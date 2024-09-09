package iterators

func Map[T any](iterable []T, mapper func(elem T) T) Generator[T] {
	gen := Generator[T]{}.New(1)

	go func() {
		for _, elem := range iterable {
			ok := gen.Put(mapper(elem))
			if !ok {
				break
			}
		}

		gen.Close()
	}()

	return gen
}
