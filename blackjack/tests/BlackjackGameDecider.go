package tests
type BlackJackGameDecider struct{
	Player Player
}
func (b *BlackJackGameDecider) Whowins()Player{
	if b.player.isBusted()==true{
		return nil
	}else{
		return b.player
	}
}