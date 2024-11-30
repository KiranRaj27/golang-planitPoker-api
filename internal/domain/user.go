package domain

import "time"

type User struct {
    ID        int       `db:"id"`         // Maps the "id" column in the database
    Name      string    `db:"name"`       // Maps the "name" column
    Email     string    `db:"email"`      // Maps the "email" column
    CreatedAt time.Time `db:"created_at"` // Maps the "created_at" column
}

type UserRepository interface {
    Create(user *User) error
    GetByID(id int) (*User, error)
    GetUsers() ([]*User, error) // New method
}

type UserUseCase interface {
    Register(user *User) error
    GetUser(id int) (*User, error)
    ListUsers() ([]*User,error)
}
