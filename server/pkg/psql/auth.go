package psql

import (
	"database/sql"
	"math/rand"
	"strconv"
	"time"
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

	query := `INSERT INTO "otp" (mobile_num, otp, created_on, last_otp) VALUES('9825447691', '123456', $1, $2)`

	_, err := a.DB.Exec(query, time.Now(), time.Now())

	if err != nil {
		panic(err)
		return "", err
	}

	return otp, nil
}

func (a *Auth) VerifyOtp(mobileNum string, otp string) (bool, error) {

	const OTP_VALIDITY = 15 * 60 // 15 minutes

	query := `SELECT otp, last_otp FROM otp WHERE mobile_num = ?`
	var resultOtp string
	var last_otp time.Time

	rows, err := a.DB.Query(query, mobileNum)

	if err != nil {
		return false, err
	}

	if rows.Next() {
		if err := rows.Scan(&resultOtp, &last_otp); err != nil {
			panic(err)
		}
	}

	currTime := time.Now()
	diff := currTime.Sub(last_otp).Seconds()

	if otp != resultOtp || diff > OTP_VALIDITY {
		return false, nil
	}

	return true, nil
}
