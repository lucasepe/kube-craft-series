// MIT License

// Copyright (c) 2023 Luca Sepe

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// This code is provided as an example to accompany the book:
//
//	Mastering and Crafting Kubernetes API Objects: Mind Mapping client-go (Vol. #1)
//

package main

import (
	"fmt"

	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
)

func main() {
	// Without exceptions
	t := types.Type{Name: types.Name{Name: "Pod"}}
	private := namer.NewPrivatePluralNamer(nil)
	plural := private.Name(&t)
	fmt.Println(t.Name, " -> ", plural)

	// Using exceptions
	exceptions := map[string]string{
		"Pod":      "manypod",
		"Endpoint": "endpoints",
	}
	private = namer.NewPrivatePluralNamer(exceptions)
	plural = private.Name(&t)
	fmt.Println(t.Name, " -> ", plural)

	t = types.Type{Name: types.Name{Name: "Endpoint"}}
	private = namer.NewPrivatePluralNamer(exceptions)
	plural = private.Name(&t)
	fmt.Println(t.Name, " -> ", plural)

	// Output:
	// Pod  ->  pods
	// Pod  ->  manypod
	// Endpoint  ->  endpoints
}
