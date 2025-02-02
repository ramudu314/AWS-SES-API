package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

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

var stats = Stats{EmailLimit: 10, EmailWarmUp: true}

func main() {
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

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Graceful shutdown logic
	go func() {
		log.Println("Starting mock SES API on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("ListenAndServe failed:", err)
		}
	}()

	// Wait for shutdown signal (Ctrl+C)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exited")
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
