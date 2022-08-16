// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	operatorv1alpha1 "github.com/openshift/api/operator/v1alpha1"
	internal "github.com/openshift/client-go/operator/applyconfigurations/internal"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ImageContentSourcePolicyApplyConfiguration represents an declarative configuration of the ImageContentSourcePolicy type for use
// with apply.
type ImageContentSourcePolicyApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *ImageContentSourcePolicySpecApplyConfiguration `json:"spec,omitempty"`
}

// ImageContentSourcePolicy constructs an declarative configuration of the ImageContentSourcePolicy type for use with
// apply.
func ImageContentSourcePolicy(name string) *ImageContentSourcePolicyApplyConfiguration {
	b := &ImageContentSourcePolicyApplyConfiguration{}
	b.WithName(name)
	b.WithKind("ImageContentSourcePolicy")
	b.WithAPIVersion("operator.openshift.io/v1alpha1")
	return b
}

// ExtractImageContentSourcePolicy extracts the applied configuration owned by fieldManager from
// imageContentSourcePolicy. If no managedFields are found in imageContentSourcePolicy for fieldManager, a
// ImageContentSourcePolicyApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// imageContentSourcePolicy must be a unmodified ImageContentSourcePolicy API object that was retrieved from the Kubernetes API.
// ExtractImageContentSourcePolicy provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractImageContentSourcePolicy(imageContentSourcePolicy *operatorv1alpha1.ImageContentSourcePolicy, fieldManager string) (*ImageContentSourcePolicyApplyConfiguration, error) {
	return extractImageContentSourcePolicy(imageContentSourcePolicy, fieldManager, "")
}

// ExtractImageContentSourcePolicyStatus is the same as ExtractImageContentSourcePolicy except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractImageContentSourcePolicyStatus(imageContentSourcePolicy *operatorv1alpha1.ImageContentSourcePolicy, fieldManager string) (*ImageContentSourcePolicyApplyConfiguration, error) {
	return extractImageContentSourcePolicy(imageContentSourcePolicy, fieldManager, "status")
}

func extractImageContentSourcePolicy(imageContentSourcePolicy *operatorv1alpha1.ImageContentSourcePolicy, fieldManager string, subresource string) (*ImageContentSourcePolicyApplyConfiguration, error) {
	b := &ImageContentSourcePolicyApplyConfiguration{}
	err := managedfields.ExtractInto(imageContentSourcePolicy, internal.Parser().Type("com.github.openshift.api.operator.v1alpha1.ImageContentSourcePolicy"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(imageContentSourcePolicy.Name)

	b.WithKind("ImageContentSourcePolicy")
	b.WithAPIVersion("operator.openshift.io/v1alpha1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ImageContentSourcePolicyApplyConfiguration) WithKind(value string) *ImageContentSourcePolicyApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ImageContentSourcePolicyApplyConfiguration) WithAPIVersion(value string) *ImageContentSourcePolicyApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ImageContentSourcePolicyApplyConfiguration) WithName(value string) *ImageContentSourcePolicyApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ImageContentSourcePolicyApplyConfiguration) WithGenerateName(value string) *ImageContentSourcePolicyApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ImageContentSourcePolicyApplyConfiguration) WithNamespace(value string) *ImageContentSourcePolicyApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ImageContentSourcePolicyApplyConfiguration) WithUID(value types.UID) *ImageContentSourcePolicyApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ImageContentSourcePolicyApplyConfiguration) WithResourceVersion(value string) *ImageContentSourcePolicyApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ImageContentSourcePolicyApplyConfiguration) WithGeneration(value int64) *ImageContentSourcePolicyApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ImageContentSourcePolicyApplyConfiguration) WithCreationTimestamp(value metav1.Time) *ImageContentSourcePolicyApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ImageContentSourcePolicyApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *ImageContentSourcePolicyApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ImageContentSourcePolicyApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ImageContentSourcePolicyApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ImageContentSourcePolicyApplyConfiguration) WithLabels(entries map[string]string) *ImageContentSourcePolicyApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *ImageContentSourcePolicyApplyConfiguration) WithAnnotations(entries map[string]string) *ImageContentSourcePolicyApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *ImageContentSourcePolicyApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *ImageContentSourcePolicyApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *ImageContentSourcePolicyApplyConfiguration) WithFinalizers(values ...string) *ImageContentSourcePolicyApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *ImageContentSourcePolicyApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *ImageContentSourcePolicyApplyConfiguration) WithSpec(value *ImageContentSourcePolicySpecApplyConfiguration) *ImageContentSourcePolicyApplyConfiguration {
	b.Spec = value
	return b
}
