package main

/* FILE DOCSTRINGS

Major help came from these sources:
  - Originally was going to use webassembly (*1)
  - A good portion was learned from youtube (*2)
  - Pieces were gathered from my classmates (*3)
  - Echo was a major resource in this go project
  - And of course Dani, Exercisms, and Go devs!!

*1. https://tutorialedge.net/golang/go-webassembly-tutorial/
*2. https://www.youtube.com/watch?v=_pww3NJuWnk
*3. https://github.com/fchikwekwe/FaithBot
*4. https://echo.labstack.com/guide/templates
*5. https://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html

END OF DOCSTRINGS */

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// index route applies for `/`
func index(context echo.Context) error {
	fmt.Println("Got to INDEX route")
	return context.File("./html/index.html")
}

// roll function will route to `/roll`
func roll(context echo.Context) error {
	var dice = context.QueryParam("dice")
	// Roller "./roller"
	// var roll = strconv.Itoa(Roller.RollDice(dice))
	roll := "12"
	return context.String(
		http.StatusOK,
		fmt.Sprintf("Your dice: %s\nYour roll: %s",
			dice,
			roll))
}

// main gets echo server up and running
func main() {
	fmt.Println("Started Main Function")
	// port is an env var
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// start echo router
	router := echo.New()
	router.GET("/", index)
	router.GET("/roll", roll)
	fmt.Println("Now listening to port...")
	router.Start(port)
}
