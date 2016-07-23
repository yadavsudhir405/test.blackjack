package tests
type Card struct{
	Value int
	Name string
}
func (c *Card) Getvalue()int {
	return c.Value
}
func (c *Card) GetName()string {
	return c.Name
}
