package counter

// Counter counts the calls made to it. The step log ADR-0001 asks for is
// not implemented yet: it is the gap the loop will close.
type Counter struct {
	n int
}

// Inc adds one to the tally.
func (c *Counter) Inc() { c.n++ }

// Tally is how many times Inc has been called.
func (c *Counter) Tally() int { return c.n }
