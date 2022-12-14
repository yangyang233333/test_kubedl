//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Alibaba Authors.

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
// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&MPIJob{}, func(obj interface{}) { SetObjectDefaults_MPIJob(obj.(*MPIJob)) })
	scheme.AddTypeDefaultingFunc(&MPIJobList{}, func(obj interface{}) { SetObjectDefaults_MPIJobList(obj.(*MPIJobList)) })
	scheme.AddTypeDefaultingFunc(&MarsJob{}, func(obj interface{}) { SetObjectDefaults_MarsJob(obj.(*MarsJob)) })
	scheme.AddTypeDefaultingFunc(&MarsJobList{}, func(obj interface{}) { SetObjectDefaults_MarsJobList(obj.(*MarsJobList)) })
	scheme.AddTypeDefaultingFunc(&PyTorchJob{}, func(obj interface{}) { SetObjectDefaults_PyTorchJob(obj.(*PyTorchJob)) })
	scheme.AddTypeDefaultingFunc(&PyTorchJobList{}, func(obj interface{}) { SetObjectDefaults_PyTorchJobList(obj.(*PyTorchJobList)) })
	scheme.AddTypeDefaultingFunc(&TFJob{}, func(obj interface{}) { SetObjectDefaults_TFJob(obj.(*TFJob)) })
	scheme.AddTypeDefaultingFunc(&TFJobList{}, func(obj interface{}) { SetObjectDefaults_TFJobList(obj.(*TFJobList)) })
	scheme.AddTypeDefaultingFunc(&XDLJob{}, func(obj interface{}) { SetObjectDefaults_XDLJob(obj.(*XDLJob)) })
	scheme.AddTypeDefaultingFunc(&XDLJobList{}, func(obj interface{}) { SetObjectDefaults_XDLJobList(obj.(*XDLJobList)) })
	scheme.AddTypeDefaultingFunc(&XGBoostJob{}, func(obj interface{}) { SetObjectDefaults_XGBoostJob(obj.(*XGBoostJob)) })
	scheme.AddTypeDefaultingFunc(&XGBoostJobList{}, func(obj interface{}) { SetObjectDefaults_XGBoostJobList(obj.(*XGBoostJobList)) })
	return nil
}

func SetObjectDefaults_MPIJob(in *MPIJob) {
	SetDefaults_MPIJob(in)
}

func SetObjectDefaults_MPIJobList(in *MPIJobList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_MPIJob(a)
	}
}

func SetObjectDefaults_MarsJob(in *MarsJob) {
	SetDefaults_MarsJob(in)
}

func SetObjectDefaults_MarsJobList(in *MarsJobList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_MarsJob(a)
	}
}

func SetObjectDefaults_PyTorchJob(in *PyTorchJob) {
	SetDefaults_PyTorchJob(in)
}

func SetObjectDefaults_PyTorchJobList(in *PyTorchJobList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_PyTorchJob(a)
	}
}

func SetObjectDefaults_TFJob(in *TFJob) {
	SetDefaults_TFJob(in)
}

func SetObjectDefaults_TFJobList(in *TFJobList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_TFJob(a)
	}
}

func SetObjectDefaults_XDLJob(in *XDLJob) {
	SetDefaults_XDLJob(in)
}

func SetObjectDefaults_XDLJobList(in *XDLJobList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_XDLJob(a)
	}
}

func SetObjectDefaults_XGBoostJob(in *XGBoostJob) {
	SetDefaults_XGBoostJob(in)
}

func SetObjectDefaults_XGBoostJobList(in *XGBoostJobList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_XGBoostJob(a)
	}
}
