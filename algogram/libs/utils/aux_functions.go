package utils

import (
	"algogram/errors"
	"algogram/libs/models"
	"bufio"
	"fmt"
	"os"
)

func OpenFile(file string) *os.File {
	f, err := os.Open(file)
	if err != nil {
		newError := new(errors.ErrorLeerArchivo)
		fmt.Fprintln(os.Stdout, newError.Error())
		os.Exit(0)
	}
	return f
}

func SaveUsers(users *os.File) models.UserList {
	userList := models.CreateUserList()
	id := 0
	userScanner := bufio.NewScanner(users)
	for userScanner.Scan() {
		name := userScanner.Text()
		userList.SaveUser(name, id)
		id++
	}
	return userList
}
