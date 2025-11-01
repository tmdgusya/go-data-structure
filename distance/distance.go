package distance

func dist(a float32, b float32) float32 {
	diff := a - b
	if diff < 0 {
		return -diff
	}
	return diff
}

func LinearScanClosetNeighbor(arr []float32, target float32) float32 {
	n := len(arr)
	if n == 0 {
		return -1
	}

	candidate := arr[0]
	closet_dist := dist(target, candidate)

	for i := 1; i < n; i++ {
		curr_dist := dist(target, arr[i])
		if curr_dist < closet_dist {
			closet_dist = curr_dist
			candidate = arr[i]
		}
	}

	return candidate
}
