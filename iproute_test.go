package iproute_test

import (
	"fmt"
	"github.com/tomponline/iproute"
)

func ExampleLocalSrcIP() {
	srcIP, _ := iproute.LocalSrcIP("127.0.0.2")
	fmt.Println(srcIP)
	srcIP, _ = iproute.LocalSrcIP("::1")
	fmt.Println(srcIP)
	// Output:
	// 127.0.0.1
	// ::1
}
