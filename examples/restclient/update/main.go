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
//
// Before launching this program make sure you
// have an active kubernetes cluster...
//
// $ kind create cluster
// $ go run main.go
//

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

	restClient, err := kubelib.CreateRESTClient(restConfig,
		kubelib.GroupVersion(corev1.SchemeGroupVersion),
		kubelib.APIPath("/api"),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// Modify the "obj" returned by Get() and
	// retry Put() until you no longer get a
	// conflict error. This way, you can preserve
	// changes made by other clients between
	// 'create' and 'update'. This is implemented below
	//	using the retry utility package included with client-go.
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Secret
		// before attempting update.
		// RetryOnConflict uses exponential backoff
		// to avoid exhausting the apiserver
		obj := corev1.Secret{}
		err := restClient.Get().Resource("secrets").
			Name("github").Namespace(metav1.NamespaceDefault).
			Do(context.TODO()).
			Into(&obj)
		if err != nil {
			return err
		}

		obj.Data["token"] = []byte("YOUR_PERSONAL_ACCESS_TOKEN_UPDATED")
		return restClient.Put().Resource("secrets").
			Name("github").Namespace(metav1.NamespaceDefault).
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

	fmt.Printf("secrets %q updated\n", "github")
}
