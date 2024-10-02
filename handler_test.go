package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		_, err := fmt.Fprint(writer, "Hello World")
		if err != nil {
			panic(err)
		}
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		_, err := fmt.Fprint(writer, "Hello World")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		_, err := fmt.Fprint(writer, "Hi Brow!")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		_, err := fmt.Fprint(writer, "Image")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/images/thumbnail", func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		_, err := fmt.Fprint(writer, "Thubmnail")
		if err != nil {
			panic(err)
		}
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprint(w, r.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
