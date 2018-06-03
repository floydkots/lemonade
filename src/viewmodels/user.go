package viewmodels

type Profile struct {
	Title string
	Active string
	User User
}

type User struct {
	Id int
	Email string
	Firstname string
	Lastname string
	Stand Stand
}

type Stand struct {
	Address string
	City string
	State string
	Zip string
}

func GetProfile() Profile {
	result := Profile{
		Title:  "Lemonade Stand Supply - Profile",
		Active: "",
		User:   User{
			Id:        0,
			Email:     "",
			Firstname: "",
			Lastname:  "",
			Stand:     Stand{
				Address: "",
				City:    "",
				State:   "",
				Zip:     "",
			},
		},
	}
	return result
}
