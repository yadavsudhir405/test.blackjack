package tests
type BlackJackGameDecider struct{
	Players []Player
}
func (b *BlackJackGameDecider) Whowins() Player{
	for _,player :=range b.Players{
		if player.isBusted()==false {
			return player
		}
	}
	return nil
}