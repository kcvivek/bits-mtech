package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/kcvivek/bits-mtech/new/app/handlers"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", handlers.Index)
    router.HandleFunc("/fact/{factId}", handlers.FactorialCall)
    router.HandleFunc("/api/v1/fact/{factId}", handlers.FactorialHandler)
    
    log.Fatal(http.ListenAndServe(":8085", router))
}
