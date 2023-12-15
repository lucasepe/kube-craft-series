package fake

import (
	"math/rand"
	"time"

	"github.com/lucasepe/kube-craft-series/examples/custom-api-object/apis/metals/v1alpha1"
	"github.com/lucasepe/kube-craft-series/examples/custom-api-object/provider"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Random() provider.SpotPriceGetter {
	source := rand.NewSource(time.Now().UnixNano())

	return &fakeSpotPriceGetter{
		rng: rand.New(source),
	}
}

type fakeSpotPriceGetter struct {
	rng *rand.Rand
}

func (s *fakeSpotPriceGetter) Get(_ *v1alpha1.SpotPriceSpec) (*v1alpha1.SpotPriceStatus, error) {
	// generate two random float64 values
	dummy := s.randFloats(64.0, 66.0, 2)

	// convert the random float64 values to resource.Quantity
	bid := resource.NewScaledQuantity(int64(dummy[0]*1000), resource.Milli)
	ask := resource.NewScaledQuantity(int64(dummy[1]*1000), resource.Milli)

	// generate a random change value and convert to resource.Quantity
	change := resource.NewScaledQuantity(
		int64(s.randFloats(0.99, 1.10, 1)[0]*1000),
		resource.Milli)

	return &v1alpha1.SpotPriceStatus{
		Date:   metav1.Now(),
		Bid:    bid.DeepCopy(),
		Ask:    ask.DeepCopy(),
		Change: change.DeepCopy(),
	}, nil
}

func (s *fakeSpotPriceGetter) randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + s.rng.Float64()*(max-min)
	}
	return res
}
