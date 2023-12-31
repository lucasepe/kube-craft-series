//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotPrice) DeepCopyInto(out *SpotPrice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotPrice.
func (in *SpotPrice) DeepCopy() *SpotPrice {
	if in == nil {
		return nil
	}
	out := new(SpotPrice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpotPrice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotPriceList) DeepCopyInto(out *SpotPriceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpotPrice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotPriceList.
func (in *SpotPriceList) DeepCopy() *SpotPriceList {
	if in == nil {
		return nil
	}
	out := new(SpotPriceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpotPriceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotPriceSpec) DeepCopyInto(out *SpotPriceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotPriceSpec.
func (in *SpotPriceSpec) DeepCopy() *SpotPriceSpec {
	if in == nil {
		return nil
	}
	out := new(SpotPriceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotPriceStatus) DeepCopyInto(out *SpotPriceStatus) {
	*out = *in
	in.Conditioned.DeepCopyInto(&out.Conditioned)
	in.Date.DeepCopyInto(&out.Date)
	out.Bid = in.Bid.DeepCopy()
	out.Ask = in.Ask.DeepCopy()
	out.Change = in.Change.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotPriceStatus.
func (in *SpotPriceStatus) DeepCopy() *SpotPriceStatus {
	if in == nil {
		return nil
	}
	out := new(SpotPriceStatus)
	in.DeepCopyInto(out)
	return out
}
