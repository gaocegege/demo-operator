/*

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

	"github.com/caicloud/mlneuron-controller/pkg/util"
	"github.com/go-logr/logr"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	demov1alpha1 "github.com/gaocegege/demo-operator/api/v1alpha1"
)

// DemoJobReconciler reconciles a DemoJob object
type DemoJobReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=demo.k8s.io,resources=demojobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=demo.k8s.io,resources=demojobs/status,verbs=get;update;patch

func (r *DemoJobReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()

	r.Log.Info("instance", "req", req)

	// Fetch the Serving instance
	instance := &demov1alpha1.DemoJob{}
	err := r.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name,
			Namespace: instance.Namespace,
		},
		Spec: batchv1.JobSpec{
			Completions:  util.NewInt32(1),
			BackoffLimit: util.NewInt32(0),
			// TTLSecondsAfterFinished: util.NewInt32(0),
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					RestartPolicy: corev1.RestartPolicyNever,
					Containers: []corev1.Container{
						{
							Name:  "resolver",
							Image: instance.Spec.Image,
						},
					},
				},
			},
		},
	}

	if err := controllerutil.SetControllerReference(
		instance, job, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	r.Log.Info("Creating new job", "job", job)
	if err := r.Create(context.Background(), job); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *DemoJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1alpha1.DemoJob{}).
		Complete(r)
}
