package models

type Product struct {
	Name   string
	Des    string
	Price  string
	Effect string
}

func GetProduct() []Product {

	ps := make([]Product, 0)

	ps = append(ps, Product{
		Name:   "产品1",
		Des:    `Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。`,
		Price:  "210",
		Effect: "祛痘",
	})

	ps = append(ps, Product{
		Name:   "产品1",
		Des:    `Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。`,
		Price:  "210",
		Effect: "祛痘",
	})

	ps = append(ps, Product{
		Name:   "产品1",
		Des:    `Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。`,
		Price:  "210",
		Effect: "祛痘",
	})

	ps = append(ps, Product{
		Name:   "产品1",
		Des:    `Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。`,
		Price:  "210",
		Effect: "祛痘",
	})

	ps = append(ps, Product{
		Name:   "产品1",
		Des:    `Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。`,
		Price:  "210",
		Effect: "祛痘",
	})

	ps = append(ps, Product{
		Name:   "产品1",
		Des:    `Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。`,
		Price:  "210",
		Effect: "祛痘",
	})

	return ps
}
