package files

import "fmt"

func Router(projectName *string) string {
	return fmt.Sprintf(`
package router

import (
	"%s/src/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", controllers.Ping)

	return r
}
`, *projectName)
}
