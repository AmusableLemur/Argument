package database

// User contains everything required for a user
type User struct {
	ID       int
	Username string
	Password string
}

// GetUser finds a user by ID (if any) and returns it
func GetUser(id int) (User, error) {
	var user User

	query := "SELECT id, username, password FROM users WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}
