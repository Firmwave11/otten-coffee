package controllers

import (
	"fmt"
	"net/http"
)

func (u *ctrl) Tracking(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, code, err := u.uc.Tracking(ctx)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, string("succes"))
}
