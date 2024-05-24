package models

import (
	"algogram/errors"
	"math"

	hash "github.com/FerBuono/go-data-structures/hash"
	heap "github.com/FerBuono/go-data-structures/heap"
)

type user struct {
	name  string
	id    int
	posts heap.PriorityQueue[*Post]
}

type userList struct {
	list hash.Dictionary[string, User]
}

func CreateUserList() UserList {
	l := new(userList)
	l.list = hash.NewHash[string, User]()
	return l
}

// UserList Primitives

func (l *userList) SaveUser(name string, id int) User {
	u := new(user)
	u.name = name
	u.id = id
	u.posts = heap.NewHeap(u.affinityFunc)
	l.list.Save(name, u)
	return u
}

func (l *userList) FindUser(name string) (User, error) {
	if !l.list.Contains(name) {
		newError := new(errors.UsuarioInexistente)
		return nil, newError
	}
	return l.list.Get(name), nil
}

func (l *userList) SavePost(post Post) {
	for iter := l.list.Iterator(); iter.HasNext(); {
		_, user := iter.Current()
		user.AddPost(post)
		iter.Next()
	}
}

// User Primitives

func (u user) ViewUser() (int, string) {
	return u.id, u.name
}

func (u *user) AddPost(post Post) {
	_, _, _, uid := post.ViewPost()
	if uid != u.id {
		u.posts.Enqueue(&post)
	}
}

func (u *user) ViewNextPost() (Post, error) {
	if u.posts.IsEmpty() {
		newError := new(errors.NoLoggeadoONoHayPosts)
		return nil, newError
	}
	return *u.posts.Dequeue(), nil
}

// Auxiliary functions and methods

func (u *user) affinityFunc(p1, p2 *Post) int {
	_, id1, _, uid1 := (*p1).ViewPost()
	_, id2, _, uid2 := (*p2).ViewPost()

	if int(math.Abs(float64(uid1)-float64(u.id))) == int(math.Abs(float64(uid2)-float64(u.id))) {
		return id2 - id1
	}
	return int(math.Abs(float64(uid2)-float64(u.id))) - int(math.Abs(float64(uid1)-float64(u.id)))
}
