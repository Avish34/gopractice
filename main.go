package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
	Number  int    `json:"number"`
	Id      int    `json:"id"`
}

var users []User

func CreateFakeUser() {
	users = append(users, User{"Avish", 22, "India", 8003588616, 1})
	users = append(users, User{"Cyrus", 22, "India", 84343674828, 2})
	users = append(users, User{"Divisha", 22, "India", 8412944343, 3})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating a user")
	w.Header().Set("Content-Type", "applicatioan/json")

	if r.Body == nil {
		w.WriteHeader(http.StatusNoContent)
	}

	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	rand.Seed(time.Now().UnixNano())
	newUser.Id = rand.Intn(100)
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	idInInterger, _ := strconv.Atoi(params["id"])
	for _, user := range users {
		if user.Id == idInInterger {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode("No User found with given id")
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(users)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a user")
	w.Header().Set("Content-Type", "applicatioan/json")
	// grab id from request
	params := mux.Vars(r)
	idInInterger, _ := strconv.Atoi(params["id"])
	for i, user := range users {
		if user.Id == idInInterger {
			users = append(users[:i], users[i+1:]...)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode("No User found with given id")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my service</h1>"))

}

func main() {
	CreateFakeUser()
	fmt.Println("Welcome to backend")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/user", GetAllUser).Methods("GET")
	r.HandleFunc("/user", CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", GetOneUser).Methods("GET")
	r.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
