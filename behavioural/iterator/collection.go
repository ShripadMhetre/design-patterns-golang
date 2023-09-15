package iterator

// User Struct - Client
type User struct {
	name string
	age  int
}

// Collection Interface
type Collection interface {
	createIterator() Iterator
}

// Concrete Collection Implementation
type UserCollection struct {
	users []*User
}

func (u *UserCollection) createIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}
