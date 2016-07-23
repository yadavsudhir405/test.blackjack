package tests

import "testing"

func Test(t *testing.T) {
	humanPlayer:=new(HumanPlayer)
	computerPlayer:=new(ComputerPlayer)
	humanPlayer.isBusted=true
	computerPlayer.isBusted=false
	blackjackGameDecider:=&blackjackGameDecider{
		Players:[]Player{&computerPlayer,&humanPlayer}
	}
	winner := blackjackGameDecider.Whowins()
	if result != nil {
		t.Errorf("Expected winner, is  %d", nil)
	}
}

