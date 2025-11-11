package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/shafikshaon/url_shortener/internal/database"
	"github.com/shafikshaon/url_shortener/internal/logger"
	"github.com/shafikshaon/url_shortener/internal/models"
)

const (
	shortCodeLength = 7
	charset         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	maxRetries      = 5
)

type LinkService struct {
	linkRepo *database.LinkRepository
	userRepo *database.UserRepository
}

func NewLinkService(linkRepo *database.LinkRepository, userRepo *database.UserRepository) *LinkService {
	return &LinkService{
		linkRepo: linkRepo,
		userRepo: userRepo,
	}
}

// GenerateShortCode generates a random short code
func (s *LinkService) GenerateShortCode() (string, error) {
	ctx := context.Background()
	logger.Debugf(ctx, "Generating short code")

	for i := 0; i < maxRetries; i++ {
		code, err := generateRandomString(shortCodeLength)
		if err != nil {
			logger.Errorf(ctx, "Failed to generate random string: %+v", err)
			return "", err
		}

		// Check if code already exists
		exists, err := s.linkRepo.ShortCodeExists(code)
		if err != nil {
			logger.Errorf(ctx, "Failed to check short code existence: %+v", err)
			return "", err
		}

		if !exists {
			logger.Infof(ctx, "Generated unique short code: %s", code)
			return code, nil
		}
		logger.Debugf(ctx, "Short code %s already exists, retrying (%d/%d)", code, i+1, maxRetries)
	}

	logger.Errorf(ctx, "Failed to generate unique short code after %d retries", maxRetries)
	return "", fmt.Errorf("failed to generate unique short code after %d retries", maxRetries)
}

// generateRandomString creates a random string of given length
func generateRandomString(length int) (string, error) {
	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := range result {
		num, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}

	return string(result), nil
}

// ValidateCustomShortCode validates a custom short code
func (s *LinkService) ValidateCustomShortCode(shortCode string) error {
	if len(shortCode) < 3 || len(shortCode) > 20 {
		return fmt.Errorf("short code must be between 3 and 20 characters")
	}

	// Only allow alphanumeric and hyphens
	for _, char := range shortCode {
		if !isAlphanumeric(char) && char != '-' {
			return fmt.Errorf("short code can only contain letters, numbers, and hyphens")
		}
	}

	// Check if code is reserved
	reserved := []string{"api", "admin", "login", "signup", "dashboard", "settings", "analytics"}
	for _, r := range reserved {
		if strings.EqualFold(shortCode, r) {
			return fmt.Errorf("short code is reserved")
		}
	}

	// Check if code already exists
	exists, err := s.linkRepo.ShortCodeExists(shortCode)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("short code already exists")
	}

	return nil
}

func isAlphanumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

// CreateLink creates a new short link
func (s *LinkService) CreateLink(link *models.Link, customCode string) error {
	ctx := context.Background()
	logger.Infof(ctx, "Creating link for user ID: %d, custom code: %s", link.UserID, customCode)

	// Check user's link limit
	user, err := s.userRepo.GetByID(link.UserID)
	if err != nil {
		logger.Errorf(ctx, "User not found: %d, error: %+v", link.UserID, err)
		return fmt.Errorf("user not found")
	}

	linkCount, err := s.linkRepo.CountByUserID(link.UserID)
	if err != nil {
		logger.Errorf(ctx, "Failed to count links for user: %+v", err)
		return err
	}

	logger.Infof(ctx, "User %d has %d/%d links", link.UserID, linkCount, user.GetLinkLimit())

	if linkCount >= user.GetLinkLimit() {
		logger.Warnf(ctx, "User %d reached link limit: %d", link.UserID, user.GetLinkLimit())
		return fmt.Errorf("link limit reached for your subscription tier")
	}

	// Generate or validate short code
	if customCode != "" {
		logger.Infof(ctx, "Validating custom short code: %s", customCode)
		if err := s.ValidateCustomShortCode(customCode); err != nil {
			logger.Errorf(ctx, "Custom short code validation failed: %+v", err)
			return err
		}
		link.ShortCode = customCode
	} else {
		code, err := s.GenerateShortCode()
		if err != nil {
			return err
		}
		link.ShortCode = code
	}

	// Validate destination URL
	if !strings.HasPrefix(link.DestinationURL, "http://") && !strings.HasPrefix(link.DestinationURL, "https://") {
		link.DestinationURL = "https://" + link.DestinationURL
		logger.Infof(ctx, "Added https:// prefix to destination URL")
	}

	logger.Infof(ctx, "Creating link with short code: %s, destination: %s", link.ShortCode, link.DestinationURL)

	// Create link
	if err := s.linkRepo.Create(link); err != nil {
		logger.Errorf(ctx, "Failed to create link: %+v", err)
		return err
	}

	logger.Infof(ctx, "Successfully created link with ID: %d, short code: %s", link.ID, link.ShortCode)
	return nil
}

// GetLink retrieves a link by ID
func (s *LinkService) GetLink(linkID int64, userID int64) (*models.Link, error) {
	ctx := context.Background()
	logger.Infof(ctx, "Fetching link ID: %d for user ID: %d", linkID, userID)

	link, err := s.linkRepo.GetByID(linkID)
	if err != nil {
		logger.Errorf(ctx, "Failed to get link: %+v", err)
		return nil, err
	}

	// Check ownership
	if link.UserID != userID {
		logger.Warnf(ctx, "Unauthorized access attempt: link ID %d by user ID %d (owner: %d)", linkID, userID, link.UserID)
		return nil, fmt.Errorf("unauthorized")
	}

	logger.Infof(ctx, "Successfully fetched link ID: %d", linkID)
	return link, nil
}

// GetLinkByShortCode retrieves a link by short code
func (s *LinkService) GetLinkByShortCode(shortCode string) (*models.Link, error) {
	ctx := context.Background()
	logger.Infof(ctx, "Fetching link by short code: %s", shortCode)

	link, err := s.linkRepo.GetByShortCode(shortCode)
	if err != nil {
		logger.Errorf(ctx, "Link not found with short code: %s, error: %+v", shortCode, err)
		return nil, err
	}

	logger.Infof(ctx, "Successfully fetched link with short code: %s (ID: %d)", shortCode, link.ID)
	return link, nil
}

// ListLinks lists all links for a user with pagination
func (s *LinkService) ListLinks(userID int64, limit, offset int, search string, sortBy string) ([]*models.Link, error) {
	return s.linkRepo.GetByUserID(userID, limit, offset, search, sortBy)
}

// UpdateLink updates a link
func (s *LinkService) UpdateLink(link *models.Link, userID int64) error {
	ctx := context.Background()
	logger.Infof(ctx, "Updating link ID: %d for user ID: %d", link.ID, userID)

	// Check ownership
	existing, err := s.linkRepo.GetByID(link.ID)
	if err != nil {
		logger.Errorf(ctx, "Failed to get existing link: %+v", err)
		return err
	}

	if existing.UserID != userID {
		logger.Warnf(ctx, "Unauthorized update attempt: link ID %d by user ID %d (owner: %d)", link.ID, userID, existing.UserID)
		return fmt.Errorf("unauthorized")
	}

	// Validate destination URL
	if !strings.HasPrefix(link.DestinationURL, "http://") && !strings.HasPrefix(link.DestinationURL, "https://") {
		link.DestinationURL = "https://" + link.DestinationURL
		logger.Infof(ctx, "Added https:// prefix to destination URL")
	}

	logger.Infof(ctx, "Updating link ID: %d with new destination: %s", link.ID, link.DestinationURL)

	if err := s.linkRepo.Update(link); err != nil {
		logger.Errorf(ctx, "Failed to update link: %+v", err)
		return err
	}

	logger.Infof(ctx, "Successfully updated link ID: %d", link.ID)
	return nil
}

// DeleteLink deletes a link
func (s *LinkService) DeleteLink(linkID int64, userID int64) error {
	ctx := context.Background()
	logger.Infof(ctx, "Deleting link ID: %d for user ID: %d", linkID, userID)

	if err := s.linkRepo.Delete(linkID, userID); err != nil {
		logger.Errorf(ctx, "Failed to delete link: %+v", err)
		return err
	}

	logger.Infof(ctx, "Successfully deleted link ID: %d", linkID)
	return nil
}

// GetUserTags returns all tags for a user
func (s *LinkService) GetUserTags(userID int64) ([]string, error) {
	return s.linkRepo.GetUserTags(userID)
}

// GetLinksByTag returns links with a specific tag
func (s *LinkService) GetLinksByTag(userID int64, tag string) ([]*models.Link, error) {
	return s.linkRepo.GetByTag(userID, tag)
}
