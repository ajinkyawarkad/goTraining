package main

var users = map[string]User{
	"1": {
		ID:        1,
		FirstName: "Jennifer",
		LastName:  "Watson",
		Contact:   9021345424,
		Address:   "Maharashtra,Pune",
	},
}

type Creds struct {
	username string `form:"uname"`
	password string `from:"pass"`
}

type User struct {
	ID        uint
	FirstName string `form:"fname"`
	LastName  string `form:"lname"`
	Address   string `form:"address"`
	Contact   int    `form:"contact"`
	Password  string `form:"password"`
}
