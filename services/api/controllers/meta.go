package controllers

import (
	"fmt"
	"net/http"
)

type MetaController struct{}

func (mc *MetaController) Healtcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Service Healthy")
}
