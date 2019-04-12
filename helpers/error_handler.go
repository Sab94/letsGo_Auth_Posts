package helpers

import (
	"github.com/posts-api/types"
	"log"
	"runtime"
)

func ErrorMessage(err error, error interface{}) (int, *types.APIErrors) {
	var apiErrors *types.APIErrors

	// This the best way to log?
	trace := make([]byte, 1024)
	runtime.Stack(trace, true)
	log.Printf("ERROR: %s\n%s", err, trace)

		switch error.(type) {
		case *types.APIError:
			apiError := error.(*types.APIError)
			apiErrors = &types.APIErrors{
				Errors: []*types.APIError{apiError},
			}
		case *types.APIErrors:
			apiErrors = error.(*types.APIErrors)
		default:
			apiErrors = &types.APIErrors{
				Errors: []*types.APIError{types.ErrUnknown},
			}
		}
	return apiErrors.Status(), apiErrors
}
