package go_rx

// Subject : Observable と Observer を兼ね備えたインターフェース
type Subject[T any] interface {
	Observable[T]
	Observer[T]
}
