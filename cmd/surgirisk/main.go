package main

import (
    "log"
    "net/http"
)

func main() {
    log.Println("Serveur lancé sur http://localhost:8080")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Serveur Go"))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
