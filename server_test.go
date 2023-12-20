package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestHandleGetUser(t *testing.T) {
	s := NewServer()
	ts := httptest.NewServer(http.HandlerFunc(s.handleGetUser))
	nreq := 1000
	wg := &sync.WaitGroup{}

	//client
	//create a loop to run 1000 request to the database
	for i := 0; 1 < nreq; i++ {
		//make a goroutine to run the requests concurrently
		wg.Add(1)
		go func(i int) {
			//cycle between 1 -100 ids
			id := i%100 + 1

			//create the route url
			url := fmt.Sprintf("%s/?id=%d", ts.URL, id)
			resp, err := http.Get(url)
			if err != nil {
				t.Error((err))
			}

			user := &User{}
			if err := json.NewDecoder(resp.Body).Decode(user); err != nil {
				t.Error(err)
			}
			fmt.Printf("%+v\n", user)
			wg.Done()
		}(i)

		time.Sleep(time.Millisecond * 1)

	}
	fmt.Println("times we query the datatbase:", s.dbhit)
	fmt.Println("times we query the cache:", s.cachehit)

	wg.Wait()
}
