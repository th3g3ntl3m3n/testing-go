package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var counterInt int

var mu sync.RWMutex

func main() {

	counterInt = 0

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/", index)

	router.GET("/counter/:number", counter)
	fmt.Println(Greet("Git"))

	PORT := os.Getenv("PORT")

	router.Run(PORT)

}

func index(ctx *gin.Context) {
	// ctx.File("main.html")
	mu.Lock()
	counterInt += 1
	mu.Unlock()
	log.Println("Log value is ", counterInt)
	PORT := os.Getenv("PORT")
	ctx.String(200, "Hello World FROM NEW SERVER"+PORT)
}

func counter(ctx *gin.Context) {

	b := ctx.Params.ByName("number")
	y, err := strconv.Atoi(b)
	if err != nil {
		log.Println("Error converting into integer", b)
		ctx.String(http.StatusBadRequest, "bad request")
		return
	}
	x := fmt.Sprintf("There you go %d FIB is %d\n", y, fib(y))
	ctx.String(200, x)
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

func ToUpper(text string) string {
	var newString = ""
	for i := 0; i < len(text); i++ {
		if text[i] >= 97 && text[i] <= (97+26) {
			newString += (string(text[i] - (97 - 65)))
		} else {
			newString += (string(text[i]))
		}
	}
	return newString
}
