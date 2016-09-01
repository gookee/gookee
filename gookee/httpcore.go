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

	var method reflect.Method
	var err bool
	controllerName := ""
	c := reflect.TypeOf(rou.Ctrl)
	if c != nil { //路由匹配成功
		method, err = c.MethodByName(strings.Title(actionName) + strings.Title(strings.ToLower(r.Method)))
		if c.Kind() == reflect.String { //controller不存在
			controllerName = reflect.ValueOf(rou.Ctrl).String()
			err = false
		} else {
			controllerName = strings.ToLower(strings.Replace(c.String(), "*controller.", "", 1))
		}
	} else {
		tmp := strings.Split(url, "/")
		controllerName = tmp[0]
		if controllerName == "" {
			controllerName = "home"
		}
		actionName = tmp[1]
		if actionName == "" {
			actionName = "index"
		}
		err = false
	}
	templatePath := "template/" + controllerName + "/" + actionName + ".gtpl"
	t := template.New(actionName + ".gtpl").Funcs(template.FuncMap{
		"nohtml": utility.NoHTML,
		"string": utility.ToStr,
	})

	//加载公用模版
	_, ferr := os.Stat("template/shared.gtpl")
	var err1 error
	if ferr == nil {
		t, err1 = t.ParseFiles(templatePath, "template/shared.gtpl")
	} else {
		t, err1 = t.ParseFiles(templatePath)
	}

	//封装httpContext对象
	httpContext := &HttpContext{t, w, r, session.Start(w, r)}

	//拦截器，返回false则终止继续执行
	if interceptFunc != nil {
		reV := reflect.ValueOf(interceptFunc)
		result := reV.Call([]reflect.Value{reflect.ValueOf(httpContext)})
		if len(result) == 1 && result[0].Kind() == reflect.Bool && !result[0].Bool() {
			return
		}
	}

	if !err && err1 != nil { //controller不存在且模版也不存在，输出报错页面
		t = template.New(actionName + ".gtpl")
		t.Parse(err1.Error())
		t.Execute(w, nil)
		return
	} else if !err { //controller不存在，输出模版页面
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
