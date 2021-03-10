// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/ca-gip/kubi-members/pkg/apis/ca-gip/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProjectMembersLister helps list ProjectMemberses.
type ProjectMembersLister interface {
	// List lists all ProjectMemberses in the indexer.
	List(selector labels.Selector) (ret []*v1.ProjectMember, err error)
	// ProjectMemberses returns an object that can list and get ProjectMemberses.
	ProjectMemberses(namespace string) ProjectMembersNamespaceLister
	ProjectMembersListerExpansion
}

// projectMembersLister implements the ProjectMembersLister interface.
type projectMembersLister struct {
	indexer cache.Indexer
}

// NewProjectMembersLister returns a new ProjectMembersLister.
func NewProjectMembersLister(indexer cache.Indexer) ProjectMembersLister {
	return &projectMembersLister{indexer: indexer}
}

// List lists all ProjectMemberses in the indexer.
func (s *projectMembersLister) List(selector labels.Selector) (ret []*v1.ProjectMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ProjectMember))
	})
	return ret, err
}

// ProjectMemberses returns an object that can list and get ProjectMemberses.
func (s *projectMembersLister) ProjectMemberses(namespace string) ProjectMembersNamespaceLister {
	return projectMembersNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProjectMembersNamespaceLister helps list and get ProjectMemberses.
type ProjectMembersNamespaceLister interface {
	// List lists all ProjectMemberses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ProjectMember, err error)
	// Get retrieves the ProjectMember from the indexer for a given namespace and name.
	Get(name string) (*v1.ProjectMember, error)
	ProjectMembersNamespaceListerExpansion
}

// projectMembersNamespaceLister implements the ProjectMembersNamespaceLister
// interface.
type projectMembersNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProjectMemberses in the indexer for a given namespace.
func (s projectMembersNamespaceLister) List(selector labels.Selector) (ret []*v1.ProjectMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ProjectMember))
	})
	return ret, err
}

// Get retrieves the ProjectMember from the indexer for a given namespace and name.
func (s projectMembersNamespaceLister) Get(name string) (*v1.ProjectMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("projectmembers"), name)
	}
	return obj.(*v1.ProjectMember), nil
}
