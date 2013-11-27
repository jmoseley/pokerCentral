package main

var userStore = make([]User, 0, 5)

type User struct {
	Entity
}
