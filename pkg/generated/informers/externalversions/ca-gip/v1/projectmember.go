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

// ProjectMemberInformer provides access to a shared informer and lister for
// ProjectMembers.
type ProjectMemberInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ProjectMemberLister
}

type projectMemberInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewProjectMemberInformer constructs a new informer for ProjectMember type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewProjectMemberInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredProjectMemberInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredProjectMemberInformer constructs a new informer for ProjectMember type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredProjectMemberInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CagipV1().ProjectMembers(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CagipV1().ProjectMembers(namespace).Watch(options)
			},
		},
		&cagipv1.ProjectMember{},
		resyncPeriod,
		indexers,
	)
}

func (f *projectMemberInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredProjectMemberInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *projectMemberInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cagipv1.ProjectMember{}, f.defaultInformer)
}

func (f *projectMemberInformer) Lister() v1.ProjectMemberLister {
	return v1.NewProjectMemberLister(f.Informer().GetIndexer())
}
