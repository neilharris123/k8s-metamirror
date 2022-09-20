package controllers

import (
  "context"
  "strings"
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
     Step 1: Ensure annotation and label lists are of equal length
  */
  reqAnnotations := strings.Split(config.Metadata.Annotations, ",")
  reqLabels := strings.Split(config.Metadata.Labels, ",")

  if len(reqAnnotations) != len(reqLabels) {
    panic("Illegal config, variable lists are of unequal length. Exiting")
  }

  /*
     Step 2: Add the label if the annotation exists, but the label does not
  */

  for i, arg := range reqAnnotations {
    _, targetAnnotation := pod.Annotations[string(arg)]
    targetLabel := pod.Labels[string(reqLabels[i])] == pod.Annotations[string(arg)]


    if targetAnnotation == targetLabel {
      log.Info("no update required")
      continue
    }

    // If the label should be set but is not, set it.
    if targetAnnotation {
      if pod.Labels == nil {
        pod.Labels = make(map[string]string)
      }
      pod.Labels[string(reqLabels[i])] = pod.Annotations[string(arg)]
      log.Info("adding label " + string(reqLabels[i]))
    }
  }

  /*
     Step 3: Push the updated pod to the Kubernetes API.
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
