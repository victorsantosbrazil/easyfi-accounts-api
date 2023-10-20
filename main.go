package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}

func main() {

	a, err := app.NewApp()

	if err != nil {
		log.Fatal(err)
	}

	a.Start()

}
