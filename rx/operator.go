package rx

func Map[T, V any](source Observable[T], f func(T) V) Observable[V] {
	panic("???")
}

func FlatMap[T, V any](source Observable[T], f func(T) Observable[V]) Observable[V] {
	panic("???")
}
