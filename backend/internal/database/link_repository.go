package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"github.com/shafikshaon/url_shortener/internal/models"
)

type LinkRepository struct {
	db *sql.DB
}

func NewLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{db: db}
}

func (r *LinkRepository) Create(link *models.Link) error {
	query := `
		INSERT INTO links (user_id, short_code, destination_url, title, tags, expires_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(
		query,
		link.UserID,
		link.ShortCode,
		link.DestinationURL,
		link.Title,
		pq.Array(link.Tags),
		link.ExpiresAt,
	).Scan(&link.ID, &link.CreatedAt, &link.UpdatedAt)

	if err != nil {
		return fmt.Errorf("error creating link: %w", err)
	}

	return nil
}

func (r *LinkRepository) GetByShortCode(shortCode string) (*models.Link, error) {
	link := &models.Link{}
	query := `
		SELECT id, user_id, short_code, destination_url, title, tags, expires_at, created_at, updated_at
		FROM links
		WHERE short_code = $1
	`

	err := r.db.QueryRow(query, shortCode).Scan(
		&link.ID,
		&link.UserID,
		&link.ShortCode,
		&link.DestinationURL,
		&link.Title,
		pq.Array(&link.Tags),
		&link.ExpiresAt,
		&link.CreatedAt,
		&link.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("link not found")
	}

	if err != nil {
		return nil, fmt.Errorf("error getting link: %w", err)
	}

	return link, nil
}

func (r *LinkRepository) GetByID(id int64) (*models.Link, error) {
	link := &models.Link{}
	query := `
		SELECT id, user_id, short_code, destination_url, title, tags, expires_at, created_at, updated_at
		FROM links
		WHERE id = $1
	`

	err := r.db.QueryRow(query, id).Scan(
		&link.ID,
		&link.UserID,
		&link.ShortCode,
		&link.DestinationURL,
		&link.Title,
		pq.Array(&link.Tags),
		&link.ExpiresAt,
		&link.CreatedAt,
		&link.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("link not found")
	}

	if err != nil {
		return nil, fmt.Errorf("error getting link: %w", err)
	}

	return link, nil
}

func (r *LinkRepository) GetByUserID(userID int64, limit, offset int, search string, sortBy string) ([]*models.Link, error) {
	query := `
		SELECT id, user_id, short_code, destination_url, title, tags, expires_at, created_at, updated_at
		FROM links
		WHERE user_id = $1
	`
	args := []interface{}{userID}
	argCount := 1

	// Add search filter
	if search != "" {
		argCount++
		query += fmt.Sprintf(" AND (short_code ILIKE $%d OR destination_url ILIKE $%d OR title ILIKE $%d)", argCount, argCount, argCount)
		args = append(args, "%"+search+"%")
	}

	// Add sorting
	switch sortBy {
	case "clicks":
		query += " ORDER BY (SELECT COUNT(*) FROM clicks WHERE link_id = links.id) DESC"
	case "created_asc":
		query += " ORDER BY created_at ASC"
	default: // created_desc
		query += " ORDER BY created_at DESC"
	}

	// Add pagination
	argCount++
	query += fmt.Sprintf(" LIMIT $%d", argCount)
	args = append(args, limit)

	argCount++
	query += fmt.Sprintf(" OFFSET $%d", argCount)
	args = append(args, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error getting links: %w", err)
	}
	defer rows.Close()

	links := []*models.Link{}
	for rows.Next() {
		link := &models.Link{}
		err := rows.Scan(
			&link.ID,
			&link.UserID,
			&link.ShortCode,
			&link.DestinationURL,
			&link.Title,
			pq.Array(&link.Tags),
			&link.ExpiresAt,
			&link.CreatedAt,
			&link.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning link: %w", err)
		}
		links = append(links, link)
	}

	return links, nil
}

func (r *LinkRepository) CountByUserID(userID int64) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM links WHERE user_id = $1`

	err := r.db.QueryRow(query, userID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error counting links: %w", err)
	}

	return count, nil
}

func (r *LinkRepository) Update(link *models.Link) error {
	query := `
		UPDATE links
		SET destination_url = $1, title = $2, tags = $3, expires_at = $4
		WHERE id = $5
		RETURNING updated_at
	`

	err := r.db.QueryRow(
		query,
		link.DestinationURL,
		link.Title,
		pq.Array(link.Tags),
		link.ExpiresAt,
		link.ID,
	).Scan(&link.UpdatedAt)

	if err != nil {
		return fmt.Errorf("error updating link: %w", err)
	}

	return nil
}

func (r *LinkRepository) Delete(id int64, userID int64) error {
	query := `DELETE FROM links WHERE id = $1 AND user_id = $2`

	result, err := r.db.Exec(query, id, userID)
	if err != nil {
		return fmt.Errorf("error deleting link: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("link not found or unauthorized")
	}

	return nil
}

func (r *LinkRepository) ShortCodeExists(shortCode string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM links WHERE short_code = $1)`

	err := r.db.QueryRow(query, shortCode).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking short code: %w", err)
	}

	return exists, nil
}

func (r *LinkRepository) GetByTag(userID int64, tag string) ([]*models.Link, error) {
	query := `
		SELECT id, user_id, short_code, destination_url, title, tags, expires_at, created_at, updated_at
		FROM links
		WHERE user_id = $1 AND $2 = ANY(tags)
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query, userID, tag)
	if err != nil {
		return nil, fmt.Errorf("error getting links by tag: %w", err)
	}
	defer rows.Close()

	links := []*models.Link{}
	for rows.Next() {
		link := &models.Link{}
		err := rows.Scan(
			&link.ID,
			&link.UserID,
			&link.ShortCode,
			&link.DestinationURL,
			&link.Title,
			pq.Array(&link.Tags),
			&link.ExpiresAt,
			&link.CreatedAt,
			&link.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning link: %w", err)
		}
		links = append(links, link)
	}

	return links, nil
}

// GetUserTags returns all unique tags for a user
func (r *LinkRepository) GetUserTags(userID int64) ([]string, error) {
	query := `
		SELECT DISTINCT unnest(tags) as tag
		FROM links
		WHERE user_id = $1 AND tags IS NOT NULL
		ORDER BY tag
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user tags: %w", err)
	}
	defer rows.Close()

	tags := []string{}
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, fmt.Errorf("error scanning tag: %w", err)
		}
		if strings.TrimSpace(tag) != "" {
			tags = append(tags, tag)
		}
	}

	return tags, nil
}
