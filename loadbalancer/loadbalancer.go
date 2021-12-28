package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LOAD BALANCER
// ROUTE REQ to available services
var services = []string{"http://localhost:9091", "http://localhost:9092", "http://localhost:9093"}

var counterInt = 0

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Any("/", handleLoad)

	router.Run(":8089")

	// sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// for {
	// 	select {
	// 	case <-sigChan:
	// 		log.Fatal("Received OS Interrupt Signal ... ")
	// 	}
	// }

}

func handleLoad(c *gin.Context) {
	client := &http.Client{}
	req, err := http.NewRequest(c.Request.Method, services[counterInt], c.Request.Body)
	if err != nil {
		log.Println("Err sending req", err, services[counterInt])
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Err sending req 2", err, services[counterInt])
	}

	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	counterInt++
	counterInt = counterInt % 3

	c.String(res.StatusCode, string(resp))
}
