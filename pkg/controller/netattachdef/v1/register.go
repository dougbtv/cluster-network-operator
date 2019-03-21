package v1

import (
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/apimachinery/pkg/runtime"
  "k8s.io/apimachinery/pkg/runtime/schema"
  // "github.com/openshift/cluster-network-operator/pkg/controller/netattachdef"
)

// SchemeGroupVersion is here
// GroupVersion is the identifier for the API which includes
// the name of the group and the version of the API
var SchemeGroupVersion = schema.GroupVersion{
  Group:   "k8s.cni.cncf.io",
  Version: "v1",
}

// create a SchemeBuilder which uses functions to add types to
// the scheme
// more comments here
var (
  SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
  AddToScheme   = SchemeBuilder.AddToScheme
)

// Resource specification
func Resource(resource string) schema.GroupResource {
  return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// addKnownTypes adds our types to the API scheme by registering
// NetworkAttachmentDefinition and NetworkAttachmentDefinitionSpec
func addKnownTypes(scheme *runtime.Scheme) error {
  scheme.AddKnownTypes(
    SchemeGroupVersion,
    &NetworkAttachmentDefinition{},
    &NetworkAttachmentDefinitionList{},
  )

  // register the type in the scheme
  metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
  return nil
}
