// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/ca-gip/kubi-members/pkg/apis/ca-gip/v1"
	scheme "github.com/ca-gip/kubi-members/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ProjectMembersGetter has a method to return a ProjectMemberInterface.
// A group's client should implement this interface.
type ProjectMembersGetter interface {
	ProjectMembers(namespace string) ProjectMemberInterface
}

// ProjectMemberInterface has methods to work with ProjectMember resources.
type ProjectMemberInterface interface {
	Create(*v1.ProjectMember) (*v1.ProjectMember, error)
	Update(*v1.ProjectMember) (*v1.ProjectMember, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ProjectMember, error)
	List(opts metav1.ListOptions) (*v1.ProjectMemberList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ProjectMember, err error)
	ProjectMemberExpansion
}

// projectMembers implements ProjectMemberInterface
type projectMembers struct {
	client rest.Interface
	ns     string
}

// newProjectMembers returns a ProjectMembers
func newProjectMembers(c *CagipV1Client, namespace string) *projectMembers {
	return &projectMembers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the projectMember, and returns the corresponding projectMember object, and an error if there is any.
func (c *projectMembers) Get(name string, options metav1.GetOptions) (result *v1.ProjectMember, err error) {
	result = &v1.ProjectMember{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("projectmembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ProjectMembers that match those selectors.
func (c *projectMembers) List(opts metav1.ListOptions) (result *v1.ProjectMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ProjectMemberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("projectmembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested projectMembers.
func (c *projectMembers) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("projectmembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a projectMember and creates it.  Returns the server's representation of the projectMember, and an error, if there is any.
func (c *projectMembers) Create(projectMember *v1.ProjectMember) (result *v1.ProjectMember, err error) {
	result = &v1.ProjectMember{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("projectmembers").
		Body(projectMember).
		Do().
		Into(result)
	return
}

// Update takes the representation of a projectMember and updates it. Returns the server's representation of the projectMember, and an error, if there is any.
func (c *projectMembers) Update(projectMember *v1.ProjectMember) (result *v1.ProjectMember, err error) {
	result = &v1.ProjectMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("projectmembers").
		Name(projectMember.Name).
		Body(projectMember).
		Do().
		Into(result)
	return
}

// Delete takes name of the projectMember and deletes it. Returns an error if one occurs.
func (c *projectMembers) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("projectmembers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *projectMembers) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("projectmembers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched projectMember.
func (c *projectMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ProjectMember, err error) {
	result = &v1.ProjectMember{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("projectmembers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
