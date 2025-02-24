package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users = make(map[string]User)

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all users")
	var userList []User
	for _, user := range users {
		userList = append(userList, user)
	}
	json.NewEncoder(w).Encode(userList)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting specific user")
	r.ParseForm()
	userID := r.Form.Get("id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	if _, ok := users[userID]; !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(users[userID])
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating new user")
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
		return
	}
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, ok := users[newUser.ID]; ok {
		http.Error(w, "User with this ID already exists", http.StatusBadRequest)
		return
	}
	users[newUser.ID] = newUser
	json.NewEncoder(w).Encode(newUser)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating existing user")
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
		return
	}
	r.ParseForm()
	userID := r.Form.Get("id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	if _, ok := users[userID]; !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	var updatedUser User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users[userID] = updatedUser
	json.NewEncoder(w).Encode(users[userID])
}

func delete_user(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting a user")
	r.ParseForm()
	userID := r.Form.Get("id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	if _, ok := users[userID]; !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	delete(users, userID)
	fmt.Fprintf(w, "User with ID %s deleted successfully", userID)
}

// main function
func main() {
	fmt.Println("Starting the server on port :8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the API!")
	})
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/user", getUser)
	http.HandleFunc("/user/create", createUser)
	http.HandleFunc("/user/update", updateUser)
	http.HandleFunc("/user/delete", delete_user)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
