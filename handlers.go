package main

import (
	"encoding/json"
	"net/http"
)

var users []User

// SignUpHandler handles user registration.
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    hashedPassword, err := HashPassword(user.Password) // Ensure this function is accessible
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    user.Password = hashedPassword
    users = append(users, user)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

// LoginHandler handles user login.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    for _, u := range users {
        if u.Username == user.Username && CheckPasswordHash(user.Password, u.Password) { // Ensure this function is accessible
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(u)
            return
        }
    }

    http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}
