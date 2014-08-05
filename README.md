gookee
======

go web mvc框架


#注册路由
	gookee.Route{"/{action}", "index", Home}.Regist()
	gookee.Route{"/manage/{action}", "index", Manage}.Regist()

#注册拦截器实现全局权限控制、注入等
	gookee.Func(power)

#示例

##拦截器示例

	func power(httpcontext *gookee.HttpContext) bool {
		r := httpcontext.Request
		w := httpcontext.ResponseWriter

		url := r.URL.Path
		if url != "/" && httpcontext.Session.Get("username") == "" {
			http.Redirect(w, r, "/", http.StatusFound)
			return false
		}

		return true
	}


##controller示例（home控制器 index页面）

###Post请求

	type Home struct{}

	func (h *Home) IndexPost(httpcontext *gookee.HttpContext) {
		r := httpcontext.Request
		w := httpcontext.ResponseWriter

		username := utility.GetRequestStr(r, "username")
		password := utility.SHA1(r.FormValue("password"))

		if manager.IsExist(username, password) {
			httpcontext.Session.Set("username", username)
			http.Redirect(w, r, "framework", http.StatusFound)
			return
		} else {
			w.Write([]byte("<script>alert('用户名或密码错误');history.back(-1);</script>"))
		}
	}

###Get请求

	func (h *Home) IndexGet(httpcontext *gookee.HttpContext) {
		w := httpcontext.ResponseWriter
		t := httpcontext.Template

		t.Execute(w, httpcontext.Session.Set("username", "gookee"))
	}