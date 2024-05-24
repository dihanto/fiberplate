package types

import "github.com/dihanto/fiberplate/model"

type UserRequest struct {
	User model.User `json:"user"`
}
