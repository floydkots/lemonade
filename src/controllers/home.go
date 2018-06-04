package controllers

import (
	"text/template"
	"net/http"
	"lemonade/src/viewmodels"
	"lemonade/src/controllers/util"
	"lemonade/src/models"
	"fmt"
)

type homeController struct {
	template *template.Template
	loginTemplate *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	this.template.Execute(responseWriter, vm)
}

func (this *homeController) login(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	responseWriter.Header().Add("Content-Type", "text/html")

	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")

		member, err := models.GetMember(email, password)

		if err == nil {
			println("Found member")
			session, err := models.CreateSession(member)
			if err == nil {
				var cookie http.Cookie
				cookie.Name = "sessionId"
				cookie.Value = session.SessionId()
				responseWriter.Header().Add("Set-Cookie", cookie.String())
			}
		} else {
			println("Did not find member")
			member, err := models.CreateMember(email, password)

			if err == nil {
				println("Created member")
				session, err := models.CreateSession(member)
				if err == nil {
					var cookie http.Cookie
					cookie.Name = "sessionId"
					cookie.Value = session.SessionId()
					responseWriter.Header().Add("Set-Cookie", cookie.String())
				}
			} else {
				fmt.Println(err)
			}
		}
	}

	vm := viewmodels.GetLogin()

	this.loginTemplate.Execute(responseWriter, vm)
}