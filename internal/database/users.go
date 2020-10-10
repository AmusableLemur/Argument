package database

// User contains everything required for a user
type User struct {
	ID       int64
	Username string
	Password string // Should be encrypted at all times
}

// CreateUser attempts to insert a new user to the database
func CreateUser(user User) (int64, error) {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	result, err := db.Exec(query, user.Username, user.Password)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// FindUser attempts to finds a user ID by name and password
func FindUser(user User) (int64, error) {
	query := "SELECT id FROM users WHERE username = ? AND password = ?"
	err := db.QueryRow(query, &user.Username, &user.Password).Scan(&user.ID)

	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

// GetUser finds a user by ID (if any) and returns it
func GetUser(id int64) (User, error) {
	var user User

	query := "SELECT id, username, password FROM users WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}
