package network

import (
	clocktesting "k8s.io/utils/clock/testing"
	"reflect"
	"testing"
	"time"

	"sigs.k8s.io/yaml"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"

	configv1 "github.com/openshift/api/config/v1"
	configlistersv1 "github.com/openshift/client-go/config/listers/config/v1"

	"github.com/openshift/library-go/pkg/operator/events"
)

func TestObserveClusterCIDRs(t *testing.T) {
	type Test struct {
		name          string
		config        *configv1.Network
		expected      []string
		expectedError bool
	}
	tests := []Test{
		{
			"clusterNetworks",
			&configv1.Network{
				ObjectMeta: metav1.ObjectMeta{Name: "cluster"},
				Status: configv1.NetworkStatus{
					ClusterNetwork: []configv1.ClusterNetworkEntry{
						{CIDR: "podCIDR1"},
						{CIDR: "podCIDR2"},
					},
				},
			},
			[]string{"podCIDR1", "podCIDR2"},
			false,
		},
		{
			"none, no old config",
			&configv1.Network{
				ObjectMeta: metav1.ObjectMeta{Name: "cluster"},
				Status:     configv1.NetworkStatus{},
			},
			nil,
			true,
		},
		{
			"none, existing config",
			&configv1.Network{
				ObjectMeta: metav1.ObjectMeta{Name: "cluster"},
				Status:     configv1.NetworkStatus{},
			},
			nil,
			true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{})
			if err := indexer.Add(test.config); err != nil {
				t.Fatal(err.Error())
			}
			result, err := GetClusterCIDRs(configlistersv1.NewNetworkLister(indexer), events.NewInMemoryRecorder("network", clocktesting.NewFakePassiveClock(time.Now())))
			if err != nil && !test.expectedError {
				t.Fatal(err)
			} else if err == nil {
				if test.expectedError {
					t.Fatalf("expected error, but got none")
				}
				if !reflect.DeepEqual(test.expected, result) {
					t.Errorf("\n===== observed config expected:\n%v\n===== observed config actual:\n%v", toYAML(test.expected), toYAML(result))
				}
			}
		})
	}
}

func TestObserveServiceClusterIPRanges(t *testing.T) {
	indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{})
	if err := indexer.Add(&configv1.Network{
		ObjectMeta: metav1.ObjectMeta{Name: "cluster"},
		Status: configv1.NetworkStatus{
			ServiceNetwork: []string{"serviceCIDRv4", "serviceCIDRv6"},
		},
	},
	); err != nil {
		t.Fatal(err.Error())
	}
	result, err := GetServiceCIDRs(configlistersv1.NewNetworkLister(indexer), events.NewInMemoryRecorder("network", clocktesting.NewFakePassiveClock(time.Now())))
	if err != nil {
		t.Fatal(err)
	}

	if expected := []string{"serviceCIDRv4", "serviceCIDRv6"}; !reflect.DeepEqual(expected, result) {
		t.Errorf("\n===== observed config expected:\n%v\n===== observed config actual:\n%v", toYAML(expected), toYAML(result))
	}
}

func TestObserveServiceNodePortRange(t *testing.T) {
	type Test struct {
		name          string
		config        *configv1.Network
		expected      string
		expectedError bool
	}
	tests := []Test{
		{
			"invalid range",
			&configv1.Network{
				ObjectMeta: metav1.ObjectMeta{Name: "cluster"},
				Spec: configv1.NetworkSpec{
					ServiceNodePortRange: "65530-67840",
				},
			},
			"",
			true,
		},
		{
			"low higher than high",
			&configv1.Network{
				ObjectMeta: metav1.ObjectMeta{Name: "cluster"},
				Spec: configv1.NetworkSpec{
					ServiceNodePortRange: "40000-38930",
				},
			},
			"",
			true,
		},
		{
			"valid range",
			&configv1.Network{
				ObjectMeta: metav1.ObjectMeta{Name: "cluster"},
				Spec: configv1.NetworkSpec{
					ServiceNodePortRange: "40000-46000",
				},
			},
			"40000-46000",
			false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{})
			if err := indexer.Add(test.config); err != nil {
				t.Fatal(err.Error())
			}
			result, err := GetServiceNodePortRange(configlistersv1.NewNetworkLister(indexer), events.NewInMemoryRecorder("network", clocktesting.NewFakePassiveClock(time.Now())))
			if err != nil && !test.expectedError {
				t.Fatal(err)
			} else if err == nil {
				if test.expectedError {
					t.Fatalf("expected error, but got none")
				}
				if !reflect.DeepEqual(test.expected, result) {
					t.Errorf("\n===== observed config expected:\n%v\n===== observed config actual:\n%v", toYAML(test.expected), toYAML(result))
				}
			}
		})
	}
}

func toYAML(o interface{}) string {
	b, e := yaml.Marshal(o)
	if e != nil {
		return e.Error()
	}
	return string(b)
}
