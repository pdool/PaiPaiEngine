package main

import "fmt"

type FriendModule struct {
}

var friendModule *FriendModule

func GetFriendModule() *FriendModule {
	if friendModule == nil {
		friendModule = new(FriendModule)
	}
	fmt.Println("xxxxxxxxx")
	return friendModule
}
