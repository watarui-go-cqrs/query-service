package categories

type Category struct {
	id   string
	name string
}

func NewCategory(id, name string) *Category {
	return &Category{
		id:   id,
		name: name,
	}
}

func (c *Category) Id() string {
	return c.id
}

func (c *Category) Name() string {
	return c.name
}
