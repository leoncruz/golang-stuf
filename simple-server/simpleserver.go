package simpleserver

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type User struct {
  Id int
  Name string
  Email string
}

var count = 0
var users = []User{}

func Index(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  json.NewEncoder(w).Encode(users)
}

func Post(w http.ResponseWriter, r *http.Request) {
  var user User

  err := json.NewDecoder(r.Body).Decode(&user)

  if err != nil {
    panic(err)
  }

  count = count + 1
  user.Id = count

  users = append(users, user)

  w.WriteHeader(http.StatusCreated)
  w.Header().Set("Content-Type", "application/json")

  json, err := json.Marshal(user)

  if err != nil {
    panic(err)
  }

  w.Write(json)
} 

func Show(w http.ResponseWriter, r *http.Request) {
  id, _ := strconv.Atoi(r.PathValue("id"))

  var user User

  for _, i := range users {
    if i.Id == id {
      user = i
      break
    }
  }

  if user == (User{}) {
    w.WriteHeader(http.StatusNotFound)
    w.Header().Set("Content-Type", "application/json")
    return
  }

  json, err := json.Marshal(user)

  if err != nil {
    panic(err)
  }

  w.Write(json)
}

func Update(w http.ResponseWriter, r *http.Request) {
  var user User
  var index int

  id, err := strconv.Atoi(r.PathValue("id"))

  if err != nil {
    panic(err)
  }

  err = json.NewDecoder(r.Body).Decode(&user)

  if err != nil {
    panic(err)
  }

  for forIndex, forUser := range users {
    if forUser.Id == id {
      user = forUser
      index = forIndex
      break
    }
  }

  if user == (User{}) {
    w.WriteHeader(http.StatusNotFound)
    w.Header().Set("Content-Type", "application/json")
    return
  }

  users[index] = user
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")

  json, err := json.Marshal(user)

  if err != nil {
    panic(err)
  }

  w.Write(json)
} 

func Run() {
  http.HandleFunc("GET /users", Index)
  http.HandleFunc("POST /users", Post)
  http.HandleFunc("GET /users/{id}", Show)
  http.HandleFunc("PUT /users/{id}", Update)

  log.Fatal(http.ListenAndServe(":5000", nil))
}
