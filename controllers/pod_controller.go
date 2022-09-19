package controllers

import (
  "context"

  "github.com/go-logr/logr"
  corev1 "k8s.io/api/core/v1"
  apierrors "k8s.io/apimachinery/pkg/api/errors"
  "k8s.io/apimachinery/pkg/runtime"
  ctrl "sigs.k8s.io/controller-runtime"
  "sigs.k8s.io/controller-runtime/pkg/client"
  "github.com/neilharris123/metamirror/config"
)

// PodReconciler reconciles a Pod object
type PodReconciler struct {
client.Client
  Log    logr.Logger
  Scheme *runtime.Scheme
}

func (r *PodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
  log := r.Log.WithValues("pod", req.NamespacedName)

  /*
     Step 0: Fetch the Pod from the Kubernetes API.
  */

  var pod corev1.Pod
  if err := r.Get(ctx, req.NamespacedName, &pod); err != nil {
    if apierrors.IsNotFound(err) {
      // we'll ignore not-found errors, since we can get them on deleted requests.
      return ctrl.Result{}, nil
    }
    log.Error(err, "unable to fetch Pod")
    return ctrl.Result{}, err
  }

  /*
     Step 1a: Ensure annotation and label lists are of equal length
  */
  reqAnnotations := config.Metadata.Annotations
  reqLabels := config.Metadata.Labels
  if len(reqAnnotations) != len(reqLabels) {
    log.Error(err, "Illegal config, variable lists are of unequal length")
  }
  /*
     Step 1: Add the label if the annotation exists, but the label does not
  */

  _, targetAnnotation := pod.Annotations[config.Metadata.Annotation]
  targetLabel := pod.Labels[config.Metadata.Label] == pod.Annotations[config.Metadata.Annotation]


  if targetAnnotation == targetLabel {
    // The desired state and actual state of the Pod are the same.
    // No further action is required by the operator at this moment.
    log.Info("no update required")
    return ctrl.Result{}, nil
  }

  if targetAnnotation {
    // If the label should be set but is not, set it.
    if pod.Labels == nil {
      pod.Labels = make(map[string]string)
    }
    pod.Labels[config.Metadata.Label] = pod.Annotations[config.Metadata.Annotation]
    log.Info("adding label")
  }

  /*
     Step 2: Push the updated pod to the Kubernetes API.
  */

  if err := r.Update(ctx, &pod); err != nil {
    if apierrors.IsConflict(err) {
      // If the Pod has been by another process since we read it.
      // Requeue the Pod to try to reconciliate again.
      return ctrl.Result{Requeue: true}, nil
    }
    if apierrors.IsNotFound(err) {
      // The Pod has been deleted since we read it.
      // Requeue the Pod to try to reconciliate again.
      return ctrl.Result{Requeue: true}, nil
    }
    log.Error(err, "unable to update Pod")
    return ctrl.Result{}, err
  }
  return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PodReconciler) SetupWithManager(mgr ctrl.Manager) error {
  return ctrl.NewControllerManagedBy(mgr).
    For(&corev1.Pod{}).
    Complete(r)
}
