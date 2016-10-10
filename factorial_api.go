package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/fact/{factId}", FactorialCall)
    router.HandleFunc("/factorial/{factId}", FactorialCall)
    router.HandleFunc("/todos/{todoId}", TodoShow)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the Factorial Portal!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}

func FactorialCall(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    num  := vars["factId"]
    i, err := strconv.Atoi(num)

    if err != nil {
       log.Fatal(err)
    }

    fmt.Fprintf(w, "Factorial for number %d is: %d \nBye, see you in next run..\n", i, Factorial(uint64(i)))
}

func Factorial(n uint64)(result uint64) {
    if (n > 0) {
       result = n * Factorial(n-1)
       return result
     }
     return 1
}
