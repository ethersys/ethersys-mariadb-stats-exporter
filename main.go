package main

import (
	"fmt"
	"mariadb_stats_exporter/collector"
	"log"
	"crypto/subtle"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"github.com/caarlos0/env/v7"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type config struct {
	IP					string		`env:"IP" envDefault:"0.0.0.0"`
	PORT				string     	`env:"PORT" envDefault:"8080"`
	HTTP_AUTH 			bool     	`env:"HTTP_AUTH" envDefault:"false"`
	HTTP_USER			string		`env:"HTTP_USER" envDefault:""`
	HTTP_PASSWORD_HASH	string		`env:"HTTP_PASSWORD_HASH" envDefault:""`
}

func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func basicAuth(username string, passwordHash string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || checkPasswordHash(pass, passwordHash) != true {
			w.Header().Set("WWW-Authenticate", `Basic realm="metrics"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorised\n"))
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Println("ERROR: MAIN01 - %+v\n", err)
	}

	collector := collector.NewRessourcesCollector()

	prometheus.MustRegister(collector)
	listen := cfg.IP + ":" + cfg.PORT

	fmt.Print("Exporter listening on: ", cfg.IP,":" , cfg.PORT," with http authentification set to ", cfg.HTTP_AUTH,"\n")

	if cfg.HTTP_AUTH {
		http.Handle("/metrics", basicAuth(cfg.HTTP_USER, cfg.HTTP_PASSWORD_HASH, promhttp.Handler()))
	} else {
		http.Handle("/metrics", promhttp.Handler())
	}

	log.Fatal(http.ListenAndServe(listen, nil))
}
