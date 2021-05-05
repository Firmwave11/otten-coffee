package usecase

import (
	"context"
)

// uc struct with value interface Repository
type uc struct {
}

// Usecases represent the Usecases contract
type Usecases interface {
	Tracking(ctx context.Context) (context.Context, interface{}, int, error)
}

/*NewUC will create an object that represent the Usecases interface (Usecases)
 * @parameter
 * r - Repository Interface
 *
 * @represent
 * interface Usecases
 *
 * @return
 * uc struct with value interface Repository
 */
func NewUC() Usecases {
	return &uc{}
}
