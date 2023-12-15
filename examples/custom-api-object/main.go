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
// $ ./scripts/gen-crd.sh
// $ kubectl apply -f crds/
// $ kubectl apply -f testdata/sample.yaml
// $ go run main.go
// $ kubectl get sp -o wide
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lucasepe/kube-craft-series/examples/custom-api-object/apis/metals/v1alpha1"
	fakeprovider "github.com/lucasepe/kube-craft-series/examples/custom-api-object/provider/fake"
	"github.com/lucasepe/kubelib"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"
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

	// Do not forget to register this API group & version to a scheme
	v1alpha1.AddToScheme(scheme.Scheme)

	client, err := kubelib.CreateRESTClient(restConfig,
		kubelib.GroupVersion(v1alpha1.SchemeGroupVersion))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// create a SpotPrice Getter
	getter := fakeprovider.Random()

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of the MetalSpotPrice
		// before attempting update.
		// RetryOnConflict uses exponential backoff
		// to avoid exhausting the apiserver
		obj := v1alpha1.SpotPrice{}
		err := client.Get().Resource("spotprice").
			Name("sample").Namespace(metav1.NamespaceDefault).
			Do(context.TODO()).
			Into(&obj)
		if err != nil {
			return err
		}

		status, err := getter.Get(&obj.Spec)
		if err != nil {
			return err
		}

		// Update the status of this SpotPrice
		obj.Status = *status.DeepCopy()
		obj.Status.SetConditions(metav1.Condition{
			Type:               "Ready",
			Status:             metav1.ConditionTrue,
			LastTransitionTime: metav1.Now(),
			Reason:             "Available",
		})

		return client.Put().
			Resource("spotprice").
			SubResource("status").
			Name("sample").
			Namespace(metav1.NamespaceDefault).
			Body(&obj).
			Do(context.TODO()).
			Error()
	})
	if retryErr != nil {
		if apierrors.IsNotFound(retryErr) {
			fmt.Fprintln(os.Stderr, retryErr.Error())
			return
		}
		fmt.Fprintf(os.Stderr, "update failed: %v\n", retryErr)
		os.Exit(1)
	}
}
