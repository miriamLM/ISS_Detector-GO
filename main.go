package main

import (
	"./issnow"
	"./wheretheiss"
	"fmt"
	"time"
)

func InfiniteLoop() {
	c := make(chan string)
	d := make(chan string)
	var changeIssNow = ""
	var changeWhereTheIss = ""
	for true {
		time.Sleep(time.Second)

		go issnow.PositionIssNow(c)
		go wheretheiss.PositionWhereTheIss(d)

		countryIssNow := <-c
		countryWhereTheIss := <-d

		if changeIssNow != countryIssNow && countryIssNow != "Water" {
			fmt.Println("issNow has changed country to: " + countryIssNow)
			changeIssNow = countryIssNow
		}
		if countryWhereTheIss == "??" {
			countryWhereTheIss = changeWhereTheIss
		}
		if changeWhereTheIss != countryWhereTheIss && countryWhereTheIss != "Water" && countryWhereTheIss != "??" {
			fmt.Println("whereTheIss has changed country to: " + countryWhereTheIss)
			changeWhereTheIss = countryWhereTheIss
		}

		fmt.Println("issNow: " + countryIssNow + " and whereTheIss: " + countryWhereTheIss)

		time.Sleep(time.Second * 10)
	}
}
func main() {
	InfiniteLoop()
}
