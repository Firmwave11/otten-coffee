package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (u *ctrl) Tracking(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, res, code, err := u.uc.Tracking(ctx)

	datares, err := json.Marshal(res)

	if err != nil || code >= 400 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		fmt.Fprintf(w, string(datares))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, string(datares))
}
