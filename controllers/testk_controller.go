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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

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

func (r *TestKReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	log := r.Log.WithValues("testk", req.NamespacedName)

	// Fetch the ReplicaSet from the cache
	appRs := &zmzappv1.TestK{}
	err := r.Get(context.TODO(), req.NamespacedName, appRs)
	if errors.IsNotFound(err) {
		log.Error(nil, "Could not find ReplicaSet")
		return ctrl.Result{}, nil
	}

	// Using a typed object.
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: req.NamespacedName.Namespace,
			Name:      "test-pod",
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
	// r is a created client.
	err = r.Create(context.Background(), pod)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("could not Create Pod: %+v", err)
	}

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

func (r *TestKReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&zmzappv1.TestK{}).
		Complete(r)
}
