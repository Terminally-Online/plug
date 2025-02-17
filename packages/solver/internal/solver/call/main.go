package call

type Caller struct{}

func New() Caller {
	return Caller{}
}

func (c *Caller) Call() error {
	return nil
}
