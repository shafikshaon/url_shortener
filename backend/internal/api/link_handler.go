package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shafikshaon/url_shortener/config"
	"github.com/shafikshaon/url_shortener/internal/analytics"
	"github.com/shafikshaon/url_shortener/internal/auth"
	"github.com/shafikshaon/url_shortener/internal/logger"
	"github.com/shafikshaon/url_shortener/internal/middleware"
	"github.com/shafikshaon/url_shortener/internal/models"
	"github.com/shafikshaon/url_shortener/internal/service"
)

type LinkHandler struct {
	linkService *service.LinkService
	tracker     *analytics.Tracker
	config      *config.Config
}

func NewLinkHandler(linkService *service.LinkService, tracker *analytics.Tracker, cfg *config.Config) *LinkHandler {
	return &LinkHandler{
		linkService: linkService,
		tracker:     tracker,
		config:      cfg,
	}
}

type CreateLinkRequest struct {
	DestinationURL string    `json:"destination_url" binding:"required,url"`
	ShortCode      string    `json:"short_code,omitempty"`
	Title          string    `json:"title,omitempty"`
	Tags           []string  `json:"tags,omitempty"`
	ExpiresAt      *string   `json:"expires_at,omitempty"`
}

type UpdateLinkRequest struct {
	DestinationURL string   `json:"destination_url" binding:"required,url"`
	Title          string   `json:"title,omitempty"`
	Tags           []string `json:"tags,omitempty"`
	ExpiresAt      *string  `json:"expires_at,omitempty"`
}

type LinkResponse struct {
	*models.Link
	ShortURL string `json:"short_url"`
}

// CreateLink handles link creation
func (h *LinkHandler) CreateLink(c *gin.Context) {
	ctx := middleware.GetContext(c)
	logger.Infof(ctx, "CreateLink handler called")

	userID, exists := auth.GetUserID(c)
	if !exists {
		logger.Warnf(ctx, "Unauthorized access attempt to create link")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	logger.Infof(ctx, "Creating link for user ID: %d", userID)

	var req CreateLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Errorf(ctx, "Invalid request body: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link := &models.Link{
		UserID:         userID,
		DestinationURL: req.DestinationURL,
		Tags:           req.Tags,
	}

	if req.Title != "" {
		link.Title = &req.Title
	}

	// Parse expiration date if provided
	if req.ExpiresAt != nil && *req.ExpiresAt != "" {
		expiresAt, err := time.Parse(time.RFC3339, *req.ExpiresAt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expiration date format"})
			return
		}
		link.ExpiresAt = &expiresAt
	}

	// Create link with optional custom short code
	if err := h.linkService.CreateLink(link, req.ShortCode); err != nil {
		logger.Errorf(ctx, "Failed to create link: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := &LinkResponse{
		Link:     link,
		ShortURL: h.config.Server.BaseURL + "/" + link.ShortCode,
	}

	logger.Infof(ctx, "Successfully created link with ID: %d, short code: %s", link.ID, link.ShortCode)
	c.JSON(http.StatusCreated, response)
}

// GetLink retrieves a single link
func (h *LinkHandler) GetLink(c *gin.Context) {
	userID, exists := auth.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	linkID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid link ID"})
		return
	}

	link, err := h.linkService.GetLink(linkID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	response := &LinkResponse{
		Link:     link,
		ShortURL: h.config.Server.BaseURL + "/" + link.ShortCode,
	}

	c.JSON(http.StatusOK, response)
}

// ListLinks retrieves all links for the current user
func (h *LinkHandler) ListLinks(c *gin.Context) {
	userID, exists := auth.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse pagination parameters
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	search := c.Query("search")
	sortBy := c.DefaultQuery("sort", "created_desc")

	if limit > 100 {
		limit = 100
	}

	links, err := h.linkService.ListLinks(userID, limit, offset, search, sortBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve links"})
		return
	}

	// Convert to response format with short URLs
	responses := make([]*LinkResponse, len(links))
	for i, link := range links {
		responses[i] = &LinkResponse{
			Link:     link,
			ShortURL: h.config.Server.BaseURL + "/" + link.ShortCode,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"links":  responses,
		"limit":  limit,
		"offset": offset,
	})
}

// UpdateLink updates an existing link
func (h *LinkHandler) UpdateLink(c *gin.Context) {
	userID, exists := auth.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	linkID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid link ID"})
		return
	}

	var req UpdateLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link := &models.Link{
		ID:             linkID,
		DestinationURL: req.DestinationURL,
		Tags:           req.Tags,
	}

	if req.Title != "" {
		link.Title = &req.Title
	}

	// Parse expiration date if provided
	if req.ExpiresAt != nil && *req.ExpiresAt != "" {
		expiresAt, err := time.Parse(time.RFC3339, *req.ExpiresAt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expiration date format"})
			return
		}
		link.ExpiresAt = &expiresAt
	}

	if err := h.linkService.UpdateLink(link, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get updated link
	updatedLink, err := h.linkService.GetLink(linkID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated link"})
		return
	}

	response := &LinkResponse{
		Link:     updatedLink,
		ShortURL: h.config.Server.BaseURL + "/" + updatedLink.ShortCode,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteLink deletes a link
func (h *LinkHandler) DeleteLink(c *gin.Context) {
	userID, exists := auth.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	linkID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid link ID"})
		return
	}

	if err := h.linkService.DeleteLink(linkID, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link deleted successfully"})
}

// GetLinkStats retrieves analytics for a link
func (h *LinkHandler) GetLinkStats(c *gin.Context) {
	userID, exists := auth.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	linkID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid link ID"})
		return
	}

	// Verify ownership
	_, err = h.linkService.GetLink(linkID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	stats, err := h.tracker.GetStats(linkID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve stats"})
		return
	}

	c.JSON(http.StatusOK, stats)
}

// GetUserTags retrieves all tags for the current user
func (h *LinkHandler) GetUserTags(c *gin.Context) {
	userID, exists := auth.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	tags, err := h.linkService.GetUserTags(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tags"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tags": tags})
}

// Redirect handles short code redirection
func (h *LinkHandler) Redirect(c *gin.Context) {
	ctx := middleware.GetContext(c)
	shortCode := c.Param("code")

	logger.Infof(ctx, "Redirect request for short code: %s", shortCode)

	link, err := h.linkService.GetLinkByShortCode(shortCode)
	if err != nil {
		logger.Warnf(ctx, "Link not found with short code: %s", shortCode)
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	// Check if link has expired
	if link.IsExpired() {
		logger.Warnf(ctx, "Link expired: short code %s, link ID: %d", shortCode, link.ID)
		c.JSON(http.StatusGone, gin.H{"error": "Link has expired"})
		return
	}

	logger.Infof(ctx, "Redirecting short code %s to: %s (link ID: %d)", shortCode, link.DestinationURL, link.ID)

	// Track click asynchronously (don't block redirect)
	go h.tracker.TrackClick(link.ID, c.ClientIP(), c.Request)

	// Perform redirect
	c.Redirect(http.StatusFound, link.DestinationURL)
}
