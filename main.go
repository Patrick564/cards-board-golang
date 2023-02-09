package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Patrick564/secret-message-wall/api"
	jwt "github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Range string `json:"range"`
	jwt.RegisteredClaims
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

	mux.HandleFunc("/register", api.RegisterRoute)
	mux.HandleFunc("/:link", resolveLinkRoute)

	err := http.ListenAndServe(":5555", mux)
	if err != nil {
		log.Fatal(err)
	}
}
