package geometry

func (path Path) Cost() float64 {
	var cost float64 = 0.0
	for ix, point := range path {
		if ix == 0 {
			continue
		} else {
			cost = cost + point.Distance(path[ix - 1])
		}
	}
	return cost
}
