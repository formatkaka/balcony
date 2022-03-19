package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/formatkaka/balcony/pkg/psql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type OTPVerif struct {
	Otp    string `json:"otp"`
	Mobile string `json:"mobile_num"`
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

	type MobileQueryParams struct {
		MobileNum string `form:"mobile_num"`
	}
	var mobileQP MobileQueryParams

	if c.ShouldBindQuery(&mobileQP) != nil {
		c.JSON(500, gin.H{
			"response": "Invalid Query Params",
		})
	}

	otp, err := app.auth.GetOtp(mobileQP.MobileNum)
	fmt.Println(err)
	c.JSON(200, gin.H{
		"otp": otp,
		"err": "",
	})
}

func (app *application) verifyOtp(c *gin.Context) {
	var otp OTPVerif
	var token string
	var err error

	if c.ShouldBind(&otp) != nil {
		log.Println("No Post data")
	}

	token, err = app.auth.VerifyOtp(otp.Mobile, otp.Otp)
	if err != nil {
		fmt.Print(err)
		c.JSON(200, gin.H{
			"response":      "Fail",
			"error_details": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"response": "Success",
		"token":    token,
	})

}

func main() {

	r := gin.Default()
	r.Use(cors.Default())

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
