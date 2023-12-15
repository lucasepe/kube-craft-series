package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/lucasepe/kube-craft-series/examples/custom-api-object-clients/generated/metals"
	"github.com/lucasepe/kube-craft-series/examples/custom-api-object/apis/metals/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func TestGeneratedClientset(t *testing.T) {
	name := "test"
	namespace := "default"

	// Let's build a dummy object
	obj := &v1alpha1.SpotPrice{}
	obj.SetName(name)
	obj.SetNamespace(namespace)
	obj.Spec.Symbol = "AU"
	obj.Spec.Currency = "EUR"
	obj.Spec.Unit = "gram"

	// Get a rest.Config as usual
	cfg, err := newRestConfig()
	if err != nil {
		t.Fatal(err)
	}

	// Create our custom api object generated clientset
	cs, err := metals.NewForConfig(cfg)
	if err != nil {
		t.Fatal(err)
	}

	// Use it
	obj, err = cs.MetalsV1alpha1().SpotPrices(namespace).
		Create(context.TODO(), obj, v1.CreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("object uid: %v", obj.UID)
}

func newRestConfig() (*rest.Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	return clientcmd.BuildConfigFromFlags("", filepath.Join(home, ".kube", "config"))
}
