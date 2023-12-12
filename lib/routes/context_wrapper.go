package routes

import (
	"backend_scrolling_simulator/lib"
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/services/user"
	"github.com/gin-gonic/gin"
)

type contextWrapper[T any] struct {
	ctx  *gin.Context
	user *models.User
	body T
}

type GinContextWrapper[T any] interface {
	ReturnErrorResponse(err error, statusCode int)
	GetCurrentUser() models.User
	GetBody() T
}

func getTokenFromHeader(ctx *gin.Context) string {
	header := ctx.GetHeader("Authorization")
	return header[7:]
}

func (ctx *contextWrapper[T]) GetCurrentUser() models.User {
	return *ctx.user
}

func (ctx *contextWrapper[T]) ReturnErrorResponse(err error, statusCode int) {
	if err != nil {
		ctx.ctx.JSON(statusCode, models.NewErrorApiResponse(lib.Error{Msg: err.Error(), Reason: ""}))
	}
}

func (ctx *contextWrapper[T]) GetBody() T {
	return ctx.body
}

// GetContextWrapper
// This function is used to get the current user and the body of the request.
// * It returns a contextWrapper that implements the GinContextWrapper interface.
// * Any errors that occur during the process are handled by the contextWrapper, returning a JSON response with the error.
func GetContextWrapper[T any](ctx *gin.Context) (GinContextWrapper[T], error) {
	var body T
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(400, models.NewErrorApiResponse(lib.Error{Msg: err.Error(), Reason: ""}))
		return nil, err
	}

	ctxWrapper := &contextWrapper[T]{ctx: ctx, body: body}

	currentUser, err := user.GetCurrentUser(getTokenFromHeader(ctx))
	if err != nil {
		// TODO: Implement forbiden
		ctxWrapper.ReturnErrorResponse(err, 401)
		return nil, err
	}

	ctxWrapper.user = &currentUser
	return ctxWrapper, nil
}
