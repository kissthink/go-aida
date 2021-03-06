package main
import (
	"github.com/gin-gonic/gin"
	"github.com/songtianyi/go-aida/restful/logic"
)

func main() {
	router := gin.Default()

	router.GET("/create", logic.Create)
	router.GET("/status", logic.Status)

	router.Run();
}
