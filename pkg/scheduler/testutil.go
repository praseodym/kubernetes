/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package scheduler

import (
	"fmt"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	clientset "k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/algorithm"
	"k8s.io/kubernetes/pkg/scheduler/algorithm/predicates"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	"k8s.io/kubernetes/pkg/scheduler/factory"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	internalqueue "k8s.io/kubernetes/pkg/scheduler/internal/queue"
)

// FakeConfigurator is an implementation for test.
type FakeConfigurator struct {
	Config *factory.Config
}

// GetPredicateMetadataProducer is not implemented yet.
func (fc *FakeConfigurator) GetPredicateMetadataProducer() (predicates.PredicateMetadataProducer, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetPredicates is not implemented yet.
func (fc *FakeConfigurator) GetPredicates(predicateKeys sets.String) (map[string]predicates.FitPredicate, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetHardPodAffinitySymmetricWeight is not implemented yet.
func (fc *FakeConfigurator) GetHardPodAffinitySymmetricWeight() int32 {
	panic("not implemented")
}

// MakeDefaultErrorFunc is not implemented yet.
func (fc *FakeConfigurator) MakeDefaultErrorFunc(backoff *internalqueue.PodBackoffMap, podQueue internalqueue.SchedulingQueue) func(pod *v1.Pod, err error) {
	return nil
}

// GetClient is not implemented yet.
func (fc *FakeConfigurator) GetClient() clientset.Interface {
	return nil
}

// GetScheduledPodLister is not implemented yet.
func (fc *FakeConfigurator) GetScheduledPodLister() corelisters.PodLister {
	return nil
}

// Create returns FakeConfigurator.Config
func (fc *FakeConfigurator) Create() (*factory.Config, error) {
	return fc.Config, nil
}

// CreateFromProvider returns FakeConfigurator.Config
func (fc *FakeConfigurator) CreateFromProvider(providerName string) (*factory.Config, error) {
	return fc.Config, nil
}

// CreateFromConfig returns FakeConfigurator.Config
func (fc *FakeConfigurator) CreateFromConfig(policy schedulerapi.Policy) (*factory.Config, error) {
	return fc.Config, nil
}

// CreateFromKeys returns FakeConfigurator.Config
func (fc *FakeConfigurator) CreateFromKeys(predicateKeys, priorityKeys sets.String, extenders []algorithm.SchedulerExtender) (*factory.Config, error) {
	return fc.Config, nil
}

// EmptyPluginRegistry is an empty plugin registry used in tests.
var EmptyPluginRegistry = framework.Registry{}
