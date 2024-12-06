package day02

type Report struct {
	levels []int
}

func (this Report) IsValid() bool {
	for i := 0; i < len(this.levels)-1; i++ {
		if(this.levels[i] > this.levels[i+1])	{
			return false
		}
	}
	return true
}
