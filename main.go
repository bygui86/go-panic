package main

import "fmt"

func main() {
	x := 20
	y := 10
	printOperations(x, y)
	fmt.Println("Everything ok!")
	fmt.Println()

	a := 20
	b := 0
	printOperations(a, b)
	fmt.Println("Everything ok!")
	fmt.Println()

	alpha := 20
	beta := 0
	printOperationsWithAlternative(alpha, beta)
	fmt.Println("Everything ok!")
}

/*
	When a function panics, it’s execution is stopped,
	any deferred functions are executed and then the function returns to its caller.

	Recover is a function provided by Go which takes control of a panicking goroutine.
	recover() can only be used inside deferred functions.
	If you call recover()during normal flow, it will simply return nil.
	However, if a goroutine is panicking, recover()will capture the panic value.
*/
func printOperations(x, y int) {
	defer panicRecover()

	sum, subtract, multiply, divide := x+y, x-y, x*y, x/y
	fmt.Printf("Operations: sum=%v, subtract=%v, multiply=%v, divide=%v \n", sum, subtract, multiply, divide)
}

func panicRecover() {
	err := recover()
	if err != nil {
		fmt.Printf("Recover from panic: %v \n", err)
	}
}

/*
	The above method is quite useful in logging panic data and providing graceful exit,
	which can be very useful in debugging and fixing issues in the application.
	But we can also use panic recovery to provide alternative logic flows when one flow panics.
*/
func printOperationsWithAlternative(x, y int) {
	defer panicRecoverWithAlternative(x, y)

	sum, subtract, multiply, divide := x+y, x-y, x*y, x/y
	fmt.Printf("Operations: sum=%v, subtract=%v, multiply=%v, divide=%v \n", sum, subtract, multiply, divide)
}

func panicRecoverWithAlternative(x, y int) {
	err := recover()
	if err != nil {
		fmt.Printf("Recover from panic: %v \n", err)
		fmt.Println("Proceeding to alternative flow skipping division.")
		printOperationsSkipDivide(x, y)
	}
}

func printOperationsSkipDivide(x, y int) {
	sum, subtract, multiply := x+y, x-y, x*y
	fmt.Printf("Operations: sum=%v, subtract=%v, multiply=%v \n", sum, subtract, multiply)
}
