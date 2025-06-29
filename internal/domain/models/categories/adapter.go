package categories

type CategoryAdapter interface {
	Convert(source *Category) any
	Rebuild(source any) (dest *Category, err error)
}
