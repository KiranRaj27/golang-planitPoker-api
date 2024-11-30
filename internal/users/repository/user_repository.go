package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/kiranraj27/sprint-planner/internal/domain"
)

type userRepository struct {
    db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
    return &userRepository{db}
}

func (r *userRepository) Create(user *domain.User) error {
    _, err := r.db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
    return err
}

func (r *userRepository) GetByID(id int) (*domain.User, error) {
    user := &domain.User{}
    err := r.db.Get(user, "SELECT * FROM users WHERE id = $1", id)
    return user, err
}

func (r *userRepository) GetUsers() ([]*domain.User, error) {
    var users []*domain.User
    query := "SELECT id, name, email, created_at FROM users"
    err := r.db.Select(&users, query)
    return users, err
}
