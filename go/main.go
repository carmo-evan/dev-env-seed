package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/carmo-evan/strtotime"
)

type response struct {
	Data time.Time `json:"data"`
	Err  string    `json:"error"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":

		keys, ok := r.URL.Query()["value"]

		if !ok || len(keys[0]) < 1 {
			return
		}

		u, err := strtotime.Parse(keys[0], time.Now().Unix())

		t := time.Unix(u, 0)

		res := response{Data: t, Err: ""}

		if err != nil {
			w.WriteHeader(400)
			res.Err = err.Error()
		}

		b, _ := json.Marshal(res)
		fmt.Fprintln(w, string(b))
	}
}

func main() {
	http.HandleFunc("/", handler)
	var port = os.Getenv("PORT")
	log.Printf("Running golang server on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
