package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Patrick564/cards-board-golang/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaims struct {
	Range string `json:"range"`
	jwt.RegisteredClaims
}

func hashAndSalt(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func RegisterRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "GET method not allowed", http.StatusMethodNotAllowed)
		return
	}

	u := models.User{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash, _ := hashAndSalt([]byte(u.Password))
	u.Password = hash
	u.CreatedAt = time.Now()

	res, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func LoginRoute(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	adminJwt := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			Range: "admin",
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			},
		},
	)
	signedAdminJwt, err := adminJwt.SignedString([]byte(""))
	if err != nil {
		log.Fatal(err)
	}

	resp := make(map[string]string)
	resp["admin_link"] = signedAdminJwt
	// resp["guest_link"] = signedGuestJwt

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(jsonResp)
}
