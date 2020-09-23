/*
Copyright 2020 zmz.

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

package controllers

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	//"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	//"sigs.k8s.io/controller-runtime/pkg/source"

	zmzappv1 "github.com/Zheng-Mz/kubebuilder-test/api/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TestKReconciler reconciles a TestK object
type TestKReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=zmzapp.zmz.example.org,resources=testks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=zmzapp.zmz.example.org,resources=testks/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=pobs,verbs=get;list;watch;create;update;patch;delete

func (r *TestKReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	log := r.Log.WithValues("testk", req.NamespacedName)

	log.Info("--- RECONCILE TestK ---", "Req ", req)

	// Fetch the ReplicaSet from the cache
	appRs := &zmzappv1.TestK{}
	err := r.Get(context.TODO(), req.NamespacedName, appRs)
	if errors.IsNotFound(err) {
		log.Error(nil, "Could not find ReplicaSet")
		return ctrl.Result{}, nil
	}
	log.Info("A new TestK", "TestK.Namespace", appRs.Namespace, "TestK.Name", appRs.Name)

	// Using a typed object.
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: appRs.Namespace,
			Name:      appRs.Name,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				corev1.Container{
					Image: "nginx",
					Name:  "nginx",
				},
			},
		},
	}

	log.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
	if err = r.Create(context.Background(), pod); err != nil {
		return ctrl.Result{}, fmt.Errorf("could not Create Pod: %v", err)
	}
	/*
		controllerutil.SetControllerReference(appRs, pod, r.Scheme)
		var as corev1.Pod
		if err := r.Get(context.Background(), types.NamespacedName{Namespace: pod.Namespace, Name: pod.Name}, &as); err != nil {
			if errors.IsNotFound(err) {
				log.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
				if err = r.Create(context.Background(), pod); err != nil {
					return ctrl.Result{}, fmt.Errorf("could not Create Pod: %v", err)
				}
			} else {
				log.Error(err, "unable to fetch Pod")
				return ctrl.Result{}, fmt.Errorf("could not Create Pod: %v", err)
			}
		}*/
	/*
		labels := make(map[string]string)
		labels["app"] = "123456"
		meta := metav1.ObjectMeta{
			Name:      "name-svc",
			Namespace: appRs.Namespace,
			Labels:    labels,
		}

		svc := &corev1.Service{
			ObjectMeta: meta,
			Spec: corev1.ServiceSpec{
				Type: "ClusterIP",
				Ports: []corev1.ServicePort{
					corev1.ServicePort{
						Port: 895,
					},
				},
				Selector: labels,
			},
		}
		controllerutil.SetControllerReference(appRs, svc, r.Scheme)
		var sf corev1.Service
		if err := r.Get(context.Background(), types.NamespacedName{Namespace: svc.Namespace, Name: svc.Name}, &sf); err != nil {
			if errors.IsNotFound(err) {
				log.Info("Creating a new service", "Service.Namespace", svc.Namespace, "Service.Name", svc.Name)
				if err = r.Create(context.Background(), svc); err != nil {
					return ctrl.Result{}, err
				}
			} else {
				log.Error(err, "Unable to fetch Service")
				return ctrl.Result{}, err
			}
		}*/
	/*
		// create a ReplicaSet
		rs := &appsv1.ReplicaSet{}
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("could not fetch ReplicaSet: %+v", err)
		}

		// Print the ReplicaSet
		log.Info("Reconciling ReplicaSet", "container name", rs.Spec.Template.Spec.Containers[0].Name)

		// Set the label if it is missing
		if rs.Labels == nil {
			rs.Labels = map[string]string{}
		}
		if rs.Labels["hello"] == "world" {
			return ctrl.Result{}, nil
		}

		// Update the ReplicaSet
		rs.Labels["hello"] = "world"
		err = r.Update(context.TODO(), rs)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("could not write ReplicaSet: %+v", err)
		}

	*/

	return ctrl.Result{}, nil
}

func (r *TestKReconciler) Map(obj handler.MapObject) []reconcile.Request {
	var reqs []reconcile.Request
	//var nf axyomv1.NetworkFunctionSpec
	var name string
	objType := "nf"
	switch v := obj.Object.(type) {
	case *corev1.Service:
		name = v.Name
		objType = "svc"
	case *corev1.Pod:
		name = v.Name
		objType = "pod"
	}

	var testKList zmzappv1.TestKList
	r.List(context.Background(), &testKList)
	for _, a := range testKList.Items {
		switch objType {
		case "svc":
			substr := []string{a.Name + "-testK", a.Name + "-n2"}
			for _, v := range substr {
				if strings.Contains(name, v) {
					reqs = append(reqs, reconcile.Request{NamespacedName: types.NamespacedName{Name: a.Name, Namespace: a.Namespace}})
				}
			}
		case "pod":
			substr := []string{a.Name + "-testK", a.Name + "-n2"}
			for _, v := range substr {
				if strings.Contains(name, v) {
					reqs = append(reqs, reconcile.Request{NamespacedName: types.NamespacedName{Name: a.Name, Namespace: a.Namespace}})
				}
			}
		}
	}

	return reqs
}

func (r *TestKReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&zmzappv1.TestK{}).
		//Owns(&corev1.Service{}).
		//Owns(&corev1.Pod{}).
		//Watches(&source.Kind{Type: &corev1.Service{}}, &handler.EnqueueRequestsFromMapFunc{ToRequests: r}).
		//Watches(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestsFromMapFunc{ToRequests: r}).
		Complete(r)
}
