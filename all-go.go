package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}

----
package main

import "fmt"

func main() {

    fmt.Println("hello!")
	fun2()

}

fun2() {

    fmt.Println("world!")

}
---
















