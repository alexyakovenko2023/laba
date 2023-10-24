package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"math/rand"
)

func randomInt() int {
	return rand.Intn(20)
}

var countWins int = 0

var result int = randomInt()

func newNumber(c *gin.Context) {
	result = randomInt()
	c.IndentedJSON(http.StatusOK, gin.H{"message": "новое число загадано"})
}

func checkingTheNnumber(c *gin.Context) {
	str := c.Param("num")
	num, err := strconv.Atoi(str)
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"ошибка": "нужно ввести число"})
		return
	}

	if num == result {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "правильно число"})
		countWins++
		return
	} else if num > result {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "ваше число больше загаданного"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "ваше число меньше загаданного"})
		return
	}

}

func numberOfWins(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"общее количество побед": countWins})
}

func main() {
	router := gin.Default()
	router.GET("/game/:num", checkingTheNnumber)
	router.GET("/wins", numberOfWins)
	router.GET("/newNumber", newNumber)

	router.Run("localhost:8080")
}
