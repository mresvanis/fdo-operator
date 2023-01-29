//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Command) DeepCopyInto(out *Command) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Command.
func (in *Command) DeepCopy() *Command {
	if in == nil {
		return nil
	}
	out := new(Command)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DIUN) DeepCopyInto(out *DIUN) {
	*out = *in
	if in.AllowedKeyStorageTypes != nil {
		in, out := &in.AllowedKeyStorageTypes, &out.AllowedKeyStorageTypes
		*out = make([]KeyStorageType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DIUN.
func (in *DIUN) DeepCopy() *DIUN {
	if in == nil {
		return nil
	}
	out := new(DIUN)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskEncryptionClevis) DeepCopyInto(out *DiskEncryptionClevis) {
	*out = *in
	if in.Binding != nil {
		in, out := &in.Binding, &out.Binding
		*out = new(ServiceInfoDiskEncryptionClevisBinding)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskEncryptionClevis.
func (in *DiskEncryptionClevis) DeepCopy() *DiskEncryptionClevis {
	if in == nil {
		return nil
	}
	out := new(DiskEncryptionClevis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDOManufacturingServer) DeepCopyInto(out *FDOManufacturingServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FDOManufacturingServer.
func (in *FDOManufacturingServer) DeepCopy() *FDOManufacturingServer {
	if in == nil {
		return nil
	}
	out := new(FDOManufacturingServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FDOManufacturingServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDOManufacturingServerList) DeepCopyInto(out *FDOManufacturingServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FDOManufacturingServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FDOManufacturingServerList.
func (in *FDOManufacturingServerList) DeepCopy() *FDOManufacturingServerList {
	if in == nil {
		return nil
	}
	out := new(FDOManufacturingServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FDOManufacturingServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDOManufacturingServerSpec) DeepCopyInto(out *FDOManufacturingServerSpec) {
	*out = *in
	if in.RendezvousServers != nil {
		in, out := &in.RendezvousServers, &out.RendezvousServers
		*out = make([]RendezvousServer, len(*in))
		copy(*out, *in)
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = new(Protocols)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FDOManufacturingServerSpec.
func (in *FDOManufacturingServerSpec) DeepCopy() *FDOManufacturingServerSpec {
	if in == nil {
		return nil
	}
	out := new(FDOManufacturingServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDOManufacturingServerStatus) DeepCopyInto(out *FDOManufacturingServerStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FDOManufacturingServerStatus.
func (in *FDOManufacturingServerStatus) DeepCopy() *FDOManufacturingServerStatus {
	if in == nil {
		return nil
	}
	out := new(FDOManufacturingServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDOOnboardingServer) DeepCopyInto(out *FDOOnboardingServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FDOOnboardingServer.
func (in *FDOOnboardingServer) DeepCopy() *FDOOnboardingServer {
	if in == nil {
		return nil
	}
	out := new(FDOOnboardingServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FDOOnboardingServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDOOnboardingServerList) DeepCopyInto(out *FDOOnboardingServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FDOOnboardingServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FDOOnboardingServerList.
func (in *FDOOnboardingServerList) DeepCopy() *FDOOnboardingServerList {
	if in == nil {
		return nil
	}
	out := new(FDOOnboardingServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FDOOnboardingServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDOOnboardingServerSpec) DeepCopyInto(out *FDOOnboardingServerSpec) {
	*out = *in
	if in.ServiceInfo != nil {
		in, out := &in.ServiceInfo, &out.ServiceInfo
		*out = new(ServiceInfo)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FDOOnboardingServerSpec.
func (in *FDOOnboardingServerSpec) DeepCopy() *FDOOnboardingServerSpec {
	if in == nil {
		return nil
	}
	out := new(FDOOnboardingServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDOOnboardingServerStatus) DeepCopyInto(out *FDOOnboardingServerStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FDOOnboardingServerStatus.
func (in *FDOOnboardingServerStatus) DeepCopy() *FDOOnboardingServerStatus {
	if in == nil {
		return nil
	}
	out := new(FDOOnboardingServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDORendezvousServer) DeepCopyInto(out *FDORendezvousServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FDORendezvousServer.
func (in *FDORendezvousServer) DeepCopy() *FDORendezvousServer {
	if in == nil {
		return nil
	}
	out := new(FDORendezvousServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FDORendezvousServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDORendezvousServerList) DeepCopyInto(out *FDORendezvousServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FDORendezvousServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FDORendezvousServerList.
func (in *FDORendezvousServerList) DeepCopy() *FDORendezvousServerList {
	if in == nil {
		return nil
	}
	out := new(FDORendezvousServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FDORendezvousServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDORendezvousServerSpec) DeepCopyInto(out *FDORendezvousServerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FDORendezvousServerSpec.
func (in *FDORendezvousServerSpec) DeepCopy() *FDORendezvousServerSpec {
	if in == nil {
		return nil
	}
	out := new(FDORendezvousServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDORendezvousServerStatus) DeepCopyInto(out *FDORendezvousServerStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FDORendezvousServerStatus.
func (in *FDORendezvousServerStatus) DeepCopy() *FDORendezvousServerStatus {
	if in == nil {
		return nil
	}
	out := new(FDORendezvousServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *File) DeepCopyInto(out *File) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new File.
func (in *File) DeepCopy() *File {
	if in == nil {
		return nil
	}
	out := new(File)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialUser) DeepCopyInto(out *InitialUser) {
	*out = *in
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialUser.
func (in *InitialUser) DeepCopy() *InitialUser {
	if in == nil {
		return nil
	}
	out := new(InitialUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Protocols) DeepCopyInto(out *Protocols) {
	*out = *in
	if in.DIUN != nil {
		in, out := &in.DIUN, &out.DIUN
		*out = new(DIUN)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Protocols.
func (in *Protocols) DeepCopy() *Protocols {
	if in == nil {
		return nil
	}
	out := new(Protocols)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RendezvousServer) DeepCopyInto(out *RendezvousServer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RendezvousServer.
func (in *RendezvousServer) DeepCopy() *RendezvousServer {
	if in == nil {
		return nil
	}
	out := new(RendezvousServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceInfo) DeepCopyInto(out *ServiceInfo) {
	*out = *in
	if in.InitialUser != nil {
		in, out := &in.InitialUser, &out.InitialUser
		*out = new(InitialUser)
		(*in).DeepCopyInto(*out)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]File, len(*in))
		copy(*out, *in)
	}
	if in.Commands != nil {
		in, out := &in.Commands, &out.Commands
		*out = make([]Command, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DiskEncryptionClevises != nil {
		in, out := &in.DiskEncryptionClevises, &out.DiskEncryptionClevises
		*out = make([]DiskEncryptionClevis, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceInfo.
func (in *ServiceInfo) DeepCopy() *ServiceInfo {
	if in == nil {
		return nil
	}
	out := new(ServiceInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceInfoDiskEncryptionClevisBinding) DeepCopyInto(out *ServiceInfoDiskEncryptionClevisBinding) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceInfoDiskEncryptionClevisBinding.
func (in *ServiceInfoDiskEncryptionClevisBinding) DeepCopy() *ServiceInfoDiskEncryptionClevisBinding {
	if in == nil {
		return nil
	}
	out := new(ServiceInfoDiskEncryptionClevisBinding)
	in.DeepCopyInto(out)
	return out
}
