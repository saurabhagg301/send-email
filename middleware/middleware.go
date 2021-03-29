package middleware

import (
	"net/http"
	"time"

	"github.com/saurabh-arch/send-email/config"
	"github.com/saurabh-arch/send-email/logger"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var cfg config.Config

func init() {
	log = logger.GetLogger()
	cfg = config.LoadConfig()
}

// Logger function to log request time, response time and time taken to serve a request
func Logger(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()

		log.Info("Making call to ", r.Host, r.RequestURI)
		log.Info("Request time:", t1)
		defer func() {
			t2 := time.Now()
			log.Println("Response time:", t2)
			timeTaken := t2.Sub(t1)
			log.Info("Time taken:", timeTaken)
		}()
		f(w, r) // original function call
	}
}

// BasicAuthMiddleware to implement basic authentication for API's
func BasicAuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || !checkUsernameAndPassword(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised\n"))
			return
		}
		handler(w, r)
	}
}

func checkUsernameAndPassword(username, password string) bool {
	return username == cfg.Username && password == cfg.Password
}
