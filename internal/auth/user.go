package auth

type User struct {
	ID       string `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	// Другие поля пользователя
}
