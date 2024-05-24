package types

type UserDTOBody struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}
type UserDTO struct {
	User UserDTOBody `json:"user"`
}
