package User

import (
	"sample/test/IUser"
)

type User struct {
	IUser IUser.IUserInterface
}

func (u *User) Use() {
	u.IUser.AddUser(1, "sample")
}
