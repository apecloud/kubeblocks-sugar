/*
Copyright (C) 2022-2024 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package sugar

import (
	"context"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
)

type kbClusterHandler struct {
	client.Client
}

func (h *kbClusterHandler) Create(ctx context.Context, event event.CreateEvent, limitingInterface workqueue.RateLimitingInterface) {
	h.mapAndEnqueue(ctx, limitingInterface, event.Object)
}

func (h *kbClusterHandler) Update(ctx context.Context, event event.UpdateEvent, limitingInterface workqueue.RateLimitingInterface) {
	h.mapAndEnqueue(ctx, limitingInterface, event.ObjectNew)
}

func (h *kbClusterHandler) Delete(ctx context.Context, event event.DeleteEvent, limitingInterface workqueue.RateLimitingInterface) {
	h.mapAndEnqueue(ctx, limitingInterface, event.Object)
}

func (h *kbClusterHandler) Generic(ctx context.Context, event event.GenericEvent, limitingInterface workqueue.RateLimitingInterface) {
}

func (h *kbClusterHandler) mapAndEnqueue(ctx context.Context, q workqueue.RateLimitingInterface, object client.Object) {
	q.Add(ctrl.Request{NamespacedName: types.NamespacedName{Namespace: object.GetNamespace(), Name: object.GetName()}})
}

func newKBClusterHandler() handler.EventHandler {
	return &kbClusterHandler{}
}

var _ handler.EventHandler = &kbClusterHandler{}
