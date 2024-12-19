package utils

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Group    []struct {
		Email    string `json:"email"`
		Title    string `json:"title"`
		Message  string `json:"message"`
		Birthday string `json:"birthday"`
		Nickname string `json:"nickname"`
	} `json:"group"`
}

type Birthday struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Title    string `json:"title"`
	Message  string `json:"message"`
	Birthday string `json:"birthday"`
	Nickname string `json:"nickname"`
}
