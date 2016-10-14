package handlers

import (
	"encoding/json"
	"net/http"
        "github.com/gorilla/mux"
        "fmt"
        "log"
        "strconv"
)

type FactorialResponse struct {
	Number int `json:"number"`
	Factorial uint64 `json:"factorial"`
	Message string `json:"message"`
	Error bool `json:"error"`
}

func FactorialHandler (w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    num  := vars["factId"]
    i, err := strconv.Atoi(num)

    if err != nil {
       log.Fatal(err)
    }
    var msg string = "valid factorial number, good job!"
    var m uint64 = 0
    var error bool 
    if (i < 66) {
	m = Factorial(uint64(i))
    } else if (i < 0) {
        msg = "can not accomodate negative number"
	error = true
    } else {
        msg = "can not accomodate number greater than 65"
	error = true
    }

    response := FactorialResponse{
		Number : i,
		Factorial : m,
		Message : msg,
		Error : error,
    }
    json.NewEncoder(w).Encode(response)
    return
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
    if (i > 65) {
      fmt.Fprintln(w, "Input error, can not calculate factorial for number greather than 65")
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
