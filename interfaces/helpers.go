package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	"reflect"
	"runtime"
	"strings"
)

func name(f func(g *gin.Context))(controllerName string, functionName string){
	path := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), "/")
	theName := strings.Split(path[len(path)-1], ".")
	controllerName = strcase.ToKebab(theName[len(theName)-2])
	functionName = strcase.ToKebab(theName[len(theName)-1])
	if strings.Contains(functionName, "-fm") {
		functionName = strings.ReplaceAll(functionName, "-fm", "")
	}
	if strings.Contains(controllerName, "*") {
		controllerName = strings.ReplaceAll(controllerName, "*", "")
	}
	if strings.Contains(controllerName, "(") {
		controllerName = strings.ReplaceAll(controllerName, "(", "")
	}
	if strings.Contains(controllerName, ")") {
		controllerName = strings.ReplaceAll(controllerName, ")", "")
	}
	controllerName = strings.Split(controllerName, "-")[0]
	return
}

func urlString(function func(g *gin.Context), params... string) string  {
	paramsString := ""
	for _, param := range params{
		paramsString += "/:" + param
	}
	controllerName , functionName := name(function)
	return controllerName+"/"+functionName+paramsString
}

func GET(r *gin.RouterGroup,function func(g *gin.Context), params... string) *gin.RouterGroup{
	//fmt.Println(urlString(function, params...))
	r.GET(urlString(function, params...), function)
	return r
}

func POST(r *gin.RouterGroup,function func(g *gin.Context), params... string) *gin.RouterGroup{
	r.POST(urlString(function, params...), function)
	return r
}

func DELETE(r *gin.RouterGroup,function func(g *gin.Context), params... string) *gin.RouterGroup{
	r.DELETE(urlString(function, params...), function)
	return r
}

func Any(r *gin.RouterGroup,function func(g *gin.Context), params... string) *gin.RouterGroup{
	r.Any(urlString(function, params...), function)
	return r
}

func HEAD(r *gin.RouterGroup,function func(g *gin.Context), params... string) *gin.RouterGroup{
	r.HEAD(urlString(function, params...), function)
	return r
}

func OPTIONS(r *gin.RouterGroup,function func(g *gin.Context), params... string) *gin.RouterGroup{
	r.OPTIONS(urlString(function, params...), function)
	return r
}

func PATCH(r *gin.RouterGroup,function func(g *gin.Context), params... string) *gin.RouterGroup{
	r.PATCH(urlString(function, params...), function)
	return r
}

func PUT(r *gin.RouterGroup,function func(g *gin.Context), params... string) *gin.RouterGroup{
	r.PUT(urlString(function, params...), function)
	return r
}

