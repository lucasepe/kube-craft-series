package v1alpha1

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/lucasepe/kube-craft-series/examples/custom-api-object/apis/metals/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// $ kind create cluster
//
//	$ kubectl apply -f \
//	    ../custom-api-object/crds/metals.example.org_spotprice.yaml
func TestCustomClient(t *testing.T) {
	cfg, err := newRestConfig()
	if err != nil {
		t.Fatal(err)
	}

	name := "test"
	namespace := "default"

	// Create our client instance
	client, err := NewClient(cfg, namespace, false)
	if err != nil {
		t.Fatal(err)
	}

	// Let's build a dummy object
	obj := &v1alpha1.SpotPrice{}
	obj.SetName(name)
	obj.SetNamespace(namespace)
	obj.Spec.Symbol = "AU"
	obj.Spec.Currency = "EUR"
	obj.Spec.Unit = "gram"

	// Create a SpotPrice instance
	_, err = client.Create(context.TODO(), obj, metav1.CreateOptions{})
	if err != nil {
		if !apierrors.IsAlreadyExists(err) {
			t.Fatal(err)
		}
	}

	// Get the created SpotPrice instance
	obj, err = client.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("object uid: %v", obj.UID)

	// Delete the object
	err = client.Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		t.Fatal(err)
	}

	// Get the object again (it should not exists)
	_, err = client.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		if !apierrors.IsNotFound(err) {
			t.Fatal(err)
		}
	}
}

func newRestConfig() (*rest.Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	return clientcmd.BuildConfigFromFlags("", filepath.Join(home, ".kube", "config"))
}
