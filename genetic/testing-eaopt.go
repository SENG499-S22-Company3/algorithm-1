package genetic

import (
	"fmt"

	"github.com/MaxHalford/eaopt"
)

func testEaopt() {
	var test eaopt.GA

	if test.Model != nil {
		fmt.Println("test")
	}
}
