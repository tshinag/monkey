package object

// Hashable is the expression of hashable object
type Hashable interface {
	HashKey() HashKey
}
