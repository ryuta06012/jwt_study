package handler

import (
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
)

type SignupRequest struct {
	ID       string `json:"id"`
	FIRST    string `json:"first_name"`
	LAST     string `json:"last_name"`
	EMAIL    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	ID       string `json:"id"`
	FIRST    string `json:"first_name"`
	LAST     string `json:"last_name"`
	EMAIL    string `json:"email"`
	Password string `json:"password"`
}

func Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("SignUp!")
		// body読み出し
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
			return
		}

		var req SignupRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			log.Fatal(err)
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		_, err = db.Exec(
			"INSERT INTO members(first_name, last_name, password, email) VALUES (?, ?, ?, ?)",
			req.FIRST,
			req.LAST,
			hash,
			req.EMAIL,
		)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func PasswordVerify(hash, pw string) error {
	// 認証に失敗した場合は error
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("login!")
		//body読み出し
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
			return
		}

		var req LoginRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println("email=", req.EMAIL, "pass=", req.Password)
		row := db.QueryRow(
			"SELECT password FROM members WHERE first_name=? AND email=?",
			req.FIRST,
			req.EMAIL,
		)
		var hash string
		err = row.Scan(&hash)
		if err != nil {
			log.Fatal(err)
		}

		err = PasswordVerify(hash, req.Password)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println("login success: first_name = ", req.FIRST)
		log.Println(hash)
	}
}
