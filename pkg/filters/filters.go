/*
Copyright 2020 CaiCloud, Inc. All rights reserved.

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

package filters

import (
	"strings"

	"github.com/caicloud/event_exporter/pkg/options"
	v1 "k8s.io/api/core/v1"
)

type EventFilter interface {
	Filter(event *v1.Event) bool
}

type EventTypeFilter struct {
	AllowedTypes []string
	CustomFilter options.CustomFilter
}

func NewEventTypeFilter(allowedTypes []string, customFilter options.CustomFilter) *EventTypeFilter {
	return &EventTypeFilter{
		AllowedTypes: allowedTypes,
		CustomFilter: customFilter,
	}
}

func (e *EventTypeFilter) Filter(event *v1.Event) bool {
	var isMatching []bool
	if e.CustomFilter.InvolvedObjectKind != "" {
		if event.InvolvedObject.Kind != "" && event.InvolvedObject.Kind == e.CustomFilter.InvolvedObjectKind {
			isMatching = append(isMatching, true)
		} else {
			isMatching = append(isMatching, false)
		}
	}
	if e.CustomFilter.InvolvedObjectName != "" {
		if event.InvolvedObject.Name != "" && event.InvolvedObject.Name == e.CustomFilter.InvolvedObjectName {
			isMatching = append(isMatching, true)
		} else {
			isMatching = append(isMatching, false)
		}
	}
	if e.CustomFilter.InvolvedObjectNamespace != "" {
		if event.InvolvedObject.Namespace != "" && event.InvolvedObject.Namespace == e.CustomFilter.InvolvedObjectNamespace {
			isMatching = append(isMatching, true)
		} else {
			isMatching = append(isMatching, false)
		}
	}

	if any(isMatching) {
		for _, allowedType := range e.CustomFilter.EventTypes {
			if strings.EqualFold(event.Type, allowedType) {
				return true
			}
		}
	} else {
		for _, allowedType := range e.AllowedTypes {
			if strings.EqualFold(event.Type, allowedType) {
				return true
			}
		}
	}

	return false
}

func any(slice []bool) bool {
	for _, value := range slice {
		if value {
			return true
		}
	}
	return false
}
