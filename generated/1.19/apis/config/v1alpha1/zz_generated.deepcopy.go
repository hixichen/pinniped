// +build !ignore_autogenerated

// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialIssuerConfig) DeepCopyInto(out *CredentialIssuerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialIssuerConfig.
func (in *CredentialIssuerConfig) DeepCopy() *CredentialIssuerConfig {
	if in == nil {
		return nil
	}
	out := new(CredentialIssuerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CredentialIssuerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialIssuerConfigKubeConfigInfo) DeepCopyInto(out *CredentialIssuerConfigKubeConfigInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialIssuerConfigKubeConfigInfo.
func (in *CredentialIssuerConfigKubeConfigInfo) DeepCopy() *CredentialIssuerConfigKubeConfigInfo {
	if in == nil {
		return nil
	}
	out := new(CredentialIssuerConfigKubeConfigInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialIssuerConfigList) DeepCopyInto(out *CredentialIssuerConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CredentialIssuerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialIssuerConfigList.
func (in *CredentialIssuerConfigList) DeepCopy() *CredentialIssuerConfigList {
	if in == nil {
		return nil
	}
	out := new(CredentialIssuerConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CredentialIssuerConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialIssuerConfigStatus) DeepCopyInto(out *CredentialIssuerConfigStatus) {
	*out = *in
	if in.Strategies != nil {
		in, out := &in.Strategies, &out.Strategies
		*out = make([]CredentialIssuerConfigStrategy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KubeConfigInfo != nil {
		in, out := &in.KubeConfigInfo, &out.KubeConfigInfo
		*out = new(CredentialIssuerConfigKubeConfigInfo)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialIssuerConfigStatus.
func (in *CredentialIssuerConfigStatus) DeepCopy() *CredentialIssuerConfigStatus {
	if in == nil {
		return nil
	}
	out := new(CredentialIssuerConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialIssuerConfigStrategy) DeepCopyInto(out *CredentialIssuerConfigStrategy) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialIssuerConfigStrategy.
func (in *CredentialIssuerConfigStrategy) DeepCopy() *CredentialIssuerConfigStrategy {
	if in == nil {
		return nil
	}
	out := new(CredentialIssuerConfigStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCProviderConfig) DeepCopyInto(out *OIDCProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCProviderConfig.
func (in *OIDCProviderConfig) DeepCopy() *OIDCProviderConfig {
	if in == nil {
		return nil
	}
	out := new(OIDCProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OIDCProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCProviderConfigList) DeepCopyInto(out *OIDCProviderConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OIDCProviderConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCProviderConfigList.
func (in *OIDCProviderConfigList) DeepCopy() *OIDCProviderConfigList {
	if in == nil {
		return nil
	}
	out := new(OIDCProviderConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OIDCProviderConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCProviderConfigSpec) DeepCopyInto(out *OIDCProviderConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCProviderConfigSpec.
func (in *OIDCProviderConfigSpec) DeepCopy() *OIDCProviderConfigSpec {
	if in == nil {
		return nil
	}
	out := new(OIDCProviderConfigSpec)
	in.DeepCopyInto(out)
	return out
}
