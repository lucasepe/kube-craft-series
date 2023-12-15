package v1alpha1

import (
	"context"
	"sync"
	"time"

	"github.com/lucasepe/kube-craft-series/examples/custom-api-object/apis/metals/v1alpha1"
	"github.com/lucasepe/kubelib"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

const (
	// resourceName is the name of our custom resource
	resourceName = "spotprice"
)

var (
	registerOnce sync.Once
)

// NewClient creates a new SpotPrice Client for the given config.
func NewClient(c *rest.Config, ns string, verbose bool) (*Client, error) {
	// makes sure that register this API group & version
	// to a scheme only once and never more than once
	registerOnce.Do(func() {
		v1alpha1.AddToScheme(scheme.Scheme)
	})

	// create adhoc configured RESTClient
	rc, err := kubelib.CreateRESTClient(c,
		kubelib.GroupVersion(v1alpha1.SchemeGroupVersion),
		kubelib.Verbose(verbose),
	)
	if err != nil {
		return nil, err
	}

	return &Client{rc: rc, ns: ns}, nil
}

type Client struct {
	rc rest.Interface
	ns string
}

// Get takes name of the obj, and returns the corresponding SpotPrice object, and an error if there is any.
func (c *Client) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1alpha1.SpotPrice, err error) {
	result = &v1alpha1.SpotPrice{}
	err = c.rc.Get().
		Namespace(c.ns).
		Resource(resourceName).
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SpotPrice that match those selectors.
func (c *Client) List(ctx context.Context, opts metav1.ListOptions) (result *v1alpha1.SpotPriceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SpotPriceList{}
	err = c.rc.Get().
		Namespace(c.ns).
		Resource(resourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested SpotPrices.
func (c *Client) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.rc.Get().
		Namespace(c.ns).
		Resource(resourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a SpotPrice and creates it.
// Returns the server's representation of the SpotPrice, and an error, if there is any.
func (c *Client) Create(ctx context.Context, obj *v1alpha1.SpotPrice, opts metav1.CreateOptions) (result *v1alpha1.SpotPrice, err error) {
	result = &v1alpha1.SpotPrice{}
	err = c.rc.Post().
		Namespace(c.ns).
		Resource(resourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(obj).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a SpotPrice and updates it.
// Returns the server's representation of the SpotPrice, and an error, if there is any.
func (c *Client) Update(ctx context.Context, obj *v1alpha1.SpotPrice, opts metav1.UpdateOptions) (result *v1alpha1.SpotPrice, err error) {
	result = &v1alpha1.SpotPrice{}
	err = c.rc.Put().
		Namespace(c.ns).
		Resource(resourceName).
		Name(obj.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(obj).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus updates the object Status member.
func (c *Client) UpdateStatus(ctx context.Context, obj *v1alpha1.SpotPrice, opts metav1.UpdateOptions) (result *v1alpha1.SpotPrice, err error) {
	result = &v1alpha1.SpotPrice{}
	err = c.rc.Put().
		Namespace(c.ns).
		Resource(resourceName).
		Name(obj.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(obj).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the SpotPrice and deletes it. Returns an error if one occurs.
func (c *Client) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.rc.Delete().
		Namespace(c.ns).
		Resource(resourceName).
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *Client) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.rc.Delete().
		Namespace(c.ns).
		Resource(resourceName).
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched object.
func (c *Client) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.SpotPrice, err error) {
	result = &v1alpha1.SpotPrice{}
	err = c.rc.Patch(pt).
		Namespace(c.ns).
		Resource(resourceName).
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
