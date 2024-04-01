package main

import (
	"github.com/gin-gonic/gin"
	"math/rand/v2"
	"net/http"
	"strconv"
)

/*
Serves as the main controller for the Number Guesser program. This program randomly chooses a number between 0-100 and
allows the user to guess what the chosen number is. After each guess, they are told whether their guess is greater than
or less than the chosen number. Once the correct number is guessed, the user is moved to the success page, where they
receive the option to play again if desired.

Author: Ryan Johnson
*/

var chosenNum int

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("", mainGET)
	router.GET("/:numGuess", func(c *gin.Context) {
		guessGET(c, chosenNum)
	})

	err := router.Run()
	if err != nil {
		return
	}
}

/*
*
Called at the beginning of each new game. A new random number is generated, and the user is taken to the main page.
*/
func mainGET(c *gin.Context) {
	chosenNum = rand.IntN(100)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

/*
Called each time a user makes a guess as to what the chosen number is. Depending upon whether the guess is higher than,
lower than, or equal to the chosen number, they are taken to the corresponding HTML template displaying the correct
information.
*/
func guessGET(c *gin.Context, chosenNum int) {
	guess := c.Query("numGuess")
	numGuess, _ := strconv.Atoi(guess)

	if numGuess < chosenNum {
		c.HTML(http.StatusOK, "guessLow.tmpl", gin.H{
			"numGuess": numGuess,
		})
	} else if numGuess > chosenNum {
		c.HTML(http.StatusOK, "guessHigh.tmpl", gin.H{
			"numGuess": numGuess,
		})
	} else {
		c.HTML(http.StatusOK, "won.tmpl", gin.H{
			"numGuess": numGuess,
		})
	}
}
