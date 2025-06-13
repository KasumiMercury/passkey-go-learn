package main

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"net/http"
)

var (
	WebAuthn *webauthn.WebAuthn
	err      error
)

func main() {
	waConfig := &webauthn.Config{
		RPDisplayName: "WebAuthn Go Learning",
		RPID:          "localhost",
		RPOrigins:     []string{"http://localhost:8080"},
	}

	if WebAuthn, err = webauthn.New(waConfig); err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.Dir("./front/")))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
