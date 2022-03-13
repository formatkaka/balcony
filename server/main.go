package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type OTPVerif struct {
	Otp int `json:"otp"`
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": "pong",
	})
}

func getOtp(c *gin.Context) {
	c.JSON(200, gin.H{
		"otp": 1234,
	})
}

func verifyOtp(c *gin.Context) {
	var otp OTPVerif

	if c.ShouldBind(&otp) == nil {
		log.Print(otp.Otp)
	} else {
		log.Println("No Post data")
	}

	if otp.Otp == 1234 {
		c.JSON(200, gin.H{
			"response": "Success",
		})
	} else {
		c.JSON(200, gin.H{
			"response": "Fail",
		})
	}

}

func main() {

	r := gin.Default()

	r.GET("/ping", ping)
	r.GET("/otp", getOtp)
	r.POST("/otp", verifyOtp)
	// mux.HandleFunc("/verifyotp", verifyOtp)
	// mux.HandleFunc("/login", loginOrSignUp)

	fmt.Println("Starting server on localhost:4000")
	r.Run("localhost:3000")
}
