package main

import (
	"db"
	"example/controller"
	"fmt"
	"gookee"
	"io/ioutil"
	"net/http"
	"session"
	"strings"
	"time"
)

func main() {
	port := "888"
	dbConnectionString := ".\\data\\data.s3db"
	bytes, err := ioutil.ReadFile("config.ini")
	if err == nil {
		config := strings.Split(string(bytes), "\r\n")
		if len(config) > 0 {
			port = config[0]
		}
		if len(config) > 1 {
			dbConnectionString = config[1]
		}
	}

	session.Init("", 0)
	db.Init(dbConnectionString)

	gookee.Route{"/{action}", "index", &controller.Home{}}.Regist()
	gookee.Func(power)

	fmt.Println("web服务器正在运行中...")
	fmt.Println("web服务器端口：" + port)
	http.HandleFunc("/", gookee.ExecHandler)
	s := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 25,
	}
	s.ListenAndServe()
}

func power(httpcontext *gookee.HttpContext) bool {
	r := httpcontext.Request
	w := httpcontext.ResponseWriter

	url := r.URL.Path
	if url != "/" && session.Start(w, r).Get("username") == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return false
	}

	return true
}
