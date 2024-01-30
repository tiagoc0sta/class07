package main

import (
	"fmt"
)

func main() {
	checkBalance()
	checkAge()
	checkTemp()
	checkScore()
	checkRain()
	checkMoney()
	checkBook()
	destination()
	watching()
	checkWeatherMoney()
	happy()
}

func checkBalance(){ 
	balance := 100.0

 if balance < 0 {
  fmt.Println("Balance is negative.")
 } else {
  fmt.Println("Balance is positive.")
 }
}

func checkAge() {
	age := 65

	if age > 65 {
		fmt.Println("You can Retire")
	} else {
		fmt.Println("You cannot Retire")
	}
}

func checkTemp(){
	temperature := 25.0


	if temperature > 30.0 {
		fmt.Println("Is hot")
	} else {
		fmt.Println("It's hot")
	}
}

func checkScore(){
	score := 85

	if score >= 70 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}

func checkRain() {
	isRaining := false

	if isRaining == true {
		fmt.Println("It's raining")
	}else {
		fmt.Println("It's not raining")
	}
}

func checkMoney() {

	money := 15
 
	if money == 5 {
	 fmt.Println("Vanilla")
	} else if money == 10 {
	 fmt.Println("Chocolate")
	} else if money == 20 {
	 fmt.Println("Strawberry")
	} else {
	 fmt.Println("Bubblegum")
	}
 
 }

 func checkBook() {

	teaFlavor := "chamomile"
	bookGenre := "mystery"
 
	if teaFlavor == "lavender" && bookGenre == "mystery" {
	 fmt.Println("Sipping lavender tea while reading a mystery novel is perfect")
	} else if teaFlavor == "chamomile" && bookGenre == "romance" {
	 fmt.Println("Enjoy your chamomile tea with a romantic tale")
	} else if teaFlavor == "matcha" && bookGenre == "sci-fi" {
	 fmt.Println("Matcha tea pairs well with an engaging sci-fi story")
	} else {
	 fmt.Println("Choose peppermint tea with self development piece")
	}
 
 }

 func destination() {

	destination := "Japan"
	season := "winter"
 
	if destination == "Japan" || season == "spring" {
	 fmt.Println("Enjoy the beautiful cherry blossoms in Japan this spring")
	} else if destination == "Iceland" && season == "winter" {
	 fmt.Println("Experience the magical northern lights in Iceland")
	} else if destination == "Thailand" && season == "summer" {
	 fmt.Println("Discover the beautiful beaches of Thailand")
	} else {
	 fmt.Println("Enjoy Canada in the fall")
	}
 
 }

 func watching() {

	snack := "chips"
	movieGenre := "romance"
 
	if snack == "popcorn" || movieGenre == "action" {
	 fmt.Println("Grab your popcorn for an action-packed movie night")
	} else if snack == "chips" || movieGenre == "comedy" {
 
	 fmt.Println("Chips are a great choice for a night of comedy films")
 
	} else if snack == "cookies" || movieGenre == "romance" {
 
	 fmt.Println("Enjoy your cookies while watching a romantic film")
 
	} else {
	 fmt.Println("Lettuce and study")
	}
 }

 func checkWeatherMoney() {

	money := 15
	isSunny := true
	iceCreamPrice := 5
 
	if isSunny && money >= iceCreamPrice {
 
	 fmt.Println("It's sunny and you have enough money for ice cream. Enjoy!!")
 
	} else if isSunny && money < iceCreamPrice {
 
	 fmt.Println("It is sunny but you do not have enough money for ice cream.")
 
	} else if !isSunny && money >= iceCreamPrice {
 
	 fmt.Println("It is not sunny but you have enough money for ice cream.")
 
	} else {
	 fmt.Println("no sunny, no money")
	}
 
 }

 func happy() {

	isHappy := false
 
	if isHappy {
	 fmt.Println("first block")
	} else if !isHappy {
	 fmt.Println("second block")
 
	} else {
	 fmt.Println("third block")
 
	}
 }