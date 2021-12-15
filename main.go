package main

import (
	"net/http"
	"context"
    "fmt"
    finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
	"github.com/gin-gonic/gin"
	"encoding/json"

)

func main() {
	router := gin.Default()
	router.GET("/", foo)
	router.GET("/test1", test1)
	router.Run(":5000")
}

func foo(c *gin.Context) {
	c.Header("Content-Type","application/xml")
	c.String(http.StatusOK, "<Response><Message>bar</Message></Response>")
}
func test1(c *gin.Context) {
	fmt.Printf("ZZZZZ","%+v\n", string(c.Query("Body")))
	ticker_string := fmt.Sprintf(string(c.Query("Body")))
	fmt.Printf("QQQQQQ: \n")
	fmt.Printf(ticker_string, "\n")
	fmt.Printf("QQQQQQ: \n")
	// c.Param("Body")

	cfg := finnhub.NewConfiguration()
    cfg.AddDefaultHeader("X-Finnhub-Token", "sandbox_c6n1ihiad3ibta79o21g")
    finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi
    
    // res, _, err := finnhubClient.Quote(context.Background()).Symbol("AAPL").Execute()
	res, _, err := finnhubClient.Quote(context.Background()).Symbol(ticker_string).Execute()
	blob, _ := json.Marshal(res.C)
    fmt.Printf("%+v\n", string(blob))
	fmt.Printf("%+v\n", err)

	response_string := fmt.Sprintf(string(blob))
	c.Header("Content-Type","application/xml")
	c.String(http.StatusOK, "<Response><Message>" + response_string + "</Message></Response>")
	
}
