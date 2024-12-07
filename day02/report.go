package day02

type Report struct {
	levels []int
}

func (this Report) IsValid() bool {
	var isIncreasing = true
	var isDecreasing = true

	for i := 0; i < len(this.levels)-1; i++ {
		if(isIncreasing && this.levels[i] >= this.levels[i+1]) {
			isIncreasing = false
		}
		if(isDecreasing && this.levels[i] <= this.levels[i+1]) {
			isDecreasing = false
		}
		
		if(!isIncreasing && !isDecreasing) {
			return false;
		}
	}
	return true
}
