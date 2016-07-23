package tests
const MAX_VALUE=21
type BlackJackGameDecider struct{
	HumanPlayer,ComputerPlayer Player
}
func (b *BlackJackGameDecider) Whowins() Player{
	computerScore:=getTotalScore(ComputerPlayer.GetCards())
	humanScore:=getTotalScore(HumanPlayer.GetCards())
	if computerScore>humanScore{
		return ComputerPlayer
	}else if humanScore>computerScore{
		return HumanPlayer
	}else{
		return nil
	}
}
func getTotalScore(cards []Card)int{
	score:=0;
	for _,card:= range cards{
		score+=card.GetValue()
	}
}