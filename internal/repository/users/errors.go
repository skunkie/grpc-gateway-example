package users

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func errIdNotFound(id string) error {
	return status.Errorf(codes.NotFound, http.StatusText(http.StatusNotFound)+": %s", id)
}
