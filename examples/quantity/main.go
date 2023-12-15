package main

import (
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	strVal = "500m"
	decVal = 0.86
)

func main() {
	// From String value
	qty1, err := resource.ParseQuantity(strVal)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse quantity %q: %v", strVal, err)
		os.Exit(1)
	}
	fmt.Println(qty1.String(), "-->", qty1.AsDec())
	// Output: 500m --> 0.500

	// From Decimal value
	qty2 := resource.NewScaledQuantity(int64(decVal*1000), resource.Milli)
	fmt.Println(qty2.String(), "-->", qty2.AsDec())
	// Output: 860m --> 0.860
}
