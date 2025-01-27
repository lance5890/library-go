// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// KubeControllerManagerStatusApplyConfiguration represents a declarative configuration of the KubeControllerManagerStatus type for use
// with apply.
type KubeControllerManagerStatusApplyConfiguration struct {
	StaticPodOperatorStatusApplyConfiguration `json:",inline"`
}

// KubeControllerManagerStatusApplyConfiguration constructs a declarative configuration of the KubeControllerManagerStatus type for use with
// apply.
func KubeControllerManagerStatus() *KubeControllerManagerStatusApplyConfiguration {
	return &KubeControllerManagerStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *KubeControllerManagerStatusApplyConfiguration) WithObservedGeneration(value int64) *KubeControllerManagerStatusApplyConfiguration {
	b.OperatorStatusApplyConfiguration.ObservedGeneration = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *KubeControllerManagerStatusApplyConfiguration) WithConditions(values ...*OperatorConditionApplyConfiguration) *KubeControllerManagerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.OperatorStatusApplyConfiguration.Conditions = append(b.OperatorStatusApplyConfiguration.Conditions, *values[i])
	}
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *KubeControllerManagerStatusApplyConfiguration) WithVersion(value string) *KubeControllerManagerStatusApplyConfiguration {
	b.OperatorStatusApplyConfiguration.Version = &value
	return b
}

// WithReadyReplicas sets the ReadyReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadyReplicas field is set to the value of the last call.
func (b *KubeControllerManagerStatusApplyConfiguration) WithReadyReplicas(value int32) *KubeControllerManagerStatusApplyConfiguration {
	b.OperatorStatusApplyConfiguration.ReadyReplicas = &value
	return b
}

// WithLatestAvailableRevision sets the LatestAvailableRevision field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LatestAvailableRevision field is set to the value of the last call.
func (b *KubeControllerManagerStatusApplyConfiguration) WithLatestAvailableRevision(value int32) *KubeControllerManagerStatusApplyConfiguration {
	b.OperatorStatusApplyConfiguration.LatestAvailableRevision = &value
	return b
}

// WithGenerations adds the given value to the Generations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Generations field.
func (b *KubeControllerManagerStatusApplyConfiguration) WithGenerations(values ...*GenerationStatusApplyConfiguration) *KubeControllerManagerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithGenerations")
		}
		b.OperatorStatusApplyConfiguration.Generations = append(b.OperatorStatusApplyConfiguration.Generations, *values[i])
	}
	return b
}

// WithLatestAvailableRevisionReason sets the LatestAvailableRevisionReason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LatestAvailableRevisionReason field is set to the value of the last call.
func (b *KubeControllerManagerStatusApplyConfiguration) WithLatestAvailableRevisionReason(value string) *KubeControllerManagerStatusApplyConfiguration {
	b.StaticPodOperatorStatusApplyConfiguration.LatestAvailableRevisionReason = &value
	return b
}

// WithNodeStatuses adds the given value to the NodeStatuses field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NodeStatuses field.
func (b *KubeControllerManagerStatusApplyConfiguration) WithNodeStatuses(values ...*NodeStatusApplyConfiguration) *KubeControllerManagerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithNodeStatuses")
		}
		b.StaticPodOperatorStatusApplyConfiguration.NodeStatuses = append(b.StaticPodOperatorStatusApplyConfiguration.NodeStatuses, *values[i])
	}
	return b
}
