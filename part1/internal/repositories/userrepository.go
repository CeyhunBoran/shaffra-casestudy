package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/CeyhunBoran/shaffra-casestudy/internal/models"
	"github.com/CeyhunBoran/shaffra-casestudy/internal/utils"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *utils.DB
}

func NewUserRepository(db *utils.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(user models.User) (*models.User, error) {
	if ur.db == nil {
		return nil, fmt.Errorf("database connection is not established")
	}
	// Generate a UUID for the user ID
	userId, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %w", err)
	}

	user.ID = userId.String()

	query := `
		INSERT INTO users (id, "name", email, age)
		VALUES ($1, $2, $3, $4)
		RETURNING id, "name", email, age
	`

	var createdUser models.User
	err = ur.db.Conn.QueryRow(query, user.ID, user.Name, user.Email, user.Age).Scan(
		&createdUser.ID,
		&createdUser.Name,
		&createdUser.Email,
		&createdUser.Age,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &createdUser, nil
}

func (ur *UserRepository) GetUser(userID string) (*models.User, error) {
	query := `
		SELECT id, "name", email, age
		FROM users
		WHERE id = $1
	`

	var user models.User
	err := ur.db.Conn.QueryRow(query, userID).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Age,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	return &user, nil
}

func (ur *UserRepository) UpdateUser(userID string, user models.User) (*models.User, error) {
	query := `
		UPDATE users
		SET name = COALESCE($2, name),
			email = COALESCE($3, email),
			age = COALESCE($4, age)
		WHERE id = $1
		RETURNING id, name, email, age
	`

	var updatedUser models.User
	err := ur.db.Conn.QueryRow(query, userID, user.Name, user.Email, user.Age).Scan(
		&updatedUser.ID,
		&updatedUser.Name,
		&updatedUser.Email,
		&updatedUser.Age,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &updatedUser, nil
}

func (ur *UserRepository) DeleteUser(userID string) error {
	query := `
        DELETE FROM users
        WHERE id = $1
        RETURNING id
    `

	var deletedID uuid.UUID
	err := ur.db.Conn.QueryRowContext(context.Background(), query, userID).Scan(&deletedID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found: %w", err)
		}
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if deletedID == uuid.Nil {
		return fmt.Errorf("user not found")
	}

	return nil
}
