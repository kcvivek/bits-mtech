package main
import(
        "fmt"
        "flag"
)

var num = flag.Int("number", 9, "The number for finding the factorial")

func main() {
	flag.Parse()
        fmt.Println("====Factorial Program Begins====")
        i := *num
        fmt.Printf("Factorial for %d is : %d \n", i, Factorial(uint64(i)))
}
func Factorial(n uint64)(result uint64) {
        if (n > 0) {
                result = n * Factorial(n-1)
                return result
        }
        return 1
}
