package tests
import{
	"fmt"
}
type HumanPlayer struct{
	isBusted bool
}
func (h *HumanPlayer)plays(){
	fmt.Println("Human Played")
}
func (h *HumanPlayer)isBusted() bool{
	return h.bool
}