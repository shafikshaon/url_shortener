package database

import (
	"database/sql"
	"fmt"

	"github.com/shafikshaon/url_shortener/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (email, password_hash, subscription_tier)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(
		query,
		user.Email,
		user.PasswordHash,
		user.SubscriptionTier,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, password_hash, api_key, subscription_tier, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.APIKey,
		&user.SubscriptionTier,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}

	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetByID(id int64) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, password_hash, api_key, subscription_tier, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.APIKey,
		&user.SubscriptionTier,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}

	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetByAPIKey(apiKey string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, password_hash, api_key, subscription_tier, created_at, updated_at
		FROM users
		WHERE api_key = $1
	`

	err := r.db.QueryRow(query, apiKey).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.APIKey,
		&user.SubscriptionTier,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}

	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	query := `
		UPDATE users
		SET email = $1, subscription_tier = $2, api_key = $3
		WHERE id = $4
		RETURNING updated_at
	`

	err := r.db.QueryRow(
		query,
		user.Email,
		user.SubscriptionTier,
		user.APIKey,
		user.ID,
	).Scan(&user.UpdatedAt)

	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}

func (r *UserRepository) UpdatePassword(userID int64, passwordHash string) error {
	query := `UPDATE users SET password_hash = $1 WHERE id = $2`

	_, err := r.db.Exec(query, passwordHash, userID)
	if err != nil {
		return fmt.Errorf("error updating password: %w", err)
	}

	return nil
}

func (r *UserRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	return nil
}
