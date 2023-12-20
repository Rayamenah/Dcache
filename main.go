package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	ID       int
	Username string
}

type Server struct {
	db       map[int]*User
	cache    map[int]*User
	dbhit    int
	cachehit int
}

func NewServer() *Server {
	//create a mock db with 100 users
	db := make(map[int]*User)

	for i := 0; i < 100; i++ {
		db[i+1] = &User{
			ID:       i + 1,
			Username: fmt.Sprintf("user_%d", i+1),
		}
	}

	return &Server{
		db:    db,
		cache: make(map[int]*User),
	}
}

// set up a cache
func (s *Server) tryCache(id int) (*User, bool) {
	user, ok := s.cache[id]
	return user, ok
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Printf("Error getting userid: %d", err)
		return
	}

	//try querying from the cache first
	user, ok := s.tryCache(id)
	if ok {
		s.cachehit++
		json.NewEncoder(w).Encode(user)
		return
	}

	//if user is not avlaible in cache query the db
	user, ok = s.db[id]
	if !ok {
		panic("user not found")
	}
	s.dbhit++

	//insert user in cache
	s.cache[id] = user

	json.NewEncoder(w).Encode(user)
}

func main() {

}
