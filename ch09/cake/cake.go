package cake

type Cake struct { state string }

func Baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake // Baker never touches it again
	}
}

func Icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake // Icer never touches it again
	}
}
