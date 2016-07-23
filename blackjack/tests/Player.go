package tests
type Player interface{
	plays()
	IsBusted() bool
	GetCards() []Card
}