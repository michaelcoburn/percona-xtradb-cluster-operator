// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	metav1 "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppStatus) DeepCopyInto(out *AppStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppStatus.
func (in *AppStatus) DeepCopy() *AppStatus {
	if in == nil {
		return nil
	}
	out := new(AppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStorageS3Spec) DeepCopyInto(out *BackupStorageS3Spec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageS3Spec.
func (in *BackupStorageS3Spec) DeepCopy() *BackupStorageS3Spec {
	if in == nil {
		return nil
	}
	out := new(BackupStorageS3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStorageSpec) DeepCopyInto(out *BackupStorageSpec) {
	*out = *in
	out.S3 = in.S3
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(VolumeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(PodResources)
		(*in).DeepCopyInto(*out)
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
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(corev1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageSpec.
func (in *BackupStorageSpec) DeepCopy() *BackupStorageSpec {
	if in == nil {
		return nil
	}
	out := new(BackupStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCondition) DeepCopyInto(out *ClusterCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCondition.
func (in *ClusterCondition) DeepCopy() *ClusterCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogCollectorSpec) DeepCopyInto(out *LogCollectorSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(PodResources)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogCollectorSpec.
func (in *LogCollectorSpec) DeepCopy() *LogCollectorSpec {
	if in == nil {
		return nil
	}
	out := new(LogCollectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PITR) DeepCopyInto(out *PITR) {
	*out = *in
	if in.BackupSource != nil {
		in, out := &in.BackupSource, &out.BackupSource
		*out = new(PXCBackupStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PITR.
func (in *PITR) DeepCopy() *PITR {
	if in == nil {
		return nil
	}
	out := new(PITR)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PITRSpec) DeepCopyInto(out *PITRSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(PodResources)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PITRSpec.
func (in *PITRSpec) DeepCopy() *PITRSpec {
	if in == nil {
		return nil
	}
	out := new(PITRSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PMMSpec) DeepCopyInto(out *PMMSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(PodResources)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PMMSpec.
func (in *PMMSpec) DeepCopy() *PMMSpec {
	if in == nil {
		return nil
	}
	out := new(PMMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PXCBackupSpec) DeepCopyInto(out *PXCBackupSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PXCBackupSpec.
func (in *PXCBackupSpec) DeepCopy() *PXCBackupSpec {
	if in == nil {
		return nil
	}
	out := new(PXCBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PXCBackupStatus) DeepCopyInto(out *PXCBackupStatus) {
	*out = *in
	if in.CompletedAt != nil {
		in, out := &in.CompletedAt, &out.CompletedAt
		*out = (*in).DeepCopy()
	}
	if in.LastScheduled != nil {
		in, out := &in.LastScheduled, &out.LastScheduled
		*out = (*in).DeepCopy()
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(BackupStorageS3Spec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PXCBackupStatus.
func (in *PXCBackupStatus) DeepCopy() *PXCBackupStatus {
	if in == nil {
		return nil
	}
	out := new(PXCBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PXCScheduledBackup) DeepCopyInto(out *PXCScheduledBackup) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = make([]PXCScheduledBackupSchedule, len(*in))
		copy(*out, *in)
	}
	if in.Storages != nil {
		in, out := &in.Storages, &out.Storages
		*out = make(map[string]*BackupStorageSpec, len(*in))
		for key, val := range *in {
			var outVal *BackupStorageSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(BackupStorageSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.PITR.DeepCopyInto(&out.PITR)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PXCScheduledBackup.
func (in *PXCScheduledBackup) DeepCopy() *PXCScheduledBackup {
	if in == nil {
		return nil
	}
	out := new(PXCScheduledBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PXCScheduledBackupSchedule) DeepCopyInto(out *PXCScheduledBackupSchedule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PXCScheduledBackupSchedule.
func (in *PXCScheduledBackupSchedule) DeepCopy() *PXCScheduledBackupSchedule {
	if in == nil {
		return nil
	}
	out := new(PXCScheduledBackupSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PXCSpec) DeepCopyInto(out *PXCSpec) {
	*out = *in
	if in.AutoRecovery != nil {
		in, out := &in.AutoRecovery, &out.AutoRecovery
		*out = new(bool)
		**out = **in
	}
	if in.ReplicationChannels != nil {
		in, out := &in.ReplicationChannels, &out.ReplicationChannels
		*out = make([]ReplicationChannel, len(*in))
		copy(*out, *in)
	}
	in.Expose.DeepCopyInto(&out.Expose)
	if in.PodSpec != nil {
		in, out := &in.PodSpec, &out.PodSpec
		*out = new(PodSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PXCSpec.
func (in *PXCSpec) DeepCopy() *PXCSpec {
	if in == nil {
		return nil
	}
	out := new(PXCSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBCluster) DeepCopyInto(out *PerconaXtraDBCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBCluster.
func (in *PerconaXtraDBCluster) DeepCopy() *PerconaXtraDBCluster {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaXtraDBCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBClusterBackup) DeepCopyInto(out *PerconaXtraDBClusterBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBClusterBackup.
func (in *PerconaXtraDBClusterBackup) DeepCopy() *PerconaXtraDBClusterBackup {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBClusterBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaXtraDBClusterBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBClusterBackupList) DeepCopyInto(out *PerconaXtraDBClusterBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaXtraDBClusterBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBClusterBackupList.
func (in *PerconaXtraDBClusterBackupList) DeepCopy() *PerconaXtraDBClusterBackupList {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBClusterBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaXtraDBClusterBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBClusterList) DeepCopyInto(out *PerconaXtraDBClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaXtraDBCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBClusterList.
func (in *PerconaXtraDBClusterList) DeepCopy() *PerconaXtraDBClusterList {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaXtraDBClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBClusterRestore) DeepCopyInto(out *PerconaXtraDBClusterRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBClusterRestore.
func (in *PerconaXtraDBClusterRestore) DeepCopy() *PerconaXtraDBClusterRestore {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBClusterRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaXtraDBClusterRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBClusterRestoreList) DeepCopyInto(out *PerconaXtraDBClusterRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaXtraDBClusterRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBClusterRestoreList.
func (in *PerconaXtraDBClusterRestoreList) DeepCopy() *PerconaXtraDBClusterRestoreList {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBClusterRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaXtraDBClusterRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBClusterRestoreSpec) DeepCopyInto(out *PerconaXtraDBClusterRestoreSpec) {
	*out = *in
	if in.BackupSource != nil {
		in, out := &in.BackupSource, &out.BackupSource
		*out = new(PXCBackupStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.PITR != nil {
		in, out := &in.PITR, &out.PITR
		*out = new(PITR)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBClusterRestoreSpec.
func (in *PerconaXtraDBClusterRestoreSpec) DeepCopy() *PerconaXtraDBClusterRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBClusterRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBClusterRestoreStatus) DeepCopyInto(out *PerconaXtraDBClusterRestoreStatus) {
	*out = *in
	if in.CompletedAt != nil {
		in, out := &in.CompletedAt, &out.CompletedAt
		*out = (*in).DeepCopy()
	}
	if in.LastScheduled != nil {
		in, out := &in.LastScheduled, &out.LastScheduled
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBClusterRestoreStatus.
func (in *PerconaXtraDBClusterRestoreStatus) DeepCopy() *PerconaXtraDBClusterRestoreStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBClusterRestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBClusterSpec) DeepCopyInto(out *PerconaXtraDBClusterSpec) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PXC != nil {
		in, out := &in.PXC, &out.PXC
		*out = new(PXCSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ProxySQL != nil {
		in, out := &in.ProxySQL, &out.ProxySQL
		*out = new(PodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.HAProxy != nil {
		in, out := &in.HAProxy, &out.HAProxy
		*out = new(PodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PMM != nil {
		in, out := &in.PMM, &out.PMM
		*out = new(PMMSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.LogCollector != nil {
		in, out := &in.LogCollector, &out.LogCollector
		*out = new(LogCollectorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(PXCScheduledBackup)
		(*in).DeepCopyInto(*out)
	}
	out.UpgradeOptions = in.UpgradeOptions
	if in.EnableCRValidationWebhook != nil {
		in, out := &in.EnableCRValidationWebhook, &out.EnableCRValidationWebhook
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBClusterSpec.
func (in *PerconaXtraDBClusterSpec) DeepCopy() *PerconaXtraDBClusterSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBClusterStatus) DeepCopyInto(out *PerconaXtraDBClusterStatus) {
	*out = *in
	out.PXC = in.PXC
	out.ProxySQL = in.ProxySQL
	out.HAProxy = in.HAProxy
	out.Backup = in.Backup
	out.PMM = in.PMM
	out.LogCollector = in.LogCollector
	if in.Messages != nil {
		in, out := &in.Messages, &out.Messages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBClusterStatus.
func (in *PerconaXtraDBClusterStatus) DeepCopy() *PerconaXtraDBClusterStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodAffinity) DeepCopyInto(out *PodAffinity) {
	*out = *in
	if in.TopologyKey != nil {
		in, out := &in.TopologyKey, &out.TopologyKey
		*out = new(string)
		**out = **in
	}
	if in.Advanced != nil {
		in, out := &in.Advanced, &out.Advanced
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodAffinity.
func (in *PodAffinity) DeepCopy() *PodAffinity {
	if in == nil {
		return nil
	}
	out := new(PodAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetSpec) DeepCopyInto(out *PodDisruptionBudgetSpec) {
	*out = *in
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetSpec.
func (in *PodDisruptionBudgetSpec) DeepCopy() *PodDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodResources) DeepCopyInto(out *PodResources) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(ResourcesList)
		**out = **in
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = new(ResourcesList)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodResources.
func (in *PodResources) DeepCopy() *PodResources {
	if in == nil {
		return nil
	}
	out := new(PodResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSpec) DeepCopyInto(out *PodSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(PodResources)
		(*in).DeepCopyInto(*out)
	}
	if in.SidecarResources != nil {
		in, out := &in.SidecarResources, &out.SidecarResources
		*out = new(PodResources)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeSpec != nil {
		in, out := &in.VolumeSpec, &out.VolumeSpec
		*out = new(VolumeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(PodAffinity)
		(*in).DeepCopyInto(*out)
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
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.TerminationGracePeriodSeconds != nil {
		in, out := &in.TerminationGracePeriodSeconds, &out.TerminationGracePeriodSeconds
		*out = new(int64)
		**out = **in
	}
	if in.LoadBalancerSourceRanges != nil {
		in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceAnnotations != nil {
		in, out := &in.ServiceAnnotations, &out.ServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReadinessInitialDelaySeconds != nil {
		in, out := &in.ReadinessInitialDelaySeconds, &out.ReadinessInitialDelaySeconds
		*out = new(int32)
		**out = **in
	}
	if in.LivenessInitialDelaySeconds != nil {
		in, out := &in.LivenessInitialDelaySeconds, &out.LivenessInitialDelaySeconds
		*out = new(int32)
		**out = **in
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(corev1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make([]corev1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSpec.
func (in *PodSpec) DeepCopy() *PodSpec {
	if in == nil {
		return nil
	}
	out := new(PodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationChannel) DeepCopyInto(out *ReplicationChannel) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationChannel.
func (in *ReplicationChannel) DeepCopy() *ReplicationChannel {
	if in == nil {
		return nil
	}
	out := new(ReplicationChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesList) DeepCopyInto(out *ResourcesList) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesList.
func (in *ResourcesList) DeepCopy() *ResourcesList {
	if in == nil {
		return nil
	}
	out := new(ResourcesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExpose) DeepCopyInto(out *ServiceExpose) {
	*out = *in
	if in.LoadBalancerSourceRanges != nil {
		in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExpose.
func (in *ServiceExpose) DeepCopy() *ServiceExpose {
	if in == nil {
		return nil
	}
	out := new(ServiceExpose)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSSpec) DeepCopyInto(out *TLSSpec) {
	*out = *in
	if in.SANs != nil {
		in, out := &in.SANs, &out.SANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IssuerConf != nil {
		in, out := &in.IssuerConf, &out.IssuerConf
		*out = new(metav1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSSpec.
func (in *TLSSpec) DeepCopy() *TLSSpec {
	if in == nil {
		return nil
	}
	out := new(TLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeOptions) DeepCopyInto(out *UpgradeOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeOptions.
func (in *UpgradeOptions) DeepCopy() *UpgradeOptions {
	if in == nil {
		return nil
	}
	out := new(UpgradeOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	if in.PVCs != nil {
		in, out := &in.PVCs, &out.PVCs
		*out = make([]corev1.PersistentVolumeClaim, len(*in))
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
	if in.EmptyDir != nil {
		in, out := &in.EmptyDir, &out.EmptyDir
		*out = new(corev1.EmptyDirVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.HostPath != nil {
		in, out := &in.HostPath, &out.HostPath
		*out = new(corev1.HostPathVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(corev1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSpec.
func (in *VolumeSpec) DeepCopy() *VolumeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeSpec)
	in.DeepCopyInto(out)
	return out
}
