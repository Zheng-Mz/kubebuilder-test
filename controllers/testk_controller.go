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
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	zmzappv1 "github.com/Zheng-Mz/kubebuilder-test/api/v1"
	appsv1 "k8s.io/api/apps/v1"
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
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

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

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      make(map[string]string),
			Annotations: make(map[string]string),
			Name:        appRs.Name,
			Namespace:   appRs.Namespace,
		},
		Spec: *appRs.Spec.DeploySpec.DeepCopy(),
	}

	controllerutil.SetControllerReference(appRs, dep, r.Scheme)
	var as appsv1.Deployment
	if err := r.Get(context.Background(), types.NamespacedName{Namespace: dep.Namespace, Name: dep.Name}, &as); err != nil {
		if errors.IsNotFound(err) {
			log.Info("Creating a new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
			if err = r.Create(context.Background(), dep); err != nil {
				return ctrl.Result{}, fmt.Errorf("could not Create Deployment: %v", err)
			}
		} else {
			log.Error(err, "unable to fetch Deployment")
			return ctrl.Result{}, fmt.Errorf("could not Create Deployment: %v", err)
		}
	}

	return ctrl.Result{}, nil
}

func (r *TestKReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&zmzappv1.TestK{}).
		Owns(&appsv1.Deployment{}).
		Complete(r)
}
