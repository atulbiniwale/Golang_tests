package package1

import "fmt"

/* Capitalization of function "M" in Myfunc makes it an exported function.
So we can call it from outside of the package.
*/

func Myfunc() {
	fmt.Println("this is from package1")

}
