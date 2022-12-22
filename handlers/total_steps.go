package costs

import (
	"context"
	"time"

	db "github.com/Ashneil2001/bou-steps/db"
	co "github.com/Ashneil2001/bou-steps/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func TotalSteps(c *gin.Context) {
	collection := db.MGI.Db.Collection("steps")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var totalSteps []co.Steps
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "status": "error"})
		return
	}
	err = cursor.All(ctx, &totalSteps)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "data": totalSteps})
}
