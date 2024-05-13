package sugar

import (
	sugar "github.com/apecloud/kubeblocks-sugar/api/sugar/v1alpha1"
	appsv1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	"github.com/apecloud/kubeblocks/pkg/controller/builder"
	"github.com/apecloud/kubeblocks/pkg/controller/kubebuilderx"
	"github.com/apecloud/kubeblocks/pkg/controller/model"
)

type translateToKubeBlocksClusterReconciler struct{}

func (r *translateToKubeBlocksClusterReconciler) PreCondition(tree *kubebuilderx.ObjectTree) *kubebuilderx.CheckResult {
	if tree.GetRoot() == nil || model.IsObjectDeleting(tree.GetRoot()) {
		return kubebuilderx.ResultUnsatisfied
	}
	return kubebuilderx.ResultSatisfied
}

func (r *translateToKubeBlocksClusterReconciler) Reconcile(tree *kubebuilderx.ObjectTree) (*kubebuilderx.ObjectTree, error) {
	acCluster, _ := tree.GetRoot().(*sugar.ApeCloudMySQL)
	kbCluster := builder.NewClusterBuilder(acCluster.Namespace, acCluster.Name).GetObject()
	object, err := tree.Get(kbCluster)
	if err != nil {
		return nil, err
	}
	if object != nil {
		kbCluster, _ = object.(*appsv1alpha1.Cluster)
	}
	kbCluster.Spec = *(&acCluster.Spec).TranslateTo()
	if err = tree.Update(kbCluster); err != nil {
		return nil, err
	}
	return tree, nil
}

func translateToKubeBlocksCluster() kubebuilderx.Reconciler {
	return &translateToKubeBlocksClusterReconciler{}
}

var _ kubebuilderx.Reconciler = &translateToKubeBlocksClusterReconciler{}
