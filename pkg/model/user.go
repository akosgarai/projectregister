package model

// User type
type User struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt string
	UpdatedAt string

	password string
}
