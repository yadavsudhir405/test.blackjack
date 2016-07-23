package tests
import{
	"fmt"
}
type ComputerPlayer struct{
	isBusted bool
}
func (h *ComputerPlayer)plays(){
	fmt.Println("ComputerPlayer Played")
}
func (h *ComputerPlayer)isBusted() bool{
	return h.bool
}