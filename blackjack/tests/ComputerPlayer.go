package tests
import(
	"fmt"
)
type ComputerPlayer struct{
	IsBusted bool
	Cards []Card
}
func (h *ComputerPlayer)plays(){
	fmt.Println("ComputerPlayer Played")
}
func (h *ComputerPlayer)isBusted() bool{
	return h.IsBusted
}
func (h *ComputerPlayer)GetCards()[]Card{
	return h.Cards
}