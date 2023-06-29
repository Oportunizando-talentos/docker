package main

import (
  "fmt"
  "net/http"
)

func greetingHandler(w http.ResponseWriter, r *http.Request) {
  name := r.URL.Query().Get("name")
  fmt.Fprintf(w, "Olá, %s!", name)
}

func main() {
  http.HandleFunc("/greeting", greetingHandler)
  fmt.Println("listening on :8000")
  http.ListenAndServe(":8000", nil)
}
