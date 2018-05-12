package iproute_test

import (
	"fmt"
	"github.com/tomponline/iproute"
    "testing"
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

func TestInvalidIP(t *testing.T) {
    srcIP, err := iproute.LocalSrcIP("invalid")

    if err.Error() != `Error: inet prefix is expected rather than "invalid".` {
        t.Error("Unexpected output: ", err.Error())
    }

    if srcIP != "" {
        t.Error("Unexpected output: ", srcIP)
    }
}
