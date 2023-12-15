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
// Before launching this program make sure you
// have an active kubernetes cluster...
//
// $ kind create cluster
// $ go run main.go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lucasepe/kubelib"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := os.Getenv(clientcmd.RecommendedConfigPathEnvVar)
	if len(kubeconfig) == 0 {
		// if `KUBECONFIG` is not defined, use the default `$HOME/.kube/config`
		kubeconfig = clientcmd.RecommendedHomeFile
	}

	// Create a rest.Config from kubeconfig.
	restConfig, err := kubelib.RESTConfigFromKubeConfig(kubeconfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// Create a Dynamic Client
	dynClient, err := dynamic.NewForConfig(restConfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// Get A POD as Unstructured object
	unstr, err := dynClient.Resource(corev1.SchemeGroupVersion.WithResource("pods")).
		Namespace("kube-system").
		Get(context.TODO(), "kube-apiserver-kind-control-plane", metav1.GetOptions{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// Convert the Unstructured object to a POD object
	pod := corev1.Pod{}
	err = runtime.DefaultUnstructuredConverter.
		FromUnstructured(unstr.UnstructuredContent(), &pod)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// Print POD containers
	for _, con := range pod.Spec.Containers {
		fmt.Println(con.Name, " -> ", con.Image)
	}
}
