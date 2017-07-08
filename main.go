package m1

import (
	"github.com/PauloASilva/m1/controllers"
	"net/http"
)

var Routes = map[string]func(w http.ResponseWriter, req *http.Request){
	"/":      controllers.Home,
	"/about": controllers.About,
}
