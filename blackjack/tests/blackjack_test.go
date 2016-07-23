package tests

import "testing"

func Test(t *testing.T) {
	humanPlayer:=new(HumanPlayer)
	humanPlayer.isBusted=true
	blackjackGameDecider:=&blackjackGameDecider{
		Player:humanPlayer
	}
	winner := blackjackGameDecider.Whowins()

	if result != nil {
		t.Errorf("Expected winner, is  %d", winner)
	}
}

