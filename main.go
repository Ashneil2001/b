package main

import (
	"log"

	"github.com/gin-gonic/gin"

	db "github.com/Ashneil2001/bou-steps/db"
	co "github.com/Ashneil2001/bou-steps/handlers"
)

func main() {
	r := gin.Default()
	db.Connect()

	r.POST("/addSteps", co.AddSteps)
	r.GET("/totalSteps", co.TotalSteps)
	r.GET("/totalAmount", co.TotalAmount)
	r.DELETE("/deleteSteps/:id", co.DeleteSteps)
	r.PUT("/updateSteps", co.UpdateSteps)

	log.Fatal(r.Run(":8080"))

}
