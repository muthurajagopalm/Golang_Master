package main

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

type Fruit struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {

	router := gin.Default()

	fruits := map[string]float64{
		"Orange":     60.5,
		"Banana":     30.0,
		"Apple":      80.0,
		"Grapes":     55.2,
		"Pineapple":  120.0,
		"Watermelon": 90.5,
	}

	router.GET("fruits/sort/name", func(ctx *gin.Context) {
		var names []string
		for name := range fruits {
			names = append(names, name)
			fmt.Println(names)

		}
		sort.Strings(names)

		var sorted []Fruit
		for _, name := range names {
			sorted = append(sorted, Fruit{Name: name, Price: fruits[name]})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"sorted_by": "name",
			"data":      sorted,
		})
	})

	router.GET("/fruits/sort/price", func(ctx *gin.Context) {
		var fruitList []Fruit
		for fruit_k, fruit_val := range fruits {
			fruitList = append(fruitList, Fruit{Name: fruit_k, Price: fruit_val})
		}

		sort.Slice(fruitList, func(i, j int) bool {
			return fruitList[i].Price < fruitList[j].Price
		})

		type Response struct {
			SortedBy string  `json:"sorted_by"`
			Data     []Fruit `json:"data"`
		}

		ctx.JSON(http.StatusOK, Response{
			SortedBy: "price",
			Data:     fruitList,
		})
	})

	router.Run(":8080")

}
