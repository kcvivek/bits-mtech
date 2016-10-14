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
	Message string `json:"factorial"`
}

func FactorialHandler (w http.ResponseWriter, r *http.Request) {
	response := FactorialResponse{
		Message: "Factorial of xyz",
	}
	json.NewEncoder(w).Encode(response)
	return
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
