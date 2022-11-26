package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	person := http.NewServeMux()
	person.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("greetings!\n"))
	})

	dog := http.NewServeMux()
	dog.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("woof\n"))
	})

	mux := http.NewServeMux()
	mux.Handle("/person/", http.StripPrefix("/person", person))
	mux.Handle("/dog/", http.StripPrefix("/dog", dog))

	// terribleSecurity := TerribleSecurityProvider("wololo")
	// mux.Handle("/hello", terribleSecurity(RequestTimer(
	// 	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.Write([]byte("Hello!\n"))
	// 	}))))

	terribleSecurity := TerribleSecurityProvider("wololo")
  wrappedMux := terribleSecurity(RequestTimer(mux))

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  30 * time.Second,
		Handler:      wrappedMux,
	}

	s.ListenAndServe()
}

func RequestTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("request time for %s: %v", r.URL.Path, end.Sub(start))
	})
}

// configurable middleware
func TerribleSecurityProvider(password string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Secret-Password") != password {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("You didn't give the secret password"))
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
