package tests
import(
	"fmt"
)
type HumanPlayer struct{
	IsBusted bool
	Cards []Card
}
func (h *HumanPlayer)plays(){
	fmt.Println("Human Played")
}
func (h *HumanPlayer)isBusted() bool{
	return h.IsBusted
}
func (h *HumanPlayer)GetCards() []Card{
	return h.Cards
}