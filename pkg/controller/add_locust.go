package controller

import (
	"github.com/amila-ku/locust-operator-opsdk/pkg/controller/locust"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, locust.Add)
}
