package main

import (
	"bufio"
	_ "embed"
	"log"
	"net/http"
	"time"

	"github.com/KrischanCS/roundly"
)

//go:embed style.css
var styleCSS string

func handle(pattern string, handleFunc func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := ResponseWriterStatusRecorder{w: w}

		handleFunc(&rw, r)
		log.Printf("SERVED %d %s %v", rw.Code, r.URL.Path, time.Since(start))
	})
}

func main() {
	handle("/", HomeHandleFunc)
	handle("/home", HomeHandleFunc)
	handle("/examples", ExamplesHandleFunc)
	handle("/style.css", StyleHandleFunc)

	log.Printf("Listening on port %d", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func StyleHandleFunc(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	//w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	//w.Header().Set("Pragma", "no-cache")

	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte(styleCSS))
	if err != nil {
		log.Println("Error writing response:", err)
	}
}

func write(w http.ResponseWriter, page roundly.Element) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//w.Header().Set("Refresh", "10")
	w.WriteHeader(http.StatusOK)

	bw := bufio.NewWriter(w)
	defer func() {
		err := bw.Flush()
		if err != nil {
			log.Print("Error flushing response:", err)
		}
	}()

	err := page.RenderElement(bw)
	if err != nil {
		log.Print("Error writing response: ", err)
	}
}

func HomeHandleFunc(w http.ResponseWriter, r *http.Request) {
	write(w, HomePage())
}

func ExamplesHandleFunc(w http.ResponseWriter, r *http.Request) {
	write(w, ExamplesPage())
}

type ResponseWriterStatusRecorder struct {
	w    http.ResponseWriter
	Code int
}

func (r *ResponseWriterStatusRecorder) Header() http.Header {
	return r.w.Header()
}

func (r *ResponseWriterStatusRecorder) Write(bytes []byte) (int, error) {
	return r.w.Write(bytes)
}

func (r *ResponseWriterStatusRecorder) WriteHeader(statusCode int) {
	r.Code = statusCode
	r.w.WriteHeader(statusCode)
}
