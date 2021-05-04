package routes

import (
	"fmt"
	"net/http"

	"github.com/Firmwave11/otten-coffee/controllers"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// route struct with value Controllers Interface
type route struct {
	ctrl controllers.Controllers
}

// Router represent the Router contract
type Router interface {
	Router(port string)
}

/*NewRouter will create an object that represent the Router interface (Router)
 * @parameter
 * c - controllers Interface
 *
 * @represent
 * interface Router
 *
 * @return
 * struct route with value Controllers Interface
 */
func NewRouter(c controllers.Controllers) Router {
	return &route{ctrl: c}
}

func (c *route) Router(port string) {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Get("/", c.ctrl.Tracking)
	})

	logrus.Infof("Server running on port : %s", port)
	logrus.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
