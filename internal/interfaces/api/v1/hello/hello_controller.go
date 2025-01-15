package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/hello"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application/service/hello/query"
)

type HelloController struct {
	helloService hello.HelloUseCase
}

func NewHelloController(helloService hello.HelloUseCase) HelloController {
	return HelloController{
		helloService: helloService,
	}
}

func (ctrl HelloController) SayHello(c *gin.Context) {
	queryResult, err := ctrl.helloService.SayHello(c.Request.Context(), query.SayHelloQuery{
		Name: "World",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, queryResult)
}
