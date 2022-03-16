package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/formatkaka/balcony/pkg/psql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type OTPVerif struct {
	Otp int `json:"otp"`
}

type application struct {
	auth *psql.Auth
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": "pong",
	})
}

func (app *application) getOtp(c *gin.Context) {
	otp, err := app.auth.GetOtp("9825447695")
	fmt.Println(err)
	c.JSON(200, gin.H{
		"otp": otp,
		"err": err,
	})
}

func (app *application) verifyOtp(c *gin.Context) {
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

	const (
		host     = "localhost"
		port     = 5432
		user     = "siddhantloya"
		password = ""
		dbname   = "balcony_db"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable", host, port, user, dbname)

	db, err1 := sql.Open("postgres", psqlInfo)

	app := &application{
		auth: &psql.Auth{DB: db},
	}

	err2 := db.Ping()
	if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
		panic("Err")
	}

	fmt.Println("DB Connected !!")
	defer db.Close()

	r.GET("/ping", ping)
	r.GET("/otp", app.getOtp)
	r.POST("/login", app.verifyOtp)
	// mux.HandleFunc("/verifyotp", verifyOtp)
	// mux.HandleFunc("/login", loginOrSignUp)

	fmt.Println("Starting server on localhost:4000")
	r.Run("localhost:3000")
}
