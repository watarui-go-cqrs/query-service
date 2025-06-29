package products

type ProductAdapter interface {
	Convert(source *Product) any
	Rebuild(source any) (dest *Product, err error)
}
