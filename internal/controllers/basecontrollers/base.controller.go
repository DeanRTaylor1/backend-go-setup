package basecontrollers

import (
	"github.com/deanrtaylor1/backend-go/internal/config"
	db "github.com/deanrtaylor1/backend-go/internal/db/sqlc"
	"github.com/go-playground/validator/v10"
)

func NewBaseController(store db.Store, config config.EnvConfig) *BaseController {
	return &BaseController{Store: store, Config: config, Validator: validator.New()}
}
