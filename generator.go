package iterators

type Generator[T any] struct {
	Results chan T
}

func (g *Generator[T]) Next() (result T, ok bool) {
	result, ok = <-g.Results
	return result, ok
}

func (g *Generator[T]) Close() {
	close(g.Results)
}
