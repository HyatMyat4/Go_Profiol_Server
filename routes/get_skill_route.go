package routes

import (
	controllers "go-profiol-server/controllers"

	"github.com/gin-gonic/gin"
)

func Get_allskill_route(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/skills", controllers.Get_AllSkill())
}
