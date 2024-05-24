package utils

import (
	"algogram/errors"
	"algogram/libs/models"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func login(input []string, loggedInUser *models.User, userList models.UserList) (string, error) {
	var err error
	if *loggedInUser != nil {
		return "", new(errors.UsuarioLoggeado)
	}
	user := strings.Join(input[1:], " ")
	*loggedInUser, err = userList.FindUser(user)
	if err != nil {
		return "", err
	}
	return user, nil
}

func logout(loggedInUser *models.User) error {
	if *loggedInUser == nil {
		return new(errors.NoLoggeado)
	}
	*loggedInUser = nil
	return nil
}

func publish(input []string, loggedInUser models.User, postList models.PostList, userList models.UserList) error {
	if loggedInUser == nil {
		return new(errors.NoLoggeado)
	}
	text := strings.Join(input[1:], " ")
	uid, user := loggedInUser.ViewUser()
	post := postList.SavePost(text, uid, user)
	userList.SavePost(post)
	return nil
}

func viewNextPost(loggedInUser models.User) (models.Post, error) {
	if loggedInUser == nil {
		return nil, new(errors.NoLoggeadoONoHayPosts)
	}
	post, newError := loggedInUser.ViewNextPost()
	if newError != nil {
		return nil, newError
	}
	return post, nil
}

func likePost(input []string, loggedInUser models.User, postList models.PostList) error {
	if loggedInUser == nil {
		return new(errors.NoLoggeadoOPostInexistente)
	}
	id, err := strconv.Atoi(input[1])
	if err != nil {
		return new(errors.ErrorParametros)
	}
	_, user := loggedInUser.ViewUser()
	err = postList.LikePost(id, user)
	if err != nil {
		return err
	}
	return nil
}

func showLikes(input []string, postList models.PostList) ([]string, error) {
	id, err := strconv.Atoi(input[1])
	if err != nil {
		return []string{}, new(errors.ErrorParametros)
	}
	likes, newError := postList.ShowLikes(id)
	if newError != nil || len(likes) == 0 {
		return []string{}, newError
	}
	return likes, nil
}

func App(loggedInUser models.User, userList models.UserList, postList models.PostList) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		input := strings.Split(scanner.Text(), " ")
		action := input[0]

		switch action {

		case "login":
			user, newError := login(input, &loggedInUser, userList)
			if newError != nil {
				fmt.Fprintln(os.Stdout, newError.Error())
				break
			}
			fmt.Println("Hola", user)

		case "logout":
			newError := logout(&loggedInUser)
			if newError != nil {
				fmt.Fprintln(os.Stdout, newError.Error())
				break
			}
			fmt.Println("Adios")

		case "publicar":
			newError := publish(input, loggedInUser, postList, userList)
			if newError != nil {
				fmt.Fprintln(os.Stdout, newError.Error())
				break
			}
			fmt.Println("Post publicado")

		case "ver_siguiente_feed":
			post, newError := viewNextPost(loggedInUser)
			if newError != nil {
				fmt.Fprintln(os.Stdout, newError.Error())
				break
			}
			text, id, user, _ := post.ViewPost()
			fmt.Println("Post ID", id)
			fmt.Println(user, "dijo:", text)
			fmt.Println("Likes:", post.Likes())

		case "likear_post":
			newError := likePost(input, loggedInUser, postList)
			if newError != nil {
				fmt.Fprintln(os.Stdout, newError.Error())
				break
			}
			fmt.Println("Post likeado")

		case "mostrar_likes":
			likes, newError := showLikes(input, postList)
			if newError != nil {
				fmt.Fprintln(os.Stdout, newError.Error())
				break
			}
			fmt.Println("El post tiene", len(likes), "likes:")
			for _, user := range likes {
				fmt.Printf("\t%s\n", user)
			}

		default:
			fmt.Fprintln(os.Stdout, "Comando incorrecto")
		}
	}
}
