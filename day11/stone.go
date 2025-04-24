package day11

type Stone struct {
	engravedNumber int64
}

func (this Stone) OnBlink() []Stone {
	return []Stone{Stone{engravedNumber: 1}}
}
