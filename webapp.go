package main

import (
	"math/rand"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

func randomInt() int {
	return rand.Intn(7)
}

var arrOfAttempts = make([]int, 0)
var countWins int = 0
var countOfAttempts int = 1

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
		arrOfAttempts = append(arrOfAttempts, countOfAttempts)
		countOfAttempts = 1
		return
	} else if num > result {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "ваше число больше загаданного"})
		countOfAttempts++
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "ваше число меньше загаданного"})
		countOfAttempts++
		return
	}

}

func numberOfWins(c *gin.Context) {
	if len(arrOfAttempts) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"ошибка": "ещё не было попыток"})
		return
	}
	sort.Ints(arrOfAttempts)

	c.IndentedJSON(http.StatusOK, gin.H{"лучший результат (наименьшее количество попыток)": arrOfAttempts[0], "общее количество побед": countWins})

}

func main() {
	router := gin.Default()
	router.GET("/game/:num", checkingTheNnumber)
	router.GET("/wins", numberOfWins)
	router.GET("/newNumber", newNumber)

	router.Run("localhost:8080")
}
