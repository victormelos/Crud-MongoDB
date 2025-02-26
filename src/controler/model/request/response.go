package request

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
