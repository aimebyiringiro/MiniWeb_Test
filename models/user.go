package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User { // returning a pointer to a specific user

	return users
}

func AddUser(u User) (User, error) { // adding a specific user and return that user and its error if there is

	// manupulating the passed in user

	if u.ID != 0 {

		return User{}, errors.New("The new user must not  incluce i ")
	}
	u.ID = nextID             // give the user an id
	nextID++                  // increment a new ID since the prevoius one is takene
	users = append(users, &u) // adding the user to the array of users. We are using & becuase the
	// the user is storing pointers  to the users not reall users
	return u, nil

}

func GetUserByID(id int) (User, error) {

	for _, u := range users {

		if u.ID == id {

			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func updateUser(u User) (User, error) {

	for i, candidate := range users {

		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v not found", u.ID)
}

func RemoveUserById(id int) error {

	for i, u := range users {

		if u.ID == id {

			users = append(users[:i], users[i+1:]...)

			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id)
}
