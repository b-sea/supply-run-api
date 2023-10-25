// Package model defines all data entities shared between front end, service, and repository layers.
package model

// StringFilter defines all properties to filter a string value.
type StringFilter struct {
	Eq *string `json:"eq"`
	Ne *string `json:"ne"`

	Re *string `json:"re"`

	In    []string `json:"in"`
	NotIn []string `json:"notIn"`

	And []StringFilter `json:"and"`
	Or  []StringFilter `json:"or"`
}
