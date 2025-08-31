package shopping

import (
	"errors"

	"github.com/b-sea/supply-run-api/internal/entity"
)

var DefaultCategory = Category{
	name: "Random Stuff",
}

type Category struct {
	name string
}

func NewCategory(name string) (Category, error) {
	if name == "" {
		return Category{}, errors.New("category name cannot be blank")
	}

	return Category{name: name}, nil
}

func (c *Category) ID() entity.ID {
	return entity.NewSeededID([]byte(c.name))
}

func (c *Category) String() string {
	return c.name
}

type ItemOption func(item *Item) error

func SetItemName(name string) ItemOption {
	return func(item *Item) error {
		if name == "" {
			return errors.New("item name cannot be blank")
		}

		item.name = name

		return nil
	}
}

func SetItemCategory(category Category) ItemOption {
	return func(item *Item) error {
		item.category = category
		return nil
	}
}

func SetItemTags(tags ...entity.Tag) ItemOption {
	return func(item *Item) error {
		item.tags = make(map[entity.Tag]bool)

		for i := range tags {
			if !tags[i].IsValid() {
				continue
			}

			item.tags[tags[i]] = true
		}

		return nil
	}
}

type Item struct {
	id       entity.ID
	name     string
	category Category
	tags     map[entity.Tag]bool
}

func NewItem(id entity.ID, name string, options ...ItemOption) (*Item, error) {
	item := &Item{
		id:       id,
		name:     name,
		category: DefaultCategory,
		tags:     make(map[entity.Tag]bool),
	}

	errs := make([]error, 0)

	if !id.IsValid() {
		errs = append(errs, errors.New("item id is invalid"))
	}

	options = append([]ItemOption{SetItemName(name)}, options...)

	for _, option := range options {
		if err := option(item); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return item, nil
}

func (item *Item) Update(options ...ItemOption) error {
	errs := make([]error, 0)

	for _, option := range options {
		if err := option(item); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (item *Item) ID() entity.ID {
	return item.id
}

func (item *Item) Name() string {
	return item.name
}

func (item *Item) Category() Category {
	return item.category
}

func (item *Item) Tags() []entity.Tag {
	tags := make([]entity.Tag, 0)

	for tag := range item.tags {
		tags = append(tags, tag)
	}

	return tags
}
