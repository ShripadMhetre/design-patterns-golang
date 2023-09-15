package iterator

import "fmt"

func Main() {
	// user objects
	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}

	// user collection with 2 user items
	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	// creating UserCollection Iterator
	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
