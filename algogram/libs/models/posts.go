package models

import (
	"algogram/errors"
	"strings"

	bst "github.com/FerBuono/go-data-structures/bst"
	hash "github.com/FerBuono/go-data-structures/hash"
)

type post struct {
	text  string
	id    int
	uid   int
	user  string
	likes bst.OrderedDictionary[string, *string]
}

type postList struct {
	postDict hash.Dictionary[int, Post]
}

func CreatePostList() PostList {
	l := new(postList)
	l.postDict = hash.NewHash[int, Post]()
	return l
}

// PostList Primitives

func (l *postList) SavePost(text string, uid int, user string) Post {
	p := new(post)
	p.text, p.uid, p.user = text, uid, user
	p.id = l.postDict.Count()
	p.likes = bst.NewBST[string, *string](func(a, b string) int { return strings.Compare(a, b) })
	l.postDict.Save(p.id, p)
	return p
}

func (l *postList) LikePost(id int, user string) error {
	if !l.postDict.Contains(id) {
		newError := new(errors.NoLoggeadoOPostInexistente)
		return newError
	}
	post := l.postDict.Get(id)
	post.Like(user)
	l.postDict.Save(id, post)
	return nil
}

func (l postList) ShowLikes(id int) ([]string, error) {
	users := []string{}
	if !l.postDict.Contains(id) {
		newError := new(errors.PostInexistenteOSinLikes)
		return users, newError
	}
	post := l.postDict.Get(id)
	if post.Likes() == 0 {
		newError := new(errors.PostInexistenteOSinLikes)
		return users, newError
	}
	users = post.ShowLikes()
	return users, nil
}

// Post Primitives

func (p *post) ViewPost() (string, int, string, int) {
	return p.text, p.id, p.user, p.uid
}

func (p *post) ShowLikes() []string {
	users := []string{}
	for iter := p.likes.Iterator(); iter.HasNext(); {
		user, _ := iter.Current()
		users = append(users, user)
		iter.Next()
	}
	return users
}

func (p *post) Likes() int {
	return p.likes.Count()
}

func (p *post) Like(user string) {
	p.likes.Save(user, nil)
}
