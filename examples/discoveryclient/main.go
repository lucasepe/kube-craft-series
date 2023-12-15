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
	"flag"
	"fmt"
	"os"

	"github.com/lucasepe/kubelib"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
)

var _ discovery.DiscoveryInterface = (*discovery.DiscoveryClient)(nil)

func main() {
	var kubeconfig string
	// First try reading the `KUBECONFIG` variable
	kubeconfig = os.Getenv(clientcmd.RecommendedConfigPathEnvVar)
	if len(kubeconfig) == 0 {
		// if `KUBECONFIG` is not defined, use the default `$HOME/.kube/config`
		kubeconfig = clientcmd.RecommendedHomeFile
	}
	// Eventually the user can specify an alternate kubeconfig file (using flags)
	flag.StringVar(&kubeconfig, clientcmd.RecommendedConfigPathFlag, kubeconfig,
		"Absolute path to the kubeconfig file.")

	var namespace string
	flag.StringVar(&namespace, "namespace", metav1.NamespaceAll, "namespace")

	flag.Parse()

	// Create a rest.Config from kubeconfig.
	restConfig, err := kubelib.RESTConfigFromKubeConfig(kubeconfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// Create a discovery client from rest.Config
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(restConfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// List available API groups
	apiGroups, err := discoveryClient.ServerGroups()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error getting API groups: %w", err)
		os.Exit(1)
	}

	// Print available API groups and versions
	fmt.Println("Available API Groups:")
	for _, group := range apiGroups.Groups {
		fmt.Printf("- Group: %s\n", group.Name)
		for _, version := range group.Versions {
			fmt.Printf("  - Version: %s\n", version.Version)
		}
	}

	// List available resources
	apiResources, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error getting API resources: %w", err)
		os.Exit(1)
	}

	// Print available API resources
	fmt.Println("\nAvailable API Resources:")
	for _, resourceList := range apiResources {
		for _, resource := range resourceList.APIResources {
			fmt.Printf("- Resource: %s\n", resource.Name)
		}
	}
}
