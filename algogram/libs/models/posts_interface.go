package models

type Post interface {
	// ViewPost returns the information of the specified post
	ViewPost() (string, int, string, int)

	// ShowLikes returns a list of users who liked the post
	ShowLikes() []string

	// Like adds a new user to the post's like dictionary
	Like(user string)

	// Likes returns the number of likes on the post
	Likes() int
}

type PostList interface {
	// SavePost adds a post to the list
	SavePost(text string, uid int, user string) Post

	// LikePost adds a new user to the like dictionary of the post identified by id
	LikePost(id int, user string) error

	// ShowLikes returns an array of users who liked a post
	ShowLikes(id int) ([]string, error)
}
