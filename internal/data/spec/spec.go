package spec

import (
	"fmt"
	"strings"
)

type ISpec interface {
	GetQuery() string
	GetValues() []interface{}
}

type ConcatSpec struct {
	Specs     []ISpec
	Separator string
}

func (s ConcatSpec) GetQuery() string {
	predicates := make([]string, len(s.Specs))
	for i, spec := range s.Specs {
		predicate := spec.GetQuery()
		if _, ok := spec.(ConcatSpec); ok {
			if predicate == "" {
				continue
			}
			if len(spec.GetValues()) > 1 {
				predicate = fmt.Sprintf("(%s)", predicate)
			}
		}
		predicates[i] = predicate
	}
	return strings.Join(predicates, fmt.Sprintf(" %s ", s.Separator))
}

func (s ConcatSpec) GetValues() []interface{} {
	values := make([]interface{}, 0)
	for _, spec := range s.Specs {
		values = append(values, spec.GetValues()...)
	}
	return values
}

func And(specs ...ISpec) ConcatSpec {
	return ConcatSpec{
		Specs:     specs,
		Separator: "AND",
	}
}

func Or(specs ...ISpec) ConcatSpec {
	return ConcatSpec{
		Specs:     specs,
		Separator: "OR",
	}
}

type OperatorSpec struct {
	Field    string
	Operator string
	Value    interface{}
}

func (s OperatorSpec) GetQuery() string {
	return fmt.Sprintf("%s %s ?", s.Field, s.Operator)
}

func (s OperatorSpec) GetValues() []interface{} {
	return []interface{}{s.Value}
}

func Equals(field string, value interface{}) OperatorSpec {
	return OperatorSpec{
		Field:    field,
		Operator: "=",
		Value:    value,
	}
}

func NotEquals(field string, value interface{}) OperatorSpec {
	return OperatorSpec{
		Field:    field,
		Operator: "!=",
		Value:    value,
	}
}

func GreaterThan(field string, value interface{}) OperatorSpec {
	return OperatorSpec{
		Field:    field,
		Operator: ">",
		Value:    value,
	}
}

func GreaterThanOrEqual(field string, value interface{}) OperatorSpec {
	return OperatorSpec{
		Field:    field,
		Operator: ">=",
		Value:    value,
	}
}

func LessThan(field string, value interface{}) OperatorSpec {
	return OperatorSpec{
		Field:    field,
		Operator: "<",
		Value:    value,
	}
}

func LessThanOrEqual(field string, value interface{}) OperatorSpec {
	return OperatorSpec{
		Field:    field,
		Operator: "<=",
		Value:    value,
	}
}

func In(field string, value interface{}) OperatorSpec {
	return OperatorSpec{
		Field:    field,
		Operator: "IN",
		Value:    value,
	}
}

func NotIn(field string, value interface{}) OperatorSpec {
	return OperatorSpec{
		Field:    field,
		Operator: "NOT IN",
		Value:    value,
	}
}
