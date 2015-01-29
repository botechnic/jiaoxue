package jiaoxue

type Users struct {
	UserIdUserMap map[string]*User
}

func NewUsers() *Users {
	users := new(Users)
	users.UserIdUserMap = make(map[string]*User)
	return users
}

func (this *Users) UserCount() int {
	return len(this.UserIdUserMap)
}

func (this *Users) AddUser(userId string, user *User) {
	this.UserIdUserMap[userId] = user
}

func (this *Users) RemoveUser(userId string) {
	delete(this.UserIdUserMap, userId)
}

func (this *Users) HasUser(userId string) *User {
	user, _ := this.UserIdUserMap[userId]

	return user
}
