package resolver

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	RegisterTokenInternalServerError = 1000
)

type errMessage map[int]string

var message errMessage = map[int]string{
	RegisterTokenInternalServerError: "internal_server_error",
}

func (e errMessage) build(errCode int, err error) error {
	ext := map[string]interface{}{
		"code": errCode,
	}

	if err != nil {
		ext = map[string]interface{}{
			"err": err.Error(),
		}
	}

	return &gqlerror.Error{
		Message:    e[errCode],
		Extensions: ext,
	}
}

func (e errMessage) add(ctx context.Context, errCode int, err error) {
	graphql.AddError(ctx, e.build(errCode, err))
}
