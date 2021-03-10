package controller

import (
	"github.com/ca-gip/kubi-members/internal/ldap"
	v1 "github.com/ca-gip/kubi-members/pkg/apis/ca-gip/v1"
	membersclientset "github.com/ca-gip/kubi-members/pkg/generated/clientset/versioned"
	kubiv1 "github.com/ca-gip/kubi/pkg/apis/ca-gip/v1"
	projectclientset "github.com/ca-gip/kubi/pkg/generated/clientset/versioned"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

type Controller struct {
	configmapclientset kubernetes.Interface
	projectclientset   projectclientset.Interface
	membersclientset   membersclientset.Interface
	ldap               *ldap.Ldap
}

func NewController(configMapClient kubernetes.Interface, projectClient projectclientset.Interface, membersClient membersclientset.Interface, ldap *ldap.Ldap) *Controller {
	return &Controller{
		configmapclientset: configMapClient,
		projectclientset:   projectClient,
		membersclientset:   membersClient,
		ldap:               ldap,
	}
}

func (c *Controller) Preflight() {
}

func (c *Controller) Run() (err error) {
	projects, err := c.projectclientset.CagipV1().Projects().List(metav1.ListOptions{})
	if err != nil {
		klog.Errorf("Could list project : %s", err)
		return err
	}

	for _, project := range projects.Items {
		if project.Status.Name == kubiv1.ProjectStatusCreated {
			c.SyncMembers(&project)
		}
	}
	return
}

func (c *Controller) SyncMembers(project *kubiv1.Project) {
	members, err := c.ldap.Search(project.Spec.SourceDN)
	if err != nil {
		klog.Errorf("Could not ldap members for %s : %s", project.Spec.SourceDN, err)
	}
	projectMembers := c.templateProjectMembers(project, members)
	c.updateMembers(project.Name, projectMembers)

	savedMembers, err := c.membersclientset.CagipV1().ProjectMembers(project.Name).List(metav1.ListOptions{})
	if err != nil {
		klog.Errorf("Could not list project members for %s : %s", project.Name, err)
		return
	}

	for _, savedMember := range savedMembers.Items {
		if !members.Exist(savedMember.Dn) {
			err := c.membersclientset.CagipV1().ProjectMembers(project.Name).Delete(savedMember.Name, &metav1.DeleteOptions{})
			if err != nil {
				klog.Errorf("Could not delete project member %s : %s", savedMember.Name, err)
			}
		}
	}

}

func (c *Controller) updateMembers(namespace string, members []*v1.ProjectMember) {
	for _, member := range members {
		_, err := c.membersclientset.CagipV1().ProjectMembers(namespace).Get(member.Username, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			_, err := c.membersclientset.CagipV1().ProjectMembers(namespace).Create(member)
			if err != nil {
				klog.Errorf("Could not create ProjectMember %s : %s", member.Username, err)
			}
		}
	}
	return
}

func (c *Controller) templateProjectMember(project *kubiv1.Project, user ldap.User) *v1.ProjectMember {
	return &v1.ProjectMember{
		ObjectMeta: metav1.ObjectMeta{
			Name:      user.Username,
			Namespace: project.Name,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(project, kubiv1.SchemeGroupVersion.WithKind("Project")),
			},
		},
		Dn:       user.Dn,
		Username: user.Username,
		Mail:     user.Mail,
	}
}

func (c *Controller) templateProjectMembers(project *kubiv1.Project, users []ldap.User) (members []*v1.ProjectMember) {
	for _, user := range users {
		member := c.templateProjectMember(project, user)
		members = append(members, member)
	}
	return
}