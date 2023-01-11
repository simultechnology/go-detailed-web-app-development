package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := time.Now()
		next.ServeHTTP(w, r)
		time.Sleep(time.Second)
		t := time.Now()
		d := t.Sub(s).Milliseconds()
		log.Printf("end:: %s(%d ms)\n", t.Format(time.RFC3339), d)
	})
}

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

func outputHtml(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
      <html>
        <head>
          <title>Chat</title>
        </head>
        <body>
          Let's chat!
        </body>
      </html>
    `))
}

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message))
	})
}

func main() {
	http.HandleFunc("/hello", hello)

	htmlHandler := http.HandlerFunc(outputHtml)
	http.Handle("/sample", MyMiddleware(htmlHandler))
	http.Handle("/message", messageHandler("message!!!"))
	http.Handle("/test", middlewareOne(MyMiddleware(messageHandler("test!!!"))))

	if err := http.ListenAndServe(":6080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
