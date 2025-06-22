package main

import (
	"log/slog"
	"net/http"
)

func BeginLogin(w http.ResponseWriter, r *http.Request) {
	user := Datastore.GetUser()

	options, session, err := WebAuthn.BeginLogin(user)
	if err != nil {
		slog.Error("could not begin login", err)
	}

	Datastore.SaveSession(session)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	j, err := json.Marshal(options)
	if err != nil {
		http.Error(w, "Failed to marshal registration options: "+err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(j)
	if err != nil {
		panic(err)
	}
}

