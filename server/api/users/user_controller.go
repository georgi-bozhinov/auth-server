package users

import (
	"database/sql"
	"encoding/json"
	"github.com/georgi-bozhinov/auth-server/server/session"
	"github.com/georgi-bozhinov/auth-server/server/storage"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type Controller struct {
	Storage storage.Storage
}

func (c *Controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []storage.User

	query := "SELECT * FROM users"
	err := c.Storage.DB().Select(&users, query)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(&users); err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	var newUser storage.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		log.Printf("Error registering user: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("Error hashing password.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	newUser.Password = string(hashedPassword)

	query := "INSERT INTO users (username, password, firstname, lastname, email) VALUES ($1, $2, $3, $4, $5)"
	tx := c.Storage.DB().MustBegin()
	tx.MustExec(query, newUser.Username, newUser.Password, newUser.FirstName, newUser.LastName, newUser.Email)
	if err := tx.Commit(); err != nil {
		log.Printf("Error registering user: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	username := params["username"]

	var user storage.User

	query := "SELECT * FROM users WHERE username=$1 LIMIT 1"
	err := c.Storage.DB().Get(&user, query, username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User with username %s not found", username)
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			log.Printf("Error getting user with username %s", username)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var loginDTO loginDTO

	if err := json.NewDecoder(r.Body).Decode(&loginDTO); err != nil {
		log.Printf("Invalid login data")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user storage.User

	query := "SELECT * FROM users WHERE username=$1 LIMIT 1"
	err := c.Storage.DB().Get(&user, query, loginDTO.Username)

	if err != nil {
		log.Printf("Unsuccessful authentication with username %s", loginDTO.Username)
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username", http.StatusUnauthorized)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password)); err != nil {
		log.Printf("Unsuccessful authentication with username %s", loginDTO.Username)
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	s := session.NewSession(session.Config{SessionTTL: time.Hour})

	if err := json.NewEncoder(w).Encode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
