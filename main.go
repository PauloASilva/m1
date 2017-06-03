package m1

import (
	"github.com/PauloASilva/m1/controllers"
	"net/http"
)

func Register(router *http.ServeMux) {
	controllers.HomeRegister(router)
	controllers.AboutRegister(router)
}
