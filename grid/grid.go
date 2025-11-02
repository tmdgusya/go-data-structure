package grid

import "math"

const threshold = 1e-9

type GridPoint struct {
	x    float64
	y    float64
	next *GridPoint
}

type Grid struct {
	num_x_bins  int
	num_y_bins  int
	x_start     float64
	x_end       float64
	y_start     float64
	y_end       float64
	x_bin_width float64
	y_bin_width float64
	bins        [][]*GridPoint
}

func approx_equal(x1 float64, y1 float64, x2 float64, y2 float64) bool {
	if math.Abs(x1-x2) > threshold {
		return false
	}

	if math.Abs(y1-y2) > threshold {
		return false
	}

	return true
}

func (g *Grid) CalculateWidth() {
	g.x_bin_width = (g.x_end - g.x_start) / float64(g.num_x_bins)
	g.y_bin_width = (g.y_end - g.y_start) / float64(g.num_y_bins)
}

func (g *Grid) GetBins(x float64, y float64) (int, int) {
	xbin := (x - g.x_start) / g.x_bin_width
	ybin := (y - g.y_start) / g.y_bin_width
	return int(math.Floor(xbin)), int(math.Floor(ybin))
}

func (g *Grid) Insert(x float64, y float64) bool {
	xbin, ybin := g.GetBins(x, y)

	if xbin < 0 || xbin >= g.num_x_bins {
		return false
	}

	if ybin < 0 || ybin >= g.num_y_bins {
		return false
	}

	next_point := g.bins[xbin][ybin]
	g.bins[xbin][ybin] = &GridPoint{x, y, next_point}

	return true
}

func (g *Grid) Delete(x float64, y float64) bool {
	xbin, ybin := g.GetBins(x, y)

	if xbin < 0 || xbin >= g.num_x_bins {
		return false
	}

	if ybin < 0 || ybin >= g.num_y_bins {
		return false
	}

	if g.bins[xbin][ybin] == nil {
		return false
	}

	current := g.bins[xbin][ybin]
	var previous *GridPoint

	for current != nil {
		if approx_equal(current.x, current.y, x, y) {
			if previous == nil {
				g.bins[xbin][ybin] = current.next
			} else {
				previous.next = current.next
			}
			return true
		}

		previous = current
		current = current.next
	}

	return false
}

func (g *Grid) MinDistToBind(xbin int, ybin int, x float64, y float64) float64 {
	if xbin < 0 || xbin >= g.num_x_bins {
		// 잘못된 값임을 나타내는
		return math.Inf(-1)
	}

	if ybin < 0 || ybin >= g.num_y_bins {
		return math.Inf(-1)
	}

	x_min := g.x_start + float64(xbin)*g.x_bin_width
	x_max := g.x_start + float64(xbin+1)*g.x_bin_width
	x_dist := float64(0)
	if x < x_min {
		x_dist = x_min - x
	}
	if x > x_max {
		x_dist = x - x_max
	}

	y_min := g.y_start + float64(ybin)*g.y_bin_width
	y_max := g.y_start + float64(ybin+1)*g.y_bin_width
	y_dist := float64(0)
	if y < y_min {
		y_dist = y_min - y
	}
	if y > y_max {
		y_dist = y - y_max
	}

	return math.Sqrt(x_dist*x_dist + y_dist*y_dist)
}
