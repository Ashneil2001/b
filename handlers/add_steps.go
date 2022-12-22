package costs

import (
	"context"
	"time"

	db "github.com/Ashneil2001/bou-steps/db"
	co "github.com/Ashneil2001/bou-steps/models"
	"github.com/gin-gonic/gin"
)

func AddSteps(c *gin.Context) {
	var steps co.Steps
	c.BindJSON(&steps)
	collection := db.MGI.Db.Collection("steps")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, steps)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "message": "steps added successfully"})
}
