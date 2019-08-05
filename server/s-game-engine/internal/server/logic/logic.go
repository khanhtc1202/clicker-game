package logic

var pastFourScores = make([]float64, 4)

func GetSize() float64 {
	oldScores := pastFourScores[0] + pastFourScores[1]
	newScores := pastFourScores[2] + pastFourScores[3]

	diff := newScores - oldScores
	if diff > 0.0 {
		size := 600.0 + diff*60.0
		if size < 2000.0 {
			return size
		}
		return 2000.0
	}
	if diff > -0.5 && diff <= 0.0 {
		return 100.0 + diff*18.0
	}
	return 10.0
}

func SetScore(score float64) bool {
	pastFourScores = append(pastFourScores, score)
	pastFourScores = pastFourScores[1:]
	return true
}
