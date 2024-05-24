package models

type User interface {
	// ViewUser returns the user's information (id and name)
	ViewUser() (int, string)

	// AddPost adds the last published post to the user's queue
	AddPost(post Post)

	// ViewNextPost returns the next post in the feed according to the affinity function
	ViewNextPost() (Post, error)
}

type UserList interface {
	// SaveUser adds a new user to the list
	SaveUser(name string, id int) User

	// FindUser searches for the user by name in the list
	FindUser(name string) (User, error)

	// SavePost saves the last published post in the post queue of each user
	SavePost(post Post)
}
