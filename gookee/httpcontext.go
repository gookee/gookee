package gookee

import (
	"net/http"
	"session"
	"text/template"
)

type HttpContext struct {
	Template       *template.Template
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Session        *session.Session
}
