// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	cagipv1 "github.com/ca-gip/kubi-members/pkg/apis/ca-gip/v1"
	versioned "github.com/ca-gip/kubi-members/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/ca-gip/kubi-members/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/ca-gip/kubi-members/pkg/generated/listers/ca-gip/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ProjectMembersInformer provides access to a shared informer and lister for
// ProjectMemberses.
type ProjectMembersInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ProjectMembersLister
}

type projectMembersInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewProjectMembersInformer constructs a new informer for ProjectMembers type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewProjectMembersInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredProjectMembersInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredProjectMembersInformer constructs a new informer for ProjectMembers type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredProjectMembersInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CagipV1().ProjectMemberses(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CagipV1().ProjectMemberses(namespace).Watch(options)
			},
		},
		&cagipv1.ProjectMembers{},
		resyncPeriod,
		indexers,
	)
}

func (f *projectMembersInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredProjectMembersInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *projectMembersInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cagipv1.ProjectMembers{}, f.defaultInformer)
}

func (f *projectMembersInformer) Lister() v1.ProjectMembersLister {
	return v1.NewProjectMembersLister(f.Informer().GetIndexer())
}
