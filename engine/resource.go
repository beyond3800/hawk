package hawk


type Resource interface {
	ToMap() H
}

type CollectionResource interface {
	ToSlice() []H
}