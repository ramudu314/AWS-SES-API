package handler

import (
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Email struct to represent the email sending request
type Email struct {
	To      string `json:"to" binding:"required,email"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

// Stats struct to hold email statistics
type Stats struct {
	sync.Mutex
	EmailsSent  int
	EmailLimit  int
	EmailWarmUp bool
}

var stats = Stats{
	EmailLimit:  10,
	EmailWarmUp: true,
}

func init() {
	if limit := os.Getenv("EMAIL_LIMIT"); limit != "" {
		// Parse the limit if set in the environment variable
		parsedLimit, err := strconv.Atoi(limit)
		if err == nil {
			stats.EmailLimit = parsedLimit
		}
	}

	if warmUp := os.Getenv("EMAIL_WARM_UP"); warmUp != "" {
		// Parse warm-up flag if set in the environment variable
		if warmUp == "false" {
			stats.EmailWarmUp = false
		}
	}
}

// Handler function for the Go app, used in Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.Default())

	// Root route for testing
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the mock SES API!"})
	})

	// Routes
	router.POST("/sendEmail", sendEmail)
	router.GET("/stats", getStats)
	router.GET("/healthcheck", healthCheck)

	// Forward the request to Gin
	router.ServeHTTP(w, r)
}

// sendEmail handles the sending of email (mock behavior)
func sendEmail(c *gin.Context) {
	var email Email
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Simulate email warming (only a few emails allowed during warm-up period)
	stats.Lock()
	defer stats.Unlock()

	if stats.EmailWarmUp && stats.EmailsSent >= stats.EmailLimit {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Email warm-up period active. Please try again later."})
		return
	}

	// Simulate email sending
	stats.EmailsSent++
	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully", "to": email.To})
}

// getStats returns the statistics of the mock SES API
func getStats(c *gin.Context) {
	stats.Lock()
	defer stats.Unlock()
	c.JSON(http.StatusOK, gin.H{
		"emails_sent": stats.EmailsSent,
		"email_limit": stats.EmailLimit,
	})
}

// healthCheck checks if the API is alive
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
