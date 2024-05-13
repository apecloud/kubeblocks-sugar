//go:build !ignore_autogenerated

/*
Copyright 2024.

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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Affinity) DeepCopyInto(out *Affinity) {
	*out = *in
	if in.TopologyKeys != nil {
		in, out := &in.TopologyKeys, &out.TopologyKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeLabels != nil {
		in, out := &in.NodeLabels, &out.NodeLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Affinity.
func (in *Affinity) DeepCopy() *Affinity {
	if in == nil {
		return nil
	}
	out := new(Affinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApeCloudMySQL) DeepCopyInto(out *ApeCloudMySQL) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApeCloudMySQL.
func (in *ApeCloudMySQL) DeepCopy() *ApeCloudMySQL {
	if in == nil {
		return nil
	}
	out := new(ApeCloudMySQL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApeCloudMySQL) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApeCloudMySQLList) DeepCopyInto(out *ApeCloudMySQLList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApeCloudMySQL, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApeCloudMySQLList.
func (in *ApeCloudMySQLList) DeepCopy() *ApeCloudMySQLList {
	if in == nil {
		return nil
	}
	out := new(ApeCloudMySQLList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApeCloudMySQLList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApeCloudMySQLSpec) DeepCopyInto(out *ApeCloudMySQLSpec) {
	*out = *in
	in.ClusterSpec.DeepCopyInto(&out.ClusterSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApeCloudMySQLSpec.
func (in *ApeCloudMySQLSpec) DeepCopy() *ApeCloudMySQLSpec {
	if in == nil {
		return nil
	}
	out := new(ApeCloudMySQLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApeCloudMySQLStatus) DeepCopyInto(out *ApeCloudMySQLStatus) {
	*out = *in
	in.ClusterStatus.DeepCopyInto(&out.ClusterStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApeCloudMySQLStatus.
func (in *ApeCloudMySQLStatus) DeepCopy() *ApeCloudMySQLStatus {
	if in == nil {
		return nil
	}
	out := new(ApeCloudMySQLStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBackup) DeepCopyInto(out *ClusterBackup) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.StartingDeadlineMinutes != nil {
		in, out := &in.StartingDeadlineMinutes, &out.StartingDeadlineMinutes
		*out = new(int64)
		**out = **in
	}
	if in.PITREnabled != nil {
		in, out := &in.PITREnabled, &out.PITREnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBackup.
func (in *ClusterBackup) DeepCopy() *ClusterBackup {
	if in == nil {
		return nil
	}
	out := new(ClusterBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterComponentService) DeepCopyInto(out *ClusterComponentService) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodService != nil {
		in, out := &in.PodService, &out.PodService
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterComponentService.
func (in *ClusterComponentService) DeepCopy() *ClusterComponentService {
	if in == nil {
		return nil
	}
	out := new(ClusterComponentService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterComponentSpec) DeepCopyInto(out *ClusterComponentSpec) {
	*out = *in
	if in.ServiceRefs != nil {
		in, out := &in.ServiceRefs, &out.ServiceRefs
		*out = make([]ServiceRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnabledLogs != nil {
		in, out := &in.EnabledLogs, &out.EnabledLogs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SchedulingPolicy != nil {
		in, out := &in.SchedulingPolicy, &out.SchedulingPolicy
		*out = new(SchedulingPolicy)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]ClusterComponentVolumeClaimTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]ClusterComponentService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(Issuer)
		(*in).DeepCopyInto(*out)
	}
	if in.UserResourceRefs != nil {
		in, out := &in.UserResourceRefs, &out.UserResourceRefs
		*out = new(UserResourceRefs)
		(*in).DeepCopyInto(*out)
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]InstanceTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OfflineInstances != nil {
		in, out := &in.OfflineInstances, &out.OfflineInstances
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MonitorEnabled != nil {
		in, out := &in.MonitorEnabled, &out.MonitorEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterComponentSpec.
func (in *ClusterComponentSpec) DeepCopy() *ClusterComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterComponentStatus) DeepCopyInto(out *ClusterComponentStatus) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = make(ComponentMessageMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodsReady != nil {
		in, out := &in.PodsReady, &out.PodsReady
		*out = new(bool)
		**out = **in
	}
	if in.PodsReadyTime != nil {
		in, out := &in.PodsReadyTime, &out.PodsReadyTime
		*out = (*in).DeepCopy()
	}
	if in.MembersStatus != nil {
		in, out := &in.MembersStatus, &out.MembersStatus
		*out = make([]MemberStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterComponentStatus.
func (in *ClusterComponentStatus) DeepCopy() *ClusterComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterComponentVolumeClaimTemplate) DeepCopyInto(out *ClusterComponentVolumeClaimTemplate) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterComponentVolumeClaimTemplate.
func (in *ClusterComponentVolumeClaimTemplate) DeepCopy() *ClusterComponentVolumeClaimTemplate {
	if in == nil {
		return nil
	}
	out := new(ClusterComponentVolumeClaimTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterService) DeepCopyInto(out *ClusterService) {
	*out = *in
	in.Service.DeepCopyInto(&out.Service)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterService.
func (in *ClusterService) DeepCopy() *ClusterService {
	if in == nil {
		return nil
	}
	out := new(ClusterService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.ComponentSpecs != nil {
		in, out := &in.ComponentSpecs, &out.ComponentSpecs
		*out = make([]ClusterComponentSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]ClusterService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SchedulingPolicy != nil {
		in, out := &in.SchedulingPolicy, &out.SchedulingPolicy
		*out = new(SchedulingPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(ClusterBackup)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make(map[string]ClusterComponentStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ComponentMessageMap) DeepCopyInto(out *ComponentMessageMap) {
	{
		in := &in
		*out = make(ComponentMessageMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentMessageMap.
func (in ComponentMessageMap) DeepCopy() ComponentMessageMap {
	if in == nil {
		return nil
	}
	out := new(ComponentMessageMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapRef) DeepCopyInto(out *ConfigMapRef) {
	*out = *in
	in.ResourceMeta.DeepCopyInto(&out.ResourceMeta)
	in.ConfigMap.DeepCopyInto(&out.ConfigMap)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapRef.
func (in *ConfigMapRef) DeepCopy() *ConfigMapRef {
	if in == nil {
		return nil
	}
	out := new(ConfigMapRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceTemplate) DeepCopyInto(out *InstanceTemplate) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.NodeName != nil {
		in, out := &in.NodeName, &out.NodeName
		*out = new(string)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]ClusterComponentVolumeClaimTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceTemplate.
func (in *InstanceTemplate) DeepCopy() *InstanceTemplate {
	if in == nil {
		return nil
	}
	out := new(InstanceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Issuer) DeepCopyInto(out *Issuer) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(TLSSecretRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Issuer.
func (in *Issuer) DeepCopy() *Issuer {
	if in == nil {
		return nil
	}
	out := new(Issuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberStatus) DeepCopyInto(out *MemberStatus) {
	*out = *in
	if in.ReplicaRole != nil {
		in, out := &in.ReplicaRole, &out.ReplicaRole
		*out = new(ReplicaRole)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberStatus.
func (in *MemberStatus) DeepCopy() *MemberStatus {
	if in == nil {
		return nil
	}
	out := new(MemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaimSpec) DeepCopyInto(out *PersistentVolumeClaimSpec) {
	*out = *in
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]corev1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.VolumeMode != nil {
		in, out := &in.VolumeMode, &out.VolumeMode
		*out = new(corev1.PersistentVolumeMode)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaimSpec.
func (in *PersistentVolumeClaimSpec) DeepCopy() *PersistentVolumeClaimSpec {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaRole) DeepCopyInto(out *ReplicaRole) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaRole.
func (in *ReplicaRole) DeepCopy() *ReplicaRole {
	if in == nil {
		return nil
	}
	out := new(ReplicaRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMeta) DeepCopyInto(out *ResourceMeta) {
	*out = *in
	if in.AsVolumeFrom != nil {
		in, out := &in.AsVolumeFrom, &out.AsVolumeFrom
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMeta.
func (in *ResourceMeta) DeepCopy() *ResourceMeta {
	if in == nil {
		return nil
	}
	out := new(ResourceMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicy) DeepCopyInto(out *SchedulingPolicy) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]corev1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicy.
func (in *SchedulingPolicy) DeepCopy() *SchedulingPolicy {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
	in.ResourceMeta.DeepCopyInto(&out.ResourceMeta)
	in.Secret.DeepCopyInto(&out.Secret)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceRef) DeepCopyInto(out *ServiceRef) {
	*out = *in
	if in.ClusterServiceSelector != nil {
		in, out := &in.ClusterServiceSelector, &out.ClusterServiceSelector
		*out = new(ServiceRefClusterSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceRef.
func (in *ServiceRef) DeepCopy() *ServiceRef {
	if in == nil {
		return nil
	}
	out := new(ServiceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceRefClusterSelector) DeepCopyInto(out *ServiceRefClusterSelector) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ServiceRefServiceSelector)
		**out = **in
	}
	if in.Credential != nil {
		in, out := &in.Credential, &out.Credential
		*out = new(ServiceRefCredentialSelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceRefClusterSelector.
func (in *ServiceRefClusterSelector) DeepCopy() *ServiceRefClusterSelector {
	if in == nil {
		return nil
	}
	out := new(ServiceRefClusterSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceRefCredentialSelector) DeepCopyInto(out *ServiceRefCredentialSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceRefCredentialSelector.
func (in *ServiceRefCredentialSelector) DeepCopy() *ServiceRefCredentialSelector {
	if in == nil {
		return nil
	}
	out := new(ServiceRefCredentialSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceRefServiceSelector) DeepCopyInto(out *ServiceRefServiceSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceRefServiceSelector.
func (in *ServiceRefServiceSelector) DeepCopy() *ServiceRefServiceSelector {
	if in == nil {
		return nil
	}
	out := new(ServiceRefServiceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSSecretRef) DeepCopyInto(out *TLSSecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSSecretRef.
func (in *TLSSecretRef) DeepCopy() *TLSSecretRef {
	if in == nil {
		return nil
	}
	out := new(TLSSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserResourceRefs) DeepCopyInto(out *UserResourceRefs) {
	*out = *in
	if in.SecretRefs != nil {
		in, out := &in.SecretRefs, &out.SecretRefs
		*out = make([]SecretRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConfigMapRefs != nil {
		in, out := &in.ConfigMapRefs, &out.ConfigMapRefs
		*out = make([]ConfigMapRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserResourceRefs.
func (in *UserResourceRefs) DeepCopy() *UserResourceRefs {
	if in == nil {
		return nil
	}
	out := new(UserResourceRefs)
	in.DeepCopyInto(out)
	return out
}
