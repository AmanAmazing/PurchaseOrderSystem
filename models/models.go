package models

type HeadInputs struct {
	Title   string
	CssFile string
}

type PostLogin struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Department string `json:"department"`
}
