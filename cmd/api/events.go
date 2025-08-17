package main

import (
	"event-app/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) createEvent(c *gin.Context) {
	var event database.Events

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.models.Events.insert(&event)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "❌ Failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, event)
}

func (app *application) getEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "❌ Invalid event ID"})
	}

	event, err := app.models.Events.Get(id)

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "❌ Event not found"})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "❌ Failed to retreive event"})
	}

	c.JSON(http.StatusOK, event)
}

func (app *application) getAllEvent(c *gin.Context)
func (app *application) updateEvent(c *gin.Context)
func (app *application) deleteEvent(c *gin.Context)
