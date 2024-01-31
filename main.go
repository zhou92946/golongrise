package main

import (
	"encoding/json"
	"fmt"
	"goasset/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("go asset")
	r := gin.Default()

	brokers := []model.Broker{}
	var i int32 = 1
	for ; i < 1000; i++ {
		broker := model.Broker{Key: string(i), Name: "西南期货", Id: i, Address: "111", RegisterTime: "2023-01-01"}
		brokers = append(brokers, broker)
	}

	r.GET("/", func(c *gin.Context) {
		ret, err := json.Marshal(brokers)
		if err != nil {
			c.String(http.StatusOK, "err")
			return
		}
		c.String(http.StatusOK, string(ret))
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
