package types

type UserRequest struct {
	Name string `json:"name" validate:"required"`
	Age  uint16 `json:"age"`
}
