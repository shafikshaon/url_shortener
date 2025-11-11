package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/shafikshaon/url_shortener/internal/database"
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
	for i := 0; i < maxRetries; i++ {
		code, err := generateRandomString(shortCodeLength)
		if err != nil {
			return "", err
		}

		// Check if code already exists
		exists, err := s.linkRepo.ShortCodeExists(code)
		if err != nil {
			return "", err
		}

		if !exists {
			return code, nil
		}
	}

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
	// Check user's link limit
	user, err := s.userRepo.GetByID(link.UserID)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	linkCount, err := s.linkRepo.CountByUserID(link.UserID)
	if err != nil {
		return err
	}

	if linkCount >= user.GetLinkLimit() {
		return fmt.Errorf("link limit reached for your subscription tier")
	}

	// Generate or validate short code
	if customCode != "" {
		if err := s.ValidateCustomShortCode(customCode); err != nil {
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
	}

	// Create link
	if err := s.linkRepo.Create(link); err != nil {
		return err
	}

	return nil
}

// GetLink retrieves a link by ID
func (s *LinkService) GetLink(linkID int64, userID int64) (*models.Link, error) {
	link, err := s.linkRepo.GetByID(linkID)
	if err != nil {
		return nil, err
	}

	// Check ownership
	if link.UserID != userID {
		return nil, fmt.Errorf("unauthorized")
	}

	return link, nil
}

// GetLinkByShortCode retrieves a link by short code
func (s *LinkService) GetLinkByShortCode(shortCode string) (*models.Link, error) {
	return s.linkRepo.GetByShortCode(shortCode)
}

// ListLinks lists all links for a user with pagination
func (s *LinkService) ListLinks(userID int64, limit, offset int, search string, sortBy string) ([]*models.Link, error) {
	return s.linkRepo.GetByUserID(userID, limit, offset, search, sortBy)
}

// UpdateLink updates a link
func (s *LinkService) UpdateLink(link *models.Link, userID int64) error {
	// Check ownership
	existing, err := s.linkRepo.GetByID(link.ID)
	if err != nil {
		return err
	}

	if existing.UserID != userID {
		return fmt.Errorf("unauthorized")
	}

	// Validate destination URL
	if !strings.HasPrefix(link.DestinationURL, "http://") && !strings.HasPrefix(link.DestinationURL, "https://") {
		link.DestinationURL = "https://" + link.DestinationURL
	}

	return s.linkRepo.Update(link)
}

// DeleteLink deletes a link
func (s *LinkService) DeleteLink(linkID int64, userID int64) error {
	return s.linkRepo.Delete(linkID, userID)
}

// GetUserTags returns all tags for a user
func (s *LinkService) GetUserTags(userID int64) ([]string, error) {
	return s.linkRepo.GetUserTags(userID)
}

// GetLinksByTag returns links with a specific tag
func (s *LinkService) GetLinksByTag(userID int64, tag string) ([]*models.Link, error) {
	return s.linkRepo.GetByTag(userID, tag)
}
