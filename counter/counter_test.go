package counter

import "testing"

func TestIncCounts(t *testing.T) {
	var c Counter
	c.Inc()
	c.Inc()
	if got := c.Tally(); got != 2 {
		t.Fatalf("Tally = %d, want 2", got)
	}
}
