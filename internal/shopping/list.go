package shopping

import (
	"errors"

	"github.com/b-sea/supply-run-api/internal/entity"
)

type Want struct {
	itemID     entity.ID
	quantity   int
	crossedOff bool
}

func (w *Want) ItemID() entity.ID {
	return w.itemID
}

func (w *Want) Quantity() int {
	return w.quantity
}

func (w *Want) CrossedOff() bool {
	return w.crossedOff
}

type ListOption func(list *List) error

func SetListName(name string) ListOption {
	return func(list *List) error {
		if name == "" {
			return errors.New("list name cannot be blank")
		}

		list.name = name

		return nil
	}
}

func SetListDescription(desc string) ListOption {
	return func(list *List) error {
		list.desc = desc
		return nil
	}
}

func UpsertListWant(itemID entity.ID, quantity int, crossedOff bool) ListOption {
	return func(list *List) error {
		if !itemID.IsValid() {
			return errors.New("list item id is invalid")
		}

		if quantity < 0 {
			if err := RemoveListWant(itemID)(list); err != nil {
				return err
			}

			return nil
		}

		list.wants[itemID] = Want{
			itemID:     itemID,
			quantity:   quantity,
			crossedOff: crossedOff,
		}

		return nil
	}
}

func RemoveListWant(itemID entity.ID) ListOption {
	return func(list *List) error {
		delete(list.wants, itemID)
		return nil
	}
}

func SetListTags(tags ...entity.Tag) ListOption {
	return func(list *List) error {
		list.tags = make(map[entity.Tag]bool)

		for i := range tags {
			if !tags[i].IsValid() {
				continue
			}

			list.tags[tags[i]] = true
		}

		return nil
	}
}

type List struct {
	id    entity.ID
	name  string
	desc  string
	wants map[entity.ID]Want
	tags  map[entity.Tag]bool
}

func NewList(id entity.ID, name string, options ...ListOption) (*List, error) {
	list := &List{
		id:    id,
		name:  name,
		desc:  "",
		wants: make(map[entity.ID]Want),
		tags:  make(map[entity.Tag]bool),
	}

	errs := make([]error, 0)

	if !id.IsValid() {
		errs = append(errs, errors.New("list id is invalid"))
	}

	options = append([]ListOption{SetListName(name)}, options...)

	for _, option := range options {
		if err := option(list); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return list, nil
}

func (list *List) Update(options ...ListOption) error {
	errs := make([]error, 0)

	for _, option := range options {
		if err := option(list); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (list *List) ID() entity.ID {
	return list.id
}

func (list *List) Name() string {
	return list.name
}

func (list *List) Description() string {
	return list.desc
}

func (list *List) Wants() []Want {
	wants := make([]Want, 0)

	for _, want := range list.wants {
		wants = append(wants, want)
	}

	return wants
}

func (list *List) Tags() []entity.Tag {
	tags := make([]entity.Tag, 0)

	for tag := range list.tags {
		tags = append(tags, tag)
	}

	return tags
}
