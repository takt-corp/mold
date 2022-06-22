package modifiers

import (
	"context"
	"reflect"

	"github.com/takt-corp/mold"
)

func nilEmpty(ctx context.Context, fl mold.FieldLevel) error {
	if fl.Field().IsZero() {
		fl.Parent().Set(reflect.Zero(fl.Parent().Type()))
	}

	return nil
}
