//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024 B-Harvest Corporation.
Copyright 2022 Strangelove Ventures LLC.

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
	"github.com/kubernetes-csi/external-snapshotter/client/v6/apis/volumesnapshot/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTemplateSpec) DeepCopyInto(out *JobTemplateSpec) {
	*out = *in
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	if in.BackoffLimit != nil {
		in, out := &in.BackoffLimit, &out.BackoffLimit
		*out = new(int32)
		**out = **in
	}
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTemplateSpec.
func (in *JobTemplateSpec) DeepCopy() *JobTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(JobTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalFullNodeRef) DeepCopyInto(out *LocalFullNodeRef) {
	*out = *in
	if in.Ordinal != nil {
		in, out := &in.Ordinal, &out.Ordinal
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalFullNodeRef.
func (in *LocalFullNodeRef) DeepCopy() *LocalFullNodeRef {
	if in == nil {
		return nil
	}
	out := new(LocalFullNodeRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledVolumeSnapshot) DeepCopyInto(out *ScheduledVolumeSnapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledVolumeSnapshot.
func (in *ScheduledVolumeSnapshot) DeepCopy() *ScheduledVolumeSnapshot {
	if in == nil {
		return nil
	}
	out := new(ScheduledVolumeSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledVolumeSnapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledVolumeSnapshotList) DeepCopyInto(out *ScheduledVolumeSnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScheduledVolumeSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledVolumeSnapshotList.
func (in *ScheduledVolumeSnapshotList) DeepCopy() *ScheduledVolumeSnapshotList {
	if in == nil {
		return nil
	}
	out := new(ScheduledVolumeSnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledVolumeSnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledVolumeSnapshotSpec) DeepCopyInto(out *ScheduledVolumeSnapshotSpec) {
	*out = *in
	in.FullNodeRef.DeepCopyInto(&out.FullNodeRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledVolumeSnapshotSpec.
func (in *ScheduledVolumeSnapshotSpec) DeepCopy() *ScheduledVolumeSnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(ScheduledVolumeSnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledVolumeSnapshotStatus) DeepCopyInto(out *ScheduledVolumeSnapshotStatus) {
	*out = *in
	if in.StatusMessage != nil {
		in, out := &in.StatusMessage, &out.StatusMessage
		*out = new(string)
		**out = **in
	}
	in.CreatedAt.DeepCopyInto(&out.CreatedAt)
	if in.Candidate != nil {
		in, out := &in.Candidate, &out.Candidate
		*out = new(SnapshotCandidate)
		(*in).DeepCopyInto(*out)
	}
	if in.LastSnapshot != nil {
		in, out := &in.LastSnapshot, &out.LastSnapshot
		*out = new(VolumeSnapshotStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledVolumeSnapshotStatus.
func (in *ScheduledVolumeSnapshotStatus) DeepCopy() *ScheduledVolumeSnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(ScheduledVolumeSnapshotStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotCandidate) DeepCopyInto(out *SnapshotCandidate) {
	*out = *in
	if in.PodLabels != nil {
		in, out := &in.PodLabels, &out.PodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotCandidate.
func (in *SnapshotCandidate) DeepCopy() *SnapshotCandidate {
	if in == nil {
		return nil
	}
	out := new(SnapshotCandidate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulJob) DeepCopyInto(out *StatefulJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulJob.
func (in *StatefulJob) DeepCopy() *StatefulJob {
	if in == nil {
		return nil
	}
	out := new(StatefulJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatefulJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulJobList) DeepCopyInto(out *StatefulJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StatefulJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulJobList.
func (in *StatefulJobList) DeepCopy() *StatefulJobList {
	if in == nil {
		return nil
	}
	out := new(StatefulJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatefulJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulJobSpec) DeepCopyInto(out *StatefulJobSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Interval = in.Interval
	in.JobTemplate.DeepCopyInto(&out.JobTemplate)
	in.PodTemplate.DeepCopyInto(&out.PodTemplate)
	in.VolumeClaimTemplate.DeepCopyInto(&out.VolumeClaimTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulJobSpec.
func (in *StatefulJobSpec) DeepCopy() *StatefulJobSpec {
	if in == nil {
		return nil
	}
	out := new(StatefulJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulJobStatus) DeepCopyInto(out *StatefulJobStatus) {
	*out = *in
	if in.StatusMessage != nil {
		in, out := &in.StatusMessage, &out.StatusMessage
		*out = new(string)
		**out = **in
	}
	if in.JobHistory != nil {
		in, out := &in.JobHistory, &out.JobHistory
		*out = make([]batchv1.JobStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulJobStatus.
func (in *StatefulJobStatus) DeepCopy() *StatefulJobStatus {
	if in == nil {
		return nil
	}
	out := new(StatefulJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulJobVolumeClaimTemplate) DeepCopyInto(out *StatefulJobVolumeClaimTemplate) {
	*out = *in
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]corev1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulJobVolumeClaimTemplate.
func (in *StatefulJobVolumeClaimTemplate) DeepCopy() *StatefulJobVolumeClaimTemplate {
	if in == nil {
		return nil
	}
	out := new(StatefulJobVolumeClaimTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSnapshotStatus) DeepCopyInto(out *VolumeSnapshotStatus) {
	*out = *in
	in.StartedAt.DeepCopyInto(&out.StartedAt)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(v1.VolumeSnapshotStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSnapshotStatus.
func (in *VolumeSnapshotStatus) DeepCopy() *VolumeSnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(VolumeSnapshotStatus)
	in.DeepCopyInto(out)
	return out
}
