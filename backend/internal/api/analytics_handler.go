package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shafikshaon/url_shortener/internal/analytics"
	"github.com/shafikshaon/url_shortener/internal/auth"
)

type AnalyticsHandler struct {
	tracker *analytics.Tracker
}

func NewAnalyticsHandler(tracker *analytics.Tracker) *AnalyticsHandler {
	return &AnalyticsHandler{
		tracker: tracker,
	}
}

// GetUserAnalytics retrieves overall analytics for the current user
func (h *AnalyticsHandler) GetUserAnalytics(c *gin.Context) {
	userID, exists := auth.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	stats, err := h.tracker.GetUserAnalytics(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve analytics"})
		return
	}

	c.JSON(http.StatusOK, stats)
}
