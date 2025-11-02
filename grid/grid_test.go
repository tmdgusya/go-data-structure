package grid

import (
	"math"
	"testing"
)

func TestCalculateWidth(t *testing.T) {
	g := Grid{
		num_x_bins: 10,
		num_y_bins: 20,
		x_start:    0.0,
		x_end:      100.0,
		y_start:    0.0,
		y_end:      200.0,
	}

	g.CalculateWidth()

	expectedXWidth := 10.0
	expectedYWidth := 10.0

	if g.x_bin_width != expectedXWidth {
		t.Errorf("Expected x_bin_width to be %f, got %f", expectedXWidth, g.x_bin_width)
	}

	if g.y_bin_width != expectedYWidth {
		t.Errorf("Expected y_bin_width to be %f, got %f", expectedYWidth, g.y_bin_width)
	}
}

func TestGetBins(t *testing.T) {
	g := Grid{
		num_x_bins:  10,
		num_y_bins:  10,
		x_start:     0.0,
		x_end:       100.0,
		y_start:     0.0,
		y_end:       100.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
	}

	tests := []struct {
		name         string
		x            float64
		y            float64
		expectedXBin int
		expectedYBin int
	}{
		{"Origin", 0.0, 0.0, 0, 0},
		{"Middle of first bin", 5.0, 5.0, 0, 0},
		{"Second bin", 15.0, 15.0, 1, 1},
		{"Last bin", 95.0, 95.0, 9, 9},
		{"Edge of bin", 10.0, 10.0, 1, 1},
		{"Different x and y bins", 25.0, 75.0, 2, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xbin, ybin := g.GetBins(tt.x, tt.y)
			if xbin != tt.expectedXBin {
				t.Errorf("Expected xbin to be %d, got %d", tt.expectedXBin, xbin)
			}
			if ybin != tt.expectedYBin {
				t.Errorf("Expected ybin to be %d, got %d", tt.expectedYBin, ybin)
			}
		})
	}
}

func TestGetBinsWithNegativeStart(t *testing.T) {
	g := Grid{
		num_x_bins:  10,
		num_y_bins:  10,
		x_start:     -50.0,
		x_end:       50.0,
		y_start:     -50.0,
		y_end:       50.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
	}

	tests := []struct {
		name         string
		x            float64
		y            float64
		expectedXBin int
		expectedYBin int
	}{
		{"Negative coordinates", -40.0, -40.0, 1, 1},
		{"Zero", 0.0, 0.0, 5, 5},
		{"Positive coordinates", 40.0, 40.0, 9, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xbin, ybin := g.GetBins(tt.x, tt.y)
			if xbin != tt.expectedXBin {
				t.Errorf("Expected xbin to be %d, got %d", tt.expectedXBin, xbin)
			}
			if ybin != tt.expectedYBin {
				t.Errorf("Expected ybin to be %d, got %d", tt.expectedYBin, ybin)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	// Test valid insertion
	result := g.Insert(5.0, 5.0)
	if !result {
		t.Error("Expected Insert to return true for valid coordinates")
	}

	// Check that point was inserted
	xbin, ybin := g.GetBins(5.0, 5.0)
	if g.bins[xbin][ybin] == nil {
		t.Error("Expected point to be inserted in bin")
	}

	if g.bins[xbin][ybin].x != 5.0 || g.bins[xbin][ybin].y != 5.0 {
		t.Errorf("Expected point coordinates (5.0, 5.0), got (%f, %f)",
			g.bins[xbin][ybin].x, g.bins[xbin][ybin].y)
	}
}

func TestInsertMultiplePoints(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	// Insert multiple points in the same bin
	g.Insert(5.0, 5.0)
	g.Insert(6.0, 6.0)
	g.Insert(7.0, 7.0)

	xbin, ybin := g.GetBins(5.0, 5.0)

	// Check that all points are in the linked list
	count := 0
	current := g.bins[xbin][ybin]
	for current != nil {
		count++
		current = current.next
	}

	if count != 3 {
		t.Errorf("Expected 3 points in bin, got %d", count)
	}

	// Check that the most recent point is at the head
	if g.bins[xbin][ybin].x != 7.0 || g.bins[xbin][ybin].y != 7.0 {
		t.Errorf("Expected most recent point (7.0, 7.0) at head, got (%f, %f)",
			g.bins[xbin][ybin].x, g.bins[xbin][ybin].y)
	}
}

func TestInsertOutOfBounds(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	tests := []struct {
		name string
		x    float64
		y    float64
	}{
		{"Negative x", -5.0, 5.0},
		{"Negative y", 5.0, -5.0},
		{"Both negative", -5.0, -5.0},
		{"X too large", 35.0, 5.0},
		{"Y too large", 5.0, 35.0},
		{"Both too large", 35.0, 35.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.Insert(tt.x, tt.y)
			if result {
				t.Errorf("Expected Insert to return false for out of bounds coordinates (%f, %f)", tt.x, tt.y)
			}
		})
	}
}

func TestDeleteSinglePoint(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	// Insert a point
	g.Insert(5.0, 5.0)

	// Delete the point
	result := g.Delete(5.0, 5.0)
	if !result {
		t.Error("Expected Delete to return true for existing point")
	}

	// Verify bin is empty
	xbin, ybin := g.GetBins(5.0, 5.0)
	if g.bins[xbin][ybin] != nil {
		t.Error("Expected bin to be empty after deletion")
	}
}

func TestDeleteHeadFromMultiplePoints(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	// Insert multiple points
	g.Insert(5.0, 5.0)
	g.Insert(6.0, 6.0)
	g.Insert(7.0, 7.0)

	xbin, ybin := g.GetBins(5.0, 5.0)

	// Delete the head (most recent: 7.0, 7.0)
	result := g.Delete(7.0, 7.0)
	if !result {
		t.Error("Expected Delete to return true for existing point")
	}

	// Verify new head is the second point
	if g.bins[xbin][ybin].x != 6.0 || g.bins[xbin][ybin].y != 6.0 {
		t.Errorf("Expected new head to be (6.0, 6.0), got (%f, %f)",
			g.bins[xbin][ybin].x, g.bins[xbin][ybin].y)
	}

	// Verify count is now 2
	count := 0
	current := g.bins[xbin][ybin]
	for current != nil {
		count++
		current = current.next
	}
	if count != 2 {
		t.Errorf("Expected 2 points remaining, got %d", count)
	}
}

func TestDeleteMiddlePoint(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	// Insert three points
	g.Insert(5.0, 5.0)
	g.Insert(6.0, 6.0)
	g.Insert(7.0, 7.0)

	xbin, ybin := g.GetBins(5.0, 5.0)

	// Delete the middle point (6.0, 6.0)
	result := g.Delete(6.0, 6.0)
	if !result {
		t.Error("Expected Delete to return true for existing point")
	}

	// Verify count is now 2
	count := 0
	current := g.bins[xbin][ybin]
	for current != nil {
		count++
		current = current.next
	}
	if count != 2 {
		t.Errorf("Expected 2 points remaining, got %d", count)
	}

	// Verify the middle point is gone
	current = g.bins[xbin][ybin]
	found := false
	for current != nil {
		if current.x == 6.0 && current.y == 6.0 {
			found = true
		}
		current = current.next
	}
	if found {
		t.Error("Expected middle point (6.0, 6.0) to be deleted")
	}
}

func TestDeleteTailPoint(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	// Insert three points
	g.Insert(5.0, 5.0)
	g.Insert(6.0, 6.0)
	g.Insert(7.0, 7.0)

	xbin, ybin := g.GetBins(5.0, 5.0)

	// Delete the tail point (first inserted: 5.0, 5.0)
	result := g.Delete(5.0, 5.0)
	if !result {
		t.Error("Expected Delete to return true for existing point")
	}

	// Verify count is now 2
	count := 0
	current := g.bins[xbin][ybin]
	for current != nil {
		count++
		current = current.next
	}
	if count != 2 {
		t.Errorf("Expected 2 points remaining, got %d", count)
	}
}

func TestDeleteNonExistentPoint(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	// Insert a point
	g.Insert(5.0, 5.0)

	// Try to delete a different point in the same bin
	result := g.Delete(6.0, 6.0)
	if result {
		t.Error("Expected Delete to return false for non-existent point")
	}

	// Verify original point is still there
	xbin, ybin := g.GetBins(5.0, 5.0)
	if g.bins[xbin][ybin] == nil {
		t.Error("Expected original point to still exist")
	}
}

func TestDeleteFromEmptyBin(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	// Try to delete from empty bin
	result := g.Delete(5.0, 5.0)
	if result {
		t.Error("Expected Delete to return false for empty bin")
	}
}

func TestDeleteOutOfBounds(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	tests := []struct {
		name string
		x    float64
		y    float64
	}{
		{"Negative x", -5.0, 5.0},
		{"Negative y", 5.0, -5.0},
		{"X too large", 35.0, 5.0},
		{"Y too large", 5.0, 35.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.Delete(tt.x, tt.y)
			if result {
				t.Errorf("Expected Delete to return false for out of bounds coordinates (%f, %f)", tt.x, tt.y)
			}
		})
	}
}

func TestDeleteWithApproxEqual(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	// Insert a point
	g.Insert(5.0, 5.0)

	// Delete with slightly different coordinates (within threshold)
	result := g.Delete(5.0+1e-10, 5.0+1e-10)
	if !result {
		t.Error("Expected Delete to return true for approximately equal coordinates")
	}

	// Verify bin is empty
	xbin, ybin := g.GetBins(5.0, 5.0)
	if g.bins[xbin][ybin] != nil {
		t.Error("Expected bin to be empty after deletion")
	}
}

func TestMinDistToBinPointInsideBin(t *testing.T) {
	g := Grid{
		num_x_bins:  10,
		num_y_bins:  10,
		x_start:     0.0,
		x_end:       100.0,
		y_start:     0.0,
		y_end:       100.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
	}

	// Point inside bin (0, 0): [0, 10) x [0, 10)
	dist := g.MinDistToBind(0, 0, 5.0, 5.0)
	if dist != 0.0 {
		t.Errorf("Expected distance 0.0 for point inside bin, got %f", dist)
	}

	// Point at the edge (x_min, y_min)
	dist = g.MinDistToBind(1, 1, 10.0, 10.0)
	if dist != 0.0 {
		t.Errorf("Expected distance 0.0 for point at edge of bin, got %f", dist)
	}
}

func TestMinDistToBinPointOutsideLeft(t *testing.T) {
	g := Grid{
		num_x_bins:  10,
		num_y_bins:  10,
		x_start:     0.0,
		x_end:       100.0,
		y_start:     0.0,
		y_end:       100.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
	}

	// Point to the left of bin (1, 1): [10, 20) x [10, 20)
	// Point at (5.0, 15.0), should be 5.0 units to the left
	dist := g.MinDistToBind(1, 1, 5.0, 15.0)
	expected := 5.0
	if dist != expected {
		t.Errorf("Expected distance %f for point to the left, got %f", expected, dist)
	}
}

func TestMinDistToBinPointOutsideRight(t *testing.T) {
	g := Grid{
		num_x_bins:  10,
		num_y_bins:  10,
		x_start:     0.0,
		x_end:       100.0,
		y_start:     0.0,
		y_end:       100.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
	}

	// Point to the right of bin (1, 1): [10, 20) x [10, 20)
	// Point at (25.0, 15.0), should be 5.0 units to the right
	dist := g.MinDistToBind(1, 1, 25.0, 15.0)
	expected := 5.0
	if dist != expected {
		t.Errorf("Expected distance %f for point to the right, got %f", expected, dist)
	}
}

func TestMinDistToBinPointOutsideBelow(t *testing.T) {
	g := Grid{
		num_x_bins:  10,
		num_y_bins:  10,
		x_start:     0.0,
		x_end:       100.0,
		y_start:     0.0,
		y_end:       100.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
	}

	// Point below bin (1, 1): [10, 20) x [10, 20)
	// Point at (15.0, 5.0), should be 5.0 units below
	dist := g.MinDistToBind(1, 1, 15.0, 5.0)
	expected := 5.0
	if dist != expected {
		t.Errorf("Expected distance %f for point below, got %f", expected, dist)
	}
}

func TestMinDistToBinPointOutsideAbove(t *testing.T) {
	g := Grid{
		num_x_bins:  10,
		num_y_bins:  10,
		x_start:     0.0,
		x_end:       100.0,
		y_start:     0.0,
		y_end:       100.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
	}

	// Point above bin (1, 1): [10, 20) x [10, 20)
	// Point at (15.0, 25.0), should be 5.0 units above
	dist := g.MinDistToBind(1, 1, 15.0, 25.0)
	expected := 5.0
	if dist != expected {
		t.Errorf("Expected distance %f for point above, got %f", expected, dist)
	}
}

func TestMinDistToBinPointAtCorner(t *testing.T) {
	g := Grid{
		num_x_bins:  10,
		num_y_bins:  10,
		x_start:     0.0,
		x_end:       100.0,
		y_start:     0.0,
		y_end:       100.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
	}

	tests := []struct {
		name     string
		xbin     int
		ybin     int
		x        float64
		y        float64
		expected float64
	}{
		{
			name:     "Bottom-left corner",
			xbin:     1,
			ybin:     1,
			x:        5.0,               // 5 units to the left of x_min=10
			y:        5.0,               // 5 units below y_min=10
			expected: 7.071067811865476, // sqrt(5^2 + 5^2) = 5*sqrt(2)
		},
		{
			name:     "Top-right corner",
			xbin:     1,
			ybin:     1,
			x:        23.0, // 3 units to the right of x_max=20
			y:        24.0, // 4 units above y_max=20
			expected: 5.0,  // sqrt(3^2 + 4^2) = 5
		},
		{
			name:     "Bottom-right corner",
			xbin:     2,
			ybin:     2,
			x:        36.0, // 6 units to the right of x_max=30
			y:        12.0, // 8 units below y_min=20
			expected: 10.0, // sqrt(6^2 + 8^2) = 10
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dist := g.MinDistToBind(tt.xbin, tt.ybin, tt.x, tt.y)
			if math.Abs(dist-tt.expected) > threshold {
				t.Errorf("Expected distance %f, got %f", tt.expected, dist)
			}
		})
	}
}

func TestMinDistToBinInvalidBins(t *testing.T) {
	g := Grid{
		num_x_bins:  10,
		num_y_bins:  10,
		x_start:     0.0,
		x_end:       100.0,
		y_start:     0.0,
		y_end:       100.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
	}

	tests := []struct {
		name string
		xbin int
		ybin int
	}{
		{"Negative xbin", -1, 5},
		{"Negative ybin", 5, -1},
		{"xbin too large", 10, 5},
		{"ybin too large", 5, 10},
		{"Both negative", -1, -1},
		{"Both too large", 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dist := g.MinDistToBind(tt.xbin, tt.ybin, 50.0, 50.0)
			if !math.IsInf(dist, -1) {
				t.Errorf("Expected -Inf for invalid bin indices, got %f", dist)
			}
		})
	}
}

func TestMinDistToBinWithNegativeCoordinates(t *testing.T) {
	g := Grid{
		num_x_bins:  10,
		num_y_bins:  10,
		x_start:     -50.0,
		x_end:       50.0,
		y_start:     -50.0,
		y_end:       50.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
	}

	// Bin (5, 5) covers [0, 10) x [0, 10)
	// Point at (5.0, 5.0) is inside
	dist := g.MinDistToBind(5, 5, 5.0, 5.0)
	if dist != 0.0 {
		t.Errorf("Expected distance 0.0 for point inside bin, got %f", dist)
	}

	// Point at (-55.0, 5.0) is 5 units to the left of bin (0, 5) which covers [-50, -40) x [0, 10)
	dist = g.MinDistToBind(0, 5, -55.0, 5.0)
	expected := 5.0
	if dist != expected {
		t.Errorf("Expected distance %f, got %f", expected, dist)
	}
}

func TestLinearScanEmptyGrid(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	// Search in empty grid
	result := g.LinearScan(15.0, 15.0)
	if result != nil {
		t.Error("Expected LinearScan to return nil for empty grid")
	}
}

func TestLinearScanSinglePoint(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	// Insert single point
	g.Insert(5.0, 5.0)

	// Search should find the only point
	result := g.LinearScan(7.0, 7.0)
	if result == nil {
		t.Fatal("Expected LinearScan to find a point")
	}

	if result.x != 5.0 || result.y != 5.0 {
		t.Errorf("Expected to find point (5.0, 5.0), got (%f, %f)", result.x, result.y)
	}
}

func TestLinearScanMultiplePointsSameBin(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	// Insert multiple points in the same bin
	g.Insert(2.0, 2.0)
	g.Insert(5.0, 5.0)
	g.Insert(8.0, 8.0)

	// Search for closest point to (6.0, 6.0)
	result := g.LinearScan(6.0, 6.0)
	if result == nil {
		t.Fatal("Expected LinearScan to find a point")
	}

	// The closest point should be (5.0, 5.0)
	// Distance from (6, 6) to (5, 5) = sqrt(2) ≈ 1.414
	// Distance from (6, 6) to (8, 8) = sqrt(8) ≈ 2.828
	// Distance from (6, 6) to (2, 2) = sqrt(32) ≈ 5.657
	if result.x != 5.0 || result.y != 5.0 {
		t.Errorf("Expected to find closest point (5.0, 5.0), got (%f, %f)", result.x, result.y)
	}
}

func TestLinearScanMultiplePointsDifferentBins(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	// Insert points in different bins
	g.Insert(5.0, 5.0)   // Bin (0, 0)
	g.Insert(15.0, 15.0) // Bin (1, 1)
	g.Insert(25.0, 25.0) // Bin (2, 2)

	// Search for closest point to (14.0, 14.0)
	result := g.LinearScan(14.0, 14.0)
	if result == nil {
		t.Fatal("Expected LinearScan to find a point")
	}

	// The closest point should be (15.0, 15.0)
	// Distance from (14, 14) to (15, 15) = sqrt(2) ≈ 1.414
	// Distance from (14, 14) to (5, 5) = sqrt(162) ≈ 12.728
	// Distance from (14, 14) to (25, 25) = sqrt(242) ≈ 15.556
	if result.x != 15.0 || result.y != 15.0 {
		t.Errorf("Expected to find closest point (15.0, 15.0), got (%f, %f)", result.x, result.y)
	}
}

func TestLinearScanExactMatch(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	// Insert points
	g.Insert(5.0, 5.0)
	g.Insert(15.0, 15.0)

	// Search for exact match
	result := g.LinearScan(5.0, 5.0)
	if result == nil {
		t.Fatal("Expected LinearScan to find a point")
	}

	// Should find exact match with distance 0
	if result.x != 5.0 || result.y != 5.0 {
		t.Errorf("Expected to find exact match (5.0, 5.0), got (%f, %f)", result.x, result.y)
	}
}

func TestLinearScanQueryOutsideGrid(t *testing.T) {
	g := Grid{
		num_x_bins:  3,
		num_y_bins:  3,
		x_start:     0.0,
		x_end:       30.0,
		y_start:     0.0,
		y_end:       30.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 3),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 3)
	}

	// Insert a point
	g.Insert(5.0, 5.0)

	// Query point outside grid bounds
	result := g.LinearScan(50.0, 50.0)
	if result == nil {
		t.Fatal("Expected LinearScan to find a point even with query outside grid")
	}

	// Should still find the closest point in the grid
	if result.x != 5.0 || result.y != 5.0 {
		t.Errorf("Expected to find point (5.0, 5.0), got (%f, %f)", result.x, result.y)
	}
}

func TestLinearScanAllCorners(t *testing.T) {
	g := Grid{
		num_x_bins:  2,
		num_y_bins:  2,
		x_start:     0.0,
		x_end:       20.0,
		y_start:     0.0,
		y_end:       20.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 2),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 2)
	}

	// Insert points in all four corners
	g.Insert(1.0, 1.0)   // Bottom-left
	g.Insert(19.0, 1.0)  // Bottom-right
	g.Insert(1.0, 19.0)  // Top-left
	g.Insert(19.0, 19.0) // Top-right

	tests := []struct {
		name      string
		queryX    float64
		queryY    float64
		expectedX float64
		expectedY float64
	}{
		{"Center should find any point", 10.0, 10.0, -1.0, -1.0}, // Don't check exact, just that something is found
		{"Near bottom-left", 2.0, 2.0, 1.0, 1.0},
		{"Near bottom-right", 18.0, 2.0, 19.0, 1.0},
		{"Near top-left", 2.0, 18.0, 1.0, 19.0},
		{"Near top-right", 18.0, 18.0, 19.0, 19.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.LinearScan(tt.queryX, tt.queryY)
			if result == nil {
				t.Error("Expected LinearScan to find a point")
				return
			}

			// For center test, just verify a point was found
			if tt.expectedX == -1.0 && tt.expectedY == -1.0 {
				return
			}

			if result.x != tt.expectedX || result.y != tt.expectedY {
				t.Errorf("Expected to find point (%f, %f), got (%f, %f)",
					tt.expectedX, tt.expectedY, result.x, result.y)
			}
		})
	}
}

func TestLinearScanWithManyPoints(t *testing.T) {
	g := Grid{
		num_x_bins:  5,
		num_y_bins:  5,
		x_start:     0.0,
		x_end:       50.0,
		y_start:     0.0,
		y_end:       50.0,
		x_bin_width: 10.0,
		y_bin_width: 10.0,
		bins:        make([][]*GridPoint, 5),
	}

	// Initialize bins
	for i := range g.bins {
		g.bins[i] = make([]*GridPoint, 5)
	}

	// Insert many points
	points := []struct{ x, y float64 }{
		{2.5, 2.5}, {7.5, 2.5}, {12.5, 12.5},
		{22.5, 22.5}, {32.5, 32.5}, {42.5, 42.5},
		{5.0, 15.0}, {25.0, 35.0}, {45.0, 5.0},
	}

	for _, p := range points {
		g.Insert(p.x, p.y)
	}

	// Query for a point and verify we get something reasonable
	result := g.LinearScan(10.0, 10.0)
	if result == nil {
		t.Fatal("Expected LinearScan to find a point")
	}

	// Verify it's one of the inserted points
	found := false
	for _, p := range points {
		if result.x == p.x && result.y == p.y {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Result point (%f, %f) is not one of the inserted points", result.x, result.y)
	}

	// The closest point to (10, 10) should be (12.5, 12.5)
	// Let's verify by calculating distances
	minDist := math.Inf(1)
	var closestPoint *struct{ x, y float64 }
	for i := range points {
		dist := euclidean_dist(10.0, 10.0, points[i].x, points[i].y)
		if dist < minDist {
			minDist = dist
			closestPoint = &points[i]
		}
	}

	if result.x != closestPoint.x || result.y != closestPoint.y {
		t.Errorf("Expected closest point (%f, %f), got (%f, %f)",
			closestPoint.x, closestPoint.y, result.x, result.y)
	}
}
