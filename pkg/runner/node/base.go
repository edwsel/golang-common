package node

// Node the single interface for services that must
// start and stop correctly without interruption (correct termination).
type Node interface {
	Name() string
	Run() error
	Close() error
}

// Instance Allows you to directly get an object without using reflection.
func Instance[T Node](node T) T {
	return node
}
