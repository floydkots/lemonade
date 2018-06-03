package controllers

import (
	"text/template"
	"net/http"
	"lemonade/src/controllers/util"
	"lemonade/src/viewmodels"
)

type profileController struct {
	template *template.Template
}

func (this *profileController) handle(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vm := viewmodels.GetProfile()
	if req.Method == "POST" {
		vm.User.Email = req.FormValue("email")
		vm.User.Firstname = req.FormValue("firstName")
		vm.User.Lastname = req.FormValue("lastName")
		vm.User.Stand.Address = req.FormValue("standAddress")
		vm.User.Stand.City = req.FormValue("standCity")
		vm.User.Stand.State = req.FormValue("standState")
		vm.User.Stand.Zip = req.FormValue("standZip")
	}

	responseWriter.Header().Add("Content-Type", "text/html")
	this.template.Execute(responseWriter, vm)
}
