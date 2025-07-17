package models

type User struct {
	ID       int
	Username string
	Password string
}

func GetUserByUsername(username string) (*User, error) {
	var u User
	err := DB.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).
		Scan(&u.ID, &u.Username, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

