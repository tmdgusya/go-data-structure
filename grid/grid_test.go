package grid

import (
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
