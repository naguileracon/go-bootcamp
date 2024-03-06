package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

func main() {
	/// handler
	handlerGreetings := func(w http.ResponseWriter, r *http.Request) {
		//read request body
		reqBody, _ := io.ReadAll(r.Body)
		// parsing json to map[string][string]
		var reqBodyMap map[string]string
		err := json.Unmarshal(reqBody, &reqBodyMap)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// writing response
		_, err = w.Write([]byte("Hello " + reqBodyMap["first_name"] + " " + reqBodyMap["last_name"]))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	// router
	router := chi.NewRouter()
	router.Post("/greetings", handlerGreetings)

	// start the server
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
