package usecase

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Firmwave11/otten-coffee/request"
)

var (
	url = "https://gist.githubusercontent.com/nubors/eecf5b8dc838d4e6cc9de9f7b5db236f/raw/d34e1823906d3ab36ccc2e687fcafedf3eacfac9/jne-awb.html"
)

func (r *uc) Tracking(ctx context.Context) (context.Context, int, error) {
	ctx, res, code, err := request.Curl(ctx, url, http.MethodGet, 10*time.Second, nil, nil)

	if err != nil {
		return ctx, 500, err
	}
	log.Println(code, res)

	return ctx, code, nil
}
