package model

type StringFilter struct {
	Eq *string `json:"eq"`
	Ne *string `json:"ne"`

	Re *string `json:"re"`

	In    []string `json:"in"`
	NotIn []string `json:"notIn"`

	And []StringFilter `json:"and"`
	Or  []StringFilter `json:"or"`
}
