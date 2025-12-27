package middlewares

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/lib/pq"
	"gorm.io/gorm"

	"looker.com/neutral-farming/internal/types"
	"looker.com/neutral-farming/pkg"
)

func HandleErr() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		if len(context.Errors) > 0 {
			state, _ := pkg.ExtractAppState(context)

			fmt.Printf("State: %#v\n", state)

			err := context.Errors.Last()

			slog.Info(fmt.Sprintf("Catching err: %s - request: [%s]", err.Error(), state.TraceID))

			var httpErr *types.HTTPError

			if errors.Is(err, gorm.ErrRecordNotFound) {
				dbErr := types.NewNotFoundError("Record not found.")

				apiError := types.ApiError{
					Message: dbErr.Message,
					UUID:    state.TraceID,
				}

				apiError.ToResponse(context, dbErr.Code)

				return
			}

			if errors.As(err.Err, &httpErr) {
				apiError := types.ApiError{
					Message: httpErr.Message,
					UUID:    state.TraceID,
				}

				apiError.ToResponse(context, httpErr.Code)

				return
			}

			var pgError *pgconn.PgError
			var pqError *pq.Error

			if errors.As(err, &pgError) || errors.As(err, &pqError) {
				slog.Error(fmt.Sprintf("Database error for request %s, %v", state.TraceID, err))

				dbErr := types.NewInternalServerError("Oops, something went wrong. Please try again later.")

				apiError := types.ApiError{
					Message: dbErr.Message,
					UUID:    state.TraceID,
				}

				apiError.ToResponse(context, dbErr.Code)

				return
			}

			slog.Warn("Caution, unhandled error")
			unhandledErr := types.NewInternalServerError("Oops, something went wrong when processing your request. Please try again later.")

			apiError := types.ApiError{
				Message: unhandledErr.Message,
				UUID:    state.TraceID,
			}

			apiError.ToResponse(context, unhandledErr.Code)
		}
	}

}
