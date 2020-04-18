package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

const DefaultPort = 8080

var port int

// this will be called before main function
func init() {
	// get viper to be able to read env
	viper.SetEnvPrefix("CFG")
	viper.AutomaticEnv()

	// set port from the environment variable CFG_PORT otherwise default it to 8080
	if port = viper.GetInt("PORT"); port == 0 {
		log.Printf("Defaulting port to %d", DefaultPort)
		port = DefaultPort
	}

}

// index handler function
func index(w http.ResponseWriter, r *http.Request) {
	log.Printf("User-Agent: %s", r.UserAgent())
	w.Write([]byte("<h1>Hello World</h1>"))
}

func main() {
	http.HandleFunc("/", index)

	log.Printf("Listening on port %d...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
