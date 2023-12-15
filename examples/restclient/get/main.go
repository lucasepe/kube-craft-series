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
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// First try reading the `KUBECONFIG` variable
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

	restClient, err := kubelib.CreateRESTClient(restConfig,
		kubelib.GroupVersion(corev1.SchemeGroupVersion),
		kubelib.APIPath("/api"),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	obj := corev1.Secret{}

	err = restClient.Get().
		Name("github").
		Namespace(metav1.NamespaceDefault).
		Resource("secrets").
		Do(context.TODO()).
		Into(&obj)
	if err != nil {
		if apierrors.IsNotFound(err) {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	token := string(obj.Data["token"])
	fmt.Printf("token: %s\n", token)
}
