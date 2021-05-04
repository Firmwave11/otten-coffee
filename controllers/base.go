package controllers

import (
	"net/http"

	"github.com/Firmwave11/otten-coffee/usecase"
)

// ctrl struct with value interface Usecases
type ctrl struct {
	uc usecase.Usecases
}

// Controllers represent the Controllers contract
type Controllers interface {
	Tracking(w http.ResponseWriter, r *http.Request)
}

/*NewCtrl will create an object that represent the Controllers interface (Controllers)
 * @parameter
 * r - Repository Interface
 *
 * @represent
 * interface Controllers
 *
 * @return
 * uc struct with value interface Usecases
 */
func NewCtrl(u usecase.Usecases) Controllers {
	return &ctrl{uc: u}
}
