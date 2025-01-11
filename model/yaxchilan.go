package model

func Yaxchilan0(p *Player) {
	return
}

func Yaxchilan1(p *Player) {
	p.resources[Wood] += 1
}

func Yaxchilan2(p *Player) {
	p.resources[Stone] += 1
	p.corn += 1
}

func Yaxchilan3(p *Player) {
	p.resources[Gold] += 1
	p.corn += 2
}

func Yaxchilan4(p *Player) {
	p.resources[Skull] += 1
}

func Yaxchilan5(p *Player) {
	p.resources[Gold] += 1
	p.resources[Stone] += 1
	p.corn += 2
}

func Yaxchilan() [][]Option {
	return [][]Option{
		[]Option{PlayerOption(Yaxchilan0)},
		[]Option{PlayerOption(Yaxchilan1)},
		[]Option{PlayerOption(Yaxchilan2)},
		[]Option{PlayerOption(Yaxchilan3)},
		[]Option{PlayerOption(Yaxchilan4)},
		[]Option{PlayerOption(Yaxchilan5)},
	}
}

func MakeYaxchilan() *Wheel {
	return MakeWheel(Yaxchilan(), 2, "Yaxchilan")
}