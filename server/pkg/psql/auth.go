package psql

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	DB *sql.DB
}

type Otp struct {
	mobile_num string
	otp        string
}

func (a *Auth) GetOtp(mobileNum string) (string, error) {

	rand.Seed(time.Now().UnixNano())
	otp := strconv.Itoa(rand.Intn(999999))

	query := `INSERT INTO otp (mobile_num, otp, created_on, last_otp) VALUES($1, $2, $3, $3)
				ON CONFLICT (mobile_num) 
				DO UPDATE 
				SET
					otp = $2, 
					last_otp = $3`

	_, err := a.DB.Exec(query, mobileNum, otp, time.Now())

	if err != nil {
		panic(err)
		return "", err
	}

	return otp, nil
}

func (a *Auth) VerifyOtp(mobileNum string, otp string) (string, error) {

	const OTP_VALIDITY = 15 * 60 // 15 minutes

	query := `SELECT otp, last_otp FROM otp WHERE mobile_num = $1`
	var resultOtp string
	var last_otp time.Time

	rows, err := a.DB.Query(query, mobileNum)

	if err != nil {
		return "", err
	}

	if rows.Next() {
		if err := rows.Scan(&resultOtp, &last_otp); err != nil {
			panic(err)
		}
	}

	currTime := time.Now()
	diff := currTime.Sub(last_otp).Seconds()

	if otp != resultOtp || diff > OTP_VALIDITY {
		return "", nil
	}

	token, err := a.GenerateToken(mobileNum)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *Auth) GenerateToken(mobileNum string) (string, error) {

	var id int64
	getUserQuery := "SELECT id FROM users WHERE mobile_num = $1"
	insertUserQuery := "INSERT INTO users (mobile_num, profile_pic, created_on) VALUES ($1, $2, $3) RETURNING id"
	hmacSampleSecret := []byte("SampleSecretKey")

	row, err := a.DB.Query(getUserQuery, mobileNum)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if row.Next() {
		fmt.Println("go next")
		if err := row.Scan(&id); err != nil {
			fmt.Print(err)
			panic(err)
		}
	} else {

		err := a.DB.QueryRow(insertUserQuery, mobileNum, "", time.Now()).Scan(&id)

		if err != nil {
			fmt.Println(err)
			return "", err
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"mobile_num": mobileNum,
		"id":         id,
	})

	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
