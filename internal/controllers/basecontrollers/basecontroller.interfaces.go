package basecontrollers

import (
	"net/http"

	"github.com/deanrtaylor1/backend-go/internal/config"
	db "github.com/deanrtaylor1/backend-go/internal/db/sqlc"
	"github.com/go-playground/validator/v10"
)

type BaseController struct {
	Store     db.Store
	Config    config.EnvConfig
	Validator *validator.Validate
}

type Controller interface {
	Routes() http.Handler
}
