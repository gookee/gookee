package gookee

import (
	"net/http"
	"os"
	"reflect"
	"regexp"
	"session"
	"strings"
	"text/template"
	"utility"
)

func ExecHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	if strings.Contains(url, ".") {
		if strings.HasSuffix(url, ".gtpl") {
			t := template.New(url)
			t.Parse("禁止访问")
			t.Execute(w, nil)
			return
		}

		http.ServeFile(w, r, "static/"+url)
		return
	}

	var rou Route
	actionName := ""
	for i := 0; i < len(routeCollection); i++ {
		rou = routeCollection[i]
		re := regexp.MustCompile(strings.Replace(rou.Url, "{action}", "(([^/]+)|)", -1))
		m := re.FindStringSubmatch(url + "/")
		if m == nil {
			continue
		}

		if len(m) > 1 && m[2] != "" {
			actionName = m[2]
		} else {
			actionName = rou.Act
		}
		break
	}

	c := reflect.TypeOf(rou.Ctrl)
	method, err := c.MethodByName(strings.Title(actionName) + strings.Title(strings.ToLower(r.Method)))
	controllerName := ""
	if c.Kind() == reflect.String {
		controllerName = reflect.ValueOf(rou.Ctrl).String()
		err = false
	} else {
		controllerName = strings.ToLower(strings.Replace(c.String(), "*controller.", "", 1))
	}
	templatePath := "template/" + controllerName + "/" + actionName + ".gtpl"
	t := template.New(actionName + ".gtpl").Funcs(template.FuncMap{
		"nohtml": utility.NoHTML,
	})
	_, ferr := os.Stat("template/shared.gtpl")
	var err1 error
	if ferr == nil {
		t, err1 = t.ParseFiles(templatePath, "template/shared.gtpl")
	} else {
		t, err1 = t.ParseFiles(templatePath)
	}

	httpContext := &HttpContext{t, w, r, session.Start(w, r)}

	if interceptFunc != nil {
		reV := reflect.ValueOf(interceptFunc)
		result := reV.Call([]reflect.Value{reflect.ValueOf(httpContext)})
		if len(result) == 1 && result[0].Kind() == reflect.Bool && !result[0].Bool() {
			return
		}
	}

	if !err && err1 != nil {
		t = template.New(actionName + ".gtpl")
		t.Parse(err1.Error())
		t.Execute(w, nil)
		return
	} else if !err {
		err2 := t.Execute(w, nil)
		if err2 != nil {
			t = template.New(actionName + ".gtpl")
			t.Parse(err2.Error())
			t.Execute(w, nil)
		}
		return
	}
	method.Func.Call([]reflect.Value{reflect.ValueOf(rou.Ctrl), reflect.ValueOf(httpContext)})
}
