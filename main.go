package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Get preferred outbound ip of this machine
// adapted from: https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5045"
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		type data struct {
			Env      string
			Hostname string
			IpAddr   net.IP
		}

		tmpl, err := template.ParseFiles("./ui/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		env := os.Getenv("ENV")
		if env == "" {
			env = "default"
		}

		hostname, _ := os.Hostname()
		ipAddr := GetOutboundIP()

		err = tmpl.Execute(w, data{Env: env, Hostname: hostname, IpAddr: ipAddr})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"status":  true,
			"message": "Server is healthy!",
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(response)
	})

	// serve static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))

	err := http.ListenAndServe(":5045", r)
	if err != nil {
		log.Fatal(err)
	}
}
