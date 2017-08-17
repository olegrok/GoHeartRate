package pull

type Worker interface {
	Work()
}

func CreatePull(n int) (chan<-Worker) {
	ch := make(chan Worker, n)
	go func() {
		for worker := range ch {
			worker.Work()
		}
	}()

	return ch
}