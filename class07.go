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
	salaryDate()	
	//energy()
	birthday()
	quiz01()
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

 func salaryDate() {

	salary := 10
	dateIdea := "surprise world tour"
 
	if salary >= 1000 && (dateIdea == "bowling" || dateIdea == "steak dinner" || dateIdea == "skydiving") {
 
	 fmt.Println("With your salary, you can afford a fun date")
 
	} else if salary >= 1000 && dateIdea == "surprise world tour" {
 
	 fmt.Println("You can afford a surprise world tour")
 
	} else if salary >= 5000 && dateIdea == "luxury world tour" {
 
	 fmt.Println("You can afford a luxury world tour")
 
	} else if salary >= 1000 {
 
	 fmt.Println("You have enough for a nice date")
 
	} else {
	 fmt.Println("heheheheh no")
	}
 
 }

 func energy() {

	var energyLevel int
	var bucketListItem string
 
	fmt.Print("Enter your energy level (0-100)")
 
	fmt.Scanln(&energyLevel)
 
	fmt.Print("Enter your bucket list item (skydiving, marathon, olympics, world tour, luxury world tour)")
 
	fmt.Scanln(&bucketListItem)
 
	if energyLevel >= 70 && (bucketListItem == "skydiving" || bucketListItem == "marathon" || bucketListItem == "olympics") {
	 fmt.Println("With your energy level, you can plan several active bucket list items")
	} else if energyLevel >= 70 && bucketListItem == "world tour" {
	 fmt.Println("With your energy level, you can go on a world tour")
	} else if energyLevel >= 80 && bucketListItem == "luxury world tour" {
	 fmt.Println("With your energy level, you can go on a luxury world tour")
	} else if energyLevel >= 70 {
	 fmt.Println("With your energy level, you can decide whatever item you would like")
	} else {
	 fmt.Println("Your energy level is too low, get some rest!!")
	}
 
 }

 func birthday(){
	hasCake := true
	holiday := "birthday"

	if hasCake && holiday == "birthday" {
			fmt.Println("You have cake on your birthday! Enjoy the celebration.")
	} else if hasCake && holiday == "anniversary" {
			fmt.Println("You have cake on your anniversary! Celebrate your special day.")
	} else if hasCake && holiday == "Christmas" {
			fmt.Println("You have cake for Christmas! Savor the festive spirit.")
	} else if hasCake {
			fmt.Println("You have cake for an unspecified holiday. Enjoy the sweet surprise!")
	} else {
			fmt.Println("There's no cake for this holiday. Consider getting one to celebrate.")
	}
 }


 func quiz01(){
	
	// Get input from the user or set them programmatically
	var budget int
	var flowerChoice string

	fmt.Print("Enter your budget: ")
	fmt.Scan(&budget)

	fmt.Print("Enter your flower choice: ")
	fmt.Scan(&flowerChoice)

	// Conditions 
	if budget < 1000 {
		fmt.Println("Your budget is too low for the mentioned types of flowers. Consider opting for more budget-friendly options.")
	} else if budget >= 1000 && budget < 5000 {
		switch flowerChoice {
		case "roses", "tulips", "lilies":
			fmt.Println("With your budget, you can afford beautiful flowers. Enjoy your choice!")
		case "exotic orchids":
			fmt.Println("With your budget, you can go for exotic orchids. They'll make a unique and impressive gift!")
		default:
			fmt.Println("You have enough for a wonderful choice of flowers, but the type is unspecified. Choose your flowers and enjoy!")
		}
	} else if budget >= 5000 && flowerChoice == "luxury flower arrangement" {
		fmt.Println("You can afford a luxury flower arrangement for a truly special occasion. Make it memorable!")
	} else {
		fmt.Println("Invalid input. Please check your budget and flower choice.")
	}
}

 
 