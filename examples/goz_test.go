package goz

import (
	"fmt"

	"github.com/idoubi/goz"
)

func ExampleNewClient() {
	cli := goz.NewClient()

	fmt.Printf("%T", cli)
	// Output: *goz.Request
}
