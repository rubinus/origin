package services

/*
@Time : 2019-03-06 19:02
@Author : rubinus.chu
@File : index
@project: origin
*/

type User interface {
	Search(name string)
}

func NewUser() User {
	return &A{}
}

type A struct {
}

func (a *A) Search(name string) {
	panic("implement me")
}
