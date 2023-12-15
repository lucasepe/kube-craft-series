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
//

package main

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func main() {
	// Creates a new Scheme
	s := runtime.NewScheme()

	// Identify the core v1 Secret kind
	gvk := schema.GroupVersionKind{
		Version: "v1",
		Kind:    "Secret",
	}

	// It should not be recognized since
	// we haven't registered it yet.
	ok := s.Recognizes(gvk)
	fmt.Printf("%v (recognized: %t)\n", gvk, ok)
	// Output: /v1, Kind=Secret (recognized: false)

	// SchemeBuilder is a slice of functions
	// to add "things" to a scheme
	builder := runtime.SchemeBuilder{
		// store the function to add corev1 objects
		corev1.AddToScheme,
	}
	// applies all the stored functions to the scheme
	builder.AddToScheme(s)

	// now check again if the Secret kind is recognized
	ok = s.Recognizes(gvk)
	fmt.Printf("%v (recognized: %t)\n", gvk, ok)
	// Output: /v1, Kind=Secret (recognized: true)

	// We can also use the scheme to create new objects.
	obj, err := s.New(gvk)
	if err != nil {
		panic(err)
	}
	fmt.Printf("obj is a %T\n", obj)
	// Output: obj is: *v1.Secret
}
