package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (u *ctrl) Tracking(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, res, code, err := u.uc.Tracking(ctx)

	if err != nil || code >= 400 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		fmt.Fprintf(w, err.Error())
		return
	}

	datares, err := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, string(datares))
}
