package models

import "fmt"

func (u *User) GetEmail() string {
	return u.Email
}

func (u *UserWithAddress) GetEmail() string {
	return u.Email
}

func (u *UserWithAddress) GetId() int {
	return u.Id
}

func RunPointer() {
	umur := 20
	umurB := &umur

	fmt.Println(umur, *umurB)
	umur = 10
	fmt.Println(umur, *umurB)
	*umurB = 30
	fmt.Println(umur, *umurB)

	var user UserWithAddress = UserWithAddress{
		Id:    10,
		Email: "test@gmail.com",
	}

	// user.Email = SetName("test2@gmail.com")
	// user.Email = SetName("test2@gmail.com")

	SetNameNative("test2@gmail.com", user)
	fmt.Println(user.GetId())
}

func SetNamePointer(newName string, user *UserWithAddress) {
	user.Email = newName
}
func SetNameNative(newName string, user UserWithAddress) {
	user.Email = newName
}

func SetName(newName string) string {
	return newName
}
