package models

type Grade int64

const (
	Dislike Grade = iota
	Like
	Undefined
)

func (g Grade) Int() int {
	return int(g)
}

type User struct {
	UserId uint
	Name   string
}

type Mem struct {
	Id       uint
	Link     string
	Likes    uint
	Dislikes uint
}

type Result struct {
	MemId  uint
	UserId uint
	IsLike Grade
}
