package player

func (p *Player) HandlerRegister() {
	p.Handlers["add_friend"] = p.AddFriend

}
