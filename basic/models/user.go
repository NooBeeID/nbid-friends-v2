package models

// exported
// var User string = "Ini Dari Models"

// unexported
// var data []string = []string{"noobee", "rully", "andre"}

// type Angka int

type User struct {
	Id    *int   `json:"id,omitempty"`
	Email string `json:"email"`
}
type UserWithAddress struct {
	Id    int
	Email string
}

// var User = struct {
// 	Id int
//	Email string
// }{
// 	Id: 10,
// 	Email: ""
// }
// var User2 = struct {
// 	Id int
//	Email string
// }{
// 	Id: 10,
// }

var User2 = UserWithAddress{
	Id:    20,
	Email: "email@gmail.com",
}

var id = 10
var User3 = User{
	// Id:    &id,
	Email: "",
}
