package provider

import "github.com/lucasepe/kube-craft-series/examples/custom-api-object/apis/metals/v1alpha1"

type SpotPriceGetter interface {
	Get(opts *v1alpha1.SpotPriceSpec) (*v1alpha1.SpotPriceStatus, error)
}
