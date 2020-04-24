package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
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

// // home handler function
// func home(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("User-Agent: %s", r.UserAgent())
// 	w.Write([]byte("<h1>Hello World</h1>"))
// }

func main() {
	// http.HandleFunc("/home", home)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	   
	http.HandleFunc("/", serveTemplate)

	log.Printf("Listening on port %d...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
  lp := filepath.Join("templates", "layout.html")
  fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

  // Return a 404 if the template doesn't exist
  info, err := os.Stat(fp)
  if err != nil {
    if os.IsNotExist(err) {
      http.NotFound(w, r)
      return
    }
  }

  // Return a 404 if the request is for a directory
  if info.IsDir() {
    http.NotFound(w, r)
    return
  }
  
  tmpl, err := template.ParseFiles(lp, fp)
  if err != nil {
    // Log the detailed error
    log.Println(err.Error())
    // Return a generic "Internal Server Error" message
    http.Error(w, http.StatusText(500), 500)
    return
  }

  // Render template
  err = tmpl.ExecuteTemplate(w, "layout", nil)
  if err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }
}
