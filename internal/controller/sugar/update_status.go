package sugar

import (
	sugar "github.com/apecloud/kubeblocks-sugar/api/sugar/v1alpha1"
	"github.com/apecloud/kubeblocks/pkg/controller/builder"
	"github.com/apecloud/kubeblocks/pkg/controller/kubebuilderx"
	"github.com/apecloud/kubeblocks/pkg/controller/model"
)

type updateStatusReconciler struct{}

func (r *updateStatusReconciler) PreCondition(tree *kubebuilderx.ObjectTree) *kubebuilderx.CheckResult {
	if tree.GetRoot() == nil || model.IsObjectDeleting(tree.GetRoot()) {
		return kubebuilderx.ResultUnsatisfied
	}
	return kubebuilderx.ResultSatisfied
}

func (r *updateStatusReconciler) Reconcile(tree *kubebuilderx.ObjectTree) (*kubebuilderx.ObjectTree, error) {
	acCluster, _ := tree.GetRoot().(*sugar.ApeCloudMySQL)
	kbCluster := builder.NewClusterBuilder(acCluster.Namespace, acCluster.Name).GetObject()
	object, err := tree.Get(kbCluster)
	if err != nil {
		return nil, err
	}
	if object == nil {
		return tree, nil
	}
	acCluster.Status.BaseStatus.ClusterStatus = kbCluster.Status
	return tree, nil
}

func updateStatus() kubebuilderx.Reconciler {
	return &updateStatusReconciler{}
}

var _ kubebuilderx.Reconciler = &updateStatusReconciler{}
