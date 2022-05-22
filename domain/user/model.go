package user

import (
	storage "github.com/septemhill/txholder/storage/user"
)

type user struct {
	id    uint64
	name  string
	email string
}

func (u *user) updateEmail(email string) {
	u.email = email
}

func (u *user) updateName(name string) {
	u.name = name
}

func (u *user) unmarshalToStorage() *storage.User {
	return unmarshalToStorage(u)
}

func unmarshalToStorage(u *user) *storage.User {
	return &storage.User{
		ID:    u.id,
		Name:  u.name,
		Email: u.email,
	}
}

func marshalToModel(u *storage.User) *user {
	return &user{
		id:    u.ID,
		name:  u.Name,
		email: u.Email,
	}
}
