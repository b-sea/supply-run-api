// Package model defines all data entities shared between front end, service, and repository layers.
package model

type Filter interface {
	IsFilter()
}

// IDFilter defines all proerties to filter an ID value.
type IDFilter struct {
	Eq    *ID      `json:"eq"`
	Ne    *ID      `json:"ne"`
	In    []string `json:"in"`
	NotIn []string `json:"notIn"`

	And []StringFilter `json:"and"`
	Or  []StringFilter `json:"or"`
}

// StringFilter defines all properties to filter a string value.
type StringFilter struct {
	Eq    *string  `json:"eq"`
	Ne    *string  `json:"ne"`
	Re    *string  `json:"re"`
	In    []string `json:"in"`
	NotIn []string `json:"notIn"`

	And []StringFilter `json:"and"`
	Or  []StringFilter `json:"or"`
}
