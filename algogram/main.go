package main

import (
	"algogram/errors"
	models "algogram/libs/models"
	utils "algogram/libs/utils"
	"fmt"
	"os"
)

func main() {
	var newError error
	var args = os.Args[1:]
	if len(args) != 1 {
		newError = new(errors.ErrorParametros)
		fmt.Fprintln(os.Stdout, newError.Error())
		os.Exit(0)
	}

	userList := utils.SaveUsers(utils.OpenFile(args[0]))
	postList := models.CreatePostList()

	var loggedInUser models.User

	utils.App(loggedInUser, userList, postList)
}
