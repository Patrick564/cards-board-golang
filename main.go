package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Range string `json:"range"`
	jwt.RegisteredClaims
}

func generateLinkRoute(w http.ResponseWriter, _ *http.Request) {
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

	// guestJwt := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.RegisteredClaims{})
	// signedGuestJwt, err := guestJwt.SignedString([]byte("secret"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	resp := make(map[string]string)
	resp["admin_link"] = signedAdminJwt
	// resp["guest_link"] = signedGuestJwt

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(jsonResp)
}

func resolveLinkRoute(w http.ResponseWriter, _ *http.Request) {
	jwt.Parse("", func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("aaaa")
		}

		return []byte(""), nil
	})

	fmt.Println("generate links admin")
	io.WriteString(w, "generate links asdmin")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/generate", generateLinkRoute)
	mux.HandleFunc("/:link", resolveLinkRoute)

	err := http.ListenAndServe(":5555", mux)
	if err != nil {
		log.Fatal(err)
	}
}
