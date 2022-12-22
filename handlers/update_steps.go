package costs

import (
	"context"
	"time"

	db "github.com/Ashneil2001/bou-steps/db"
	co "github.com/Ashneil2001/bou-steps/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateSteps(c *gin.Context) {
	var steps co.Steps
	c.BindJSON(&steps)
	collection := db.MGI.Db.Collection("steps")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.UpdateOne(ctx, bson.D{{}}, bson.D{{
		"$set", bson.D{

			{"steps", steps.Steps},
			{"stepsType", steps.StepsType},
			{"stepsDate", steps.StepsDate},
			{"stepsDescription", steps.StepsDescription},
		},
	}})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "Steps updated", "message": "Steps updated successfully"})
}
