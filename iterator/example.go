package main

type collection interface {
	GetIterator() iterator
}

type user struct {
	name string
}

type usercollection struct {
	users []*user
}
func NewUserCollection(users []*user) collection {
	return usercollection{users: users}
}
func (uc usercollection) GetIterator() iterator {
	return NewUserIterator(uc.users)
}

// -------------------------------------------
type iterator interface {
	HasNext() bool
	// get current
	// Next()
}
type userIterator struct {
	index int
	users []*user
}

func NewUserIterator(users []*user) iterator {
	return userIterator{
		index: 0,
		users: users,
	}
}

func (u userIterator) HasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

// other methods

func main() {
	u1 := &user{name: "u1"}
	u2 := &user{name: "u2"}
	users := []*user{u1, u2}

	usercollection := NewUserCollection(users)
	iter := usercollection.GetIterator()
	for iter.HasNext() {
		// do something

		// break for program to run
		break
	}
}
