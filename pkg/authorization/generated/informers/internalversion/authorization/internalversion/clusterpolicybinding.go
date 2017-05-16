// This file was automatically generated by informer-gen

package internalversion

import (
	api "github.com/openshift/origin/pkg/authorization/api"
	internalinterfaces "github.com/openshift/origin/pkg/authorization/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/authorization/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/authorization/generated/listers/authorization/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ClusterPolicyBindingInformer provides access to a shared informer and lister for
// ClusterPolicyBindings.
type ClusterPolicyBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ClusterPolicyBindingLister
}

type clusterPolicyBindingInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newClusterPolicyBindingInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Authorization().ClusterPolicyBindings().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Authorization().ClusterPolicyBindings().Watch(options)
			},
		},
		&api.ClusterPolicyBinding{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *clusterPolicyBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&api.ClusterPolicyBinding{}, newClusterPolicyBindingInformer)
}

func (f *clusterPolicyBindingInformer) Lister() internalversion.ClusterPolicyBindingLister {
	return internalversion.NewClusterPolicyBindingLister(f.Informer().GetIndexer())
}
