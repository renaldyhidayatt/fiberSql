package dto

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewUsers struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
