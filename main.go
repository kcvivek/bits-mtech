package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    
    //"github.com/braintree/manners"
    //"github.com/kcvivek/bits-mtech/app/handlers"
    //"github.com/kcvivek/bits-mtech/app/health"
    //"github.com/kcvivek/bits-mtech/app/user"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/fact/{factId}", FactorialCall)
    
    log.Fatal(http.ListenAndServe(":8085", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the Factorial Portal!")
}


func FactorialCall(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    num  := vars["factId"]
    i, err := strconv.Atoi(num)

    if err != nil {
       log.Fatal(err)
    }

    var f uint64
    f = 0
    if (i > 20) {
      fmt.Fprintln(w, "Input error, can not calculate factorial for number greather than 20")
    } else if (i < 0) {
      fmt.Fprintln(w, "Input error, can not calculate factorial for negative numbers")
    } else {
        f = Factorial(uint64(i))
        fmt.Fprintf(w, "Factorial for number %d is: %d \nBye, see you in next run..\n", i, f)
    }
        
}

func Factorial(n uint64)(result uint64) {
    if (n > 0) {
       result = n * Factorial(n-1)
       return result
     }
     return 1
}
