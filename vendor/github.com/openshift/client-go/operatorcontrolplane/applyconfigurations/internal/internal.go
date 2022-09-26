// Code generated by applyconfiguration-gen. DO NOT EDIT.

package internal

import (
	"fmt"
	"sync"

	typed "sigs.k8s.io/structured-merge-diff/v4/typed"
)

func Parser() *typed.Parser {
	parserOnce.Do(func() {
		var err error
		parser, err = typed.NewParser(schemaYAML)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse schema: %v", err))
		}
	})
	return parser
}

var parserOnce sync.Once
var parser *typed.Parser
var schemaYAML = typed.YAMLObject(`types:
- name: com.github.openshift.api.config.v1.SecretNameReference
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.operatorcontrolplane.v1alpha1.LogEntry
  map:
    fields:
    - name: latency
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Duration
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: success
      type:
        scalar: boolean
      default: false
    - name: time
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
- name: com.github.openshift.api.operatorcontrolplane.v1alpha1.OutageEntry
  map:
    fields:
    - name: end
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: endLogs
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.operatorcontrolplane.v1alpha1.LogEntry
          elementRelationship: atomic
    - name: message
      type:
        scalar: string
    - name: start
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: startLogs
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.operatorcontrolplane.v1alpha1.LogEntry
          elementRelationship: atomic
- name: com.github.openshift.api.operatorcontrolplane.v1alpha1.PodNetworkConnectivityCheck
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.operatorcontrolplane.v1alpha1.PodNetworkConnectivityCheckSpec
      default: {}
    - name: status
      type:
        namedType: com.github.openshift.api.operatorcontrolplane.v1alpha1.PodNetworkConnectivityCheckStatus
      default: {}
- name: com.github.openshift.api.operatorcontrolplane.v1alpha1.PodNetworkConnectivityCheckCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.operatorcontrolplane.v1alpha1.PodNetworkConnectivityCheckSpec
  map:
    fields:
    - name: sourcePod
      type:
        scalar: string
      default: ""
    - name: targetEndpoint
      type:
        scalar: string
      default: ""
    - name: tlsClientCert
      type:
        namedType: com.github.openshift.api.config.v1.SecretNameReference
      default: {}
- name: com.github.openshift.api.operatorcontrolplane.v1alpha1.PodNetworkConnectivityCheckStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.operatorcontrolplane.v1alpha1.PodNetworkConnectivityCheckCondition
          elementRelationship: associative
          keys:
          - type
    - name: failures
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.operatorcontrolplane.v1alpha1.LogEntry
          elementRelationship: atomic
    - name: outages
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.operatorcontrolplane.v1alpha1.OutageEntry
          elementRelationship: atomic
    - name: successes
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.operatorcontrolplane.v1alpha1.LogEntry
          elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Duration
  scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: fieldsType
      type:
        scalar: string
    - name: fieldsV1
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
    - name: manager
      type:
        scalar: string
    - name: operation
      type:
        scalar: string
    - name: subresource
      type:
        scalar: string
    - name: time
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
  map:
    fields:
    - name: annotations
      type:
        map:
          elementType:
            scalar: string
    - name: creationTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: deletionGracePeriodSeconds
      type:
        scalar: numeric
    - name: deletionTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: finalizers
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: generateName
      type:
        scalar: string
    - name: generation
      type:
        scalar: numeric
    - name: labels
      type:
        map:
          elementType:
            scalar: string
    - name: managedFields
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
          elementRelationship: atomic
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: ownerReferences
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
          elementRelationship: associative
          keys:
          - uid
    - name: resourceVersion
      type:
        scalar: string
    - name: selfLink
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
      default: ""
    - name: blockOwnerDeletion
      type:
        scalar: boolean
    - name: controller
      type:
        scalar: boolean
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: uid
      type:
        scalar: string
      default: ""
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Time
  scalar: untyped
- name: __untyped_atomic_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
- name: __untyped_deduced_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_deduced_
    elementRelationship: separable
`)
