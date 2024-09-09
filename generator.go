package iterators

type Generator[T any] struct {
	results  chan T
	finish   chan struct{}
	IsClosed bool
}

func (g Generator[T]) New(size int) Generator[T] {
	return Generator[T]{
		results: make(chan T, size),
		finish:  make(chan struct{}),
	}
}

func (g Generator[T]) Next() (elem T, ok bool) {
	select {
	case <-g.finish:
	case elem, ok = <-g.results:
	}

	return elem, ok
}

func (g Generator[T]) Put(elem T) (ok bool) {
	select {
	case <-g.finish:
		return false
	case g.results <- elem:
		return true
	}
}

func (g *Generator[T]) Close() {
	if !g.IsClosed {
		g.IsClosed = true
		close(g.finish)
		close(g.results)
	}
}
