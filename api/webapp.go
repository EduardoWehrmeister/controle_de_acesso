package main

import (
	"fmt"
	"net/http"

	"controle_de_acesso.com/api/controller"
	"controle_de_acesso.com/api/http_adapter"
	"controle_de_acesso.com/api/repository"
	"controle_de_acesso.com/api/service"
)

var (
	httpRouter http_adapter.IRouter = http_adapter.NewMuxRouter()

	//Instanciar repositórios
	casbinRepository repository.ICasbinRepository = repository.NewCasbinMongoRepository()

	//Instanciar serviços
	policyService service.IPolicyService = service.NewPolicyService(casbinRepository)

	//Instanciar controller
	policyController controller.IPolicyController = controller.NewPolicyController(policyService)
)

func main() {
	//TODO - Colocar esse valor no appsettings
	const port string = ":8000"

	httpRouter.GET("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and Runnning")
	})

	httpRouter.POST("/policy", policyController.AddPolicy)

	httpRouter.SERVE(port)
}
