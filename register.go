package main

import (
	"encoding/json"
	"net/http"
)

func BeginRegistration(w http.ResponseWriter, r *http.Request) {
	user := Datastore.GetUser()

	// TODO: session store
	options, _, err := WebAuthn.BeginRegistration(user)
	if err != nil {
		http.Error(w, "Failed to begin registration: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	j, err := json.Marshal(options)
	if err != nil {
		http.Error(w, "Failed to marshal registration options: "+err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(j)
}
