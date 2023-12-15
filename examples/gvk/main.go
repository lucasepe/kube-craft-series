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

	"k8s.io/apimachinery/pkg/runtime/schema"
)

func main() {
	// GroupVersion contains the "group" and
	// the "version", which uniquely identifies the API.
	gv := schema.GroupVersion{Group: "", Version: "v1"}
	fmt.Println(gv) // Output: v1

	// WithKind creates a GroupVersionKind based on the
	// GroupVersion and the passed Kind.
	gvk := gv.WithKind("Secret")
	fmt.Println(gvk) // Output: /v1, Kind=Secret

	// WithResource creates a GroupVersionResource based on the
	// GroupVersion and the passed Resource.
	gvr := gv.WithResource("secrets")
	fmt.Println(gvr) // Output: /v1, Resource=secrets

	// ParseGroupVersion turns "group/version" string
	// into a GroupVersion struct
	gv, _ = schema.ParseGroupVersion("apps/v1")
	fmt.Println(gv.Group) // Output: apps
}
