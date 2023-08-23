package model

import "time"

type StringFilter struct {
	Eq *string `json:"eq"`
	Ne *string `json:"ne"`

	Re *string `json:"re"`

	In    []string `json:"in"`
	NotIn []string `json:"notIn"`

	And []StringFilter `json:"and"`
	Or  []StringFilter `json:"or"`
}

type IDFilter struct {
	Eq *GlobalID `json:"eq"`
	Ne *GlobalID `json:"ne"`

	In    []GlobalID `json:"in"`
	NotIn []GlobalID `json:"notIn"`

	And []IDFilter `json:"and"`
	Or  []IDFilter `json:"or"`
}

type TimeFilter struct {
	Eq *time.Time `json:"eq"`
	Ne *time.Time `json:"ne"`

	Gt *time.Time `json:"gt"`
	Ge *time.Time `json:"ge"`
	Lt *time.Time `json:"lt"`
	Le *time.Time `json:"le"`

	And []TimeFilter `json:"and"`
	Or  []TimeFilter `json:"or"`
}

type NodeFilter struct {
	ID *IDFilter `json:"id"`

	CreatedBy *StringFilter `json:"createdBy"`
	CreatedAt *TimeFilter   `json:"createdAt"`

	UpdatedBy *StringFilter `json:"updatedBy"`
	UpdatedAt *TimeFilter   `json:"updatedAt"`
}
