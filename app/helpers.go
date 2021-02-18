package app

import (
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
)

func Page(g *gin.Context) int  {
	page, _ := strconv.Atoi(g.DefaultQuery("page", "1"))
	return  page
}

func Limit(g *gin.Context) int  {
	page, _ := strconv.Atoi(g.DefaultQuery("limit", os.Getenv("LIMIT")))
	return  page
}

func Order(g *gin.Context, orders ...string)[]string   {
	var o []string

	if  g.Query("order") != ""{

		orders := strings.SplitN(g.Query("order"), "|", -1)

		for _ , order := range orders{
			o = append(o , order)
		}
		return o

	}else{

		for _ , orderBy := range orders{
			o = append(o , orderBy)
		}
		return o

	}
}

//func Order(order ...string)[]string   {
//	var o []string
//	for _ , orderBy := range order{
//		o = append(o , orderBy)
//	}
//	return o
//}