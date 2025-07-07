package controller

import (
	"context"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	cozyv1alpha1 "github.com/cozystack/cozystack/api/v1alpha1"
)

type CozystackReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *CozystackReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	c := &cozyv1alpha1.Cozystack{}

	if err := r.Get(ctx, req.NamespacedName, c); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		l.Error(err, "Unable to fetch Cozystack")
		return ctrl.Result{}, err
	}

	if !c.DeletionTimestamp.IsZero() {
		return ctrl.Result{}, nil
	}

	if c.Name != "cozystack" || c.Namespace != "cozy-system" {
		l.Info("only cozy-system/cozystack Cozystack is allowed in a cluster")
		err := r.Delete(ctx, c)
		if err != nil {
			l.Error(err, "Unable to delete invalid Cozystack")
		}
		return ctrl.Result{}, err
	}

	panic("not implemented")
}

func (r *CozystackReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cozyv1alpha1.Cozystack{}).
		Complete(r)
}
