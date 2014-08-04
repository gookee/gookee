package controller

import (
	. "example/model/dal"
	"gookee"
	"net/http"
	"utility"
)

type Home struct{}

func (h *Home) UserinfoPost(httpcontext *gookee.HttpContext) {
	r := httpcontext.Request
	w := httpcontext.ResponseWriter

	action := utility.GetRequestStr(r, "action")
	id := utility.GetRequestInt(r, "id")
	ids := utility.GetRequestStr(r, "ids")
	username := utility.GetRequestStr(r, "username")
	address := utility.GetRequestStr(r, "address")
	phone := utility.GetRequestStr(r, "phone")
	isspend := utility.GetRequestInt(r, "isspend")
	pid := utility.GetRequestInt(r, "pid")

	f_username := utility.GetRequestStr(r, "f_username")

	switch action {
	case "getAll":
		w.Write([]byte(Userinfodal.GetAll(f_username, pid)))
	case "find":
		w.Write([]byte(Userinfodal.GetRow(id)))
	case "del":
		Userinfodal.Del(id)
	case "batchDel":
		Userinfodal.DelAll(ids)
	case "add":
		r := Userinfodal.Insert(username, address, phone, isspend, pid)
		if r == 0 {
			w.Write([]byte("下级加盟商数量已有20位，不能继续添加！"))
		}
	case "edit":
		Userinfodal.Update(username, address, phone, isspend, pid, id)
	}
	w.Write([]byte(""))
}

func (h *Home) IndexPost(httpcontext *gookee.HttpContext) {
	r := httpcontext.Request
	w := httpcontext.ResponseWriter

	username := utility.GetRequestStr(r, "username")
	password := utility.SHA1(r.FormValue("password"))

	if Managerdal.IsExist(username, password) {
		httpcontext.Session.Set("username", username)
		http.Redirect(w, r, "framework", http.StatusFound)
		return
	} else {
		w.Write([]byte("<script>alert('用户名或密码错误');history.back(-1);</script>"))
	}
}

func (h *Home) LogoutGet(httpcontext *gookee.HttpContext) {
	r := httpcontext.Request
	w := httpcontext.ResponseWriter

	httpcontext.Session.Clear()
	http.Redirect(w, r, "/", http.StatusFound)
	return
}

func (h *Home) FrameworkGet(httpcontext *gookee.HttpContext) {
	w := httpcontext.ResponseWriter
	t := httpcontext.Template

	t.Execute(w, httpcontext.Session.Get("username"))
}

func (h *Home) FrameworkPost(httpcontext *gookee.HttpContext) {
	r := httpcontext.Request
	w := httpcontext.ResponseWriter

	password := utility.SHA1(r.FormValue("password"))

	Managerdal.Update(httpcontext.Session.Get("username"), password)
	w.Write([]byte(""))
}
