package database

type users struct{ *collection }

type Account interface {
	AccountCreateTest()
}

func InitAccount(collec *collection) Account {
	return &users{collec}
}

func (user *users) AccountCreateTest() {

}
