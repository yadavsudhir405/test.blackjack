package tests

import "testing"

func TestGameDecider(t *testing.T) {
	tensBlackCard:=&Card{
		Name:"TenBlack",
		Value:10,
	}
	twosBlackCard:=&Card{
		Name:"TwosBlackcard",
		Value:2,
	}
	eightsBlackCard:=&Card{
		Name:"EightBlack",
		Value:8,
	}
	sevensBlackCard:=&Card{
		Name:"SevensBlackcard",
		Value:7,
	}
	

	humansCardHolder:=[]Card{}
	humansCardHolder=append(humansCardHolder,tensBlackCard)
	humansCardHolder=append(humansCardHolder,twosBlackCard)
	
	computerCardHolder:=[]Card{}
	computerCardHolder=append(computerCardHolder,eightsBlackCard)
	computerCardHolder=append(computerCardHolder,SevensBlackcard)

	human:=&HumanPlayer{
		IsBusted:false,
		Cards:humansCardHolder,
	}

	computer:=&computerPlayer{
		IsBusted:false,
		Cards:computerCardHolder,
	}
	
	
	blackjackGameDecider:=&blackjackGameDecider{
		HumanPlayer:human,
		ComputerPlayer:computer,
		
	}
	winner := blackjackGameDecider.Whowins()
	if winner != computer {
		t.Errorf("Expected winner, is  %d", computer)
	}
}


