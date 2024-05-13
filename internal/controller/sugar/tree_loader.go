package sugar

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sugar "github.com/apecloud/kubeblocks-sugar/api/sugar/v1alpha1"
	appsv1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	"github.com/apecloud/kubeblocks/pkg/controller/kubebuilderx"
)

type treeLoader struct{}

func (t *treeLoader) Load(ctx context.Context, reader client.Reader, request ctrl.Request, recorder record.EventRecorder, logger logr.Logger) (*kubebuilderx.ObjectTree, error) {
	tree, err := kubebuilderx.ReadObjectTree[*sugar.ApeCloudMySQL](ctx, reader, request, nil)
	if err != nil {
		return nil, err
	}
	kbCluster := &appsv1alpha1.Cluster{}
	if err = reader.Get(ctx, request.NamespacedName, kbCluster); err != nil && !errors.IsNotFound(err) {
		return nil, err
	} else if err == nil {
		if err = tree.Add(kbCluster); err != nil {
			return nil, err
		}
	}
	tree.EventRecorder = recorder
	tree.Logger = logger
	return tree, nil
}

func objectTree() kubebuilderx.TreeLoader {
	return &treeLoader{}
}

var _ kubebuilderx.TreeLoader = &treeLoader{}
