package staticpodcontroller

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticPodConfigStatus) DeepCopyInto(out *StaticPodConfigStatus) {
	*out = *in
	in.OperatorStatus.DeepCopyInto(&out.OperatorStatus)
	if in.TargetKubeletStates != nil {
		in, out := &in.TargetKubeletStates, &out.TargetKubeletStates
		*out = make([]KubeletState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeAPIServerOperatorConfigStatus.
func (in *StaticPodConfigStatus) DeepCopy() *StaticPodConfigStatus {
	if in == nil {
		return nil
	}
	out := new(StaticPodConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletState) DeepCopyInto(out *KubeletState) {
	*out = *in
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletState.
func (in *KubeletState) DeepCopy() *KubeletState {
	if in == nil {
		return nil
	}
	out := new(KubeletState)
	in.DeepCopyInto(out)
	return out
}
