package myquote

import "rsc.io/quote"

func MyGlass() string {
	myGlass := quote.Glass()
	return myGlass
}

func MyGo() string {
	return quote.Go()
}

func MyHello() string {
	return quote.Hello()
}

func MyOpt() string {
	return quote.Opt()
}
