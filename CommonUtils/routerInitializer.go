package CommonUtils

import (
	"github.com/gin-gonic/gin"
)

func RouterInitializer() {
	router := gin.Default()
	// calling routerCollector from urls.go
	routerCollector(router)
	listen(router, "7777")

}

func listen(router *gin.Engine, port string) {
	router.Run("localhost:" + port)
}
