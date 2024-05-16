package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lexample

extern void helloFromC();
extern int add(int a, int b);
*/
import "C"
import "fmt"

func main() {
	// Call the C function helloFromC()
	C.helloFromC()

	// Call the C function add() and print the result
	a := 10
	b := 20
	result := C.add(C.int(a), C.int(b))
	fmt.Printf("Result of %d + %d = %d\n", a, b, result)
}
