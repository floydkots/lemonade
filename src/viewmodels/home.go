package viewmodels

import (

)

type Home struct {
	Title string
	Active string
}

func GetHome() Home {
	result := Home{
		Title: "Lemonade Stand Society",
		Active: "home",
	}
	
	return result
}

type Login struct {
	Title string
	Active string
}

func GetLogin() Login {
	result := Login {
		Title: "Lemonade Stand Society - Login",
		Active: "login",
	}
	return result
}