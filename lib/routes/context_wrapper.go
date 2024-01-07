package routes

import (
	"backend_scrolling_simulator/lib"
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/services/user"
	"fmt"
	"github.com/gin-gonic/gin"
)

type contextWrapper[Q, B any] struct {
	ctx         *gin.Context
	user        *models.User
	body        B
	queryParams Q
}

func (ctx *contextWrapper[Q, B]) GetQueryParams() Q {
	return ctx.queryParams
}

type GinContextWrapper[Q, B any] interface {
	ReturnErrorResponse(err error, statusCode int)
	GetCurrentUser() models.User
	GetBody() B
	GetQueryParams() Q
	ReturnJSON(statusCode int, data interface{})
}

func getTokenFromHeader(ctx *gin.Context) string {
	header := ctx.GetHeader("Authorization")
	return header[7:]
}

func (ctx *contextWrapper[Q, B]) GetCurrentUser() models.User {
	return *ctx.user
}

func (ctx *contextWrapper[Q, B]) ReturnErrorResponse(err error, statusCode int) {
	if err != nil {
		ctx.ctx.JSON(statusCode, models.NewErrorApiResponse(lib.Error{Msg: err.Error(), Reason: ""}))
	}
}

func (ctx *contextWrapper[Q, B]) GetBody() B {
	return ctx.body
}

// GetContextWrapper
// This function is used to get the current user and the body of the request.
// * It returns a contextWrapper that implements the GinContextWrapper interface.
// * Any errors that occur during the process are handled by the contextWrapper, returning a JSON response with the error.
func GetContextWrapper[Q, B any](ctx *gin.Context) (GinContextWrapper[Q, B], error) {
	var body B
	if !lib.IsNone(body) {
		err := ctx.BindJSON(&body)
		if err != nil {
			fmt.Print(err.Error())
			ctx.JSON(400, models.NewErrorApiResponse(lib.Error{Msg: err.Error(), Reason: ""}))
			return nil, err
		}
	}

	queryParams, err := getQueryParams[Q](ctx)
	if err != nil {
		fmt.Print(err.Error())
		ctx.JSON(400, models.NewErrorApiResponse(lib.Error{Msg: err.Error(), Reason: ""}))
		return nil, err
	}

	ctxWrapper := &contextWrapper[Q, B]{ctx: ctx, body: body, queryParams: *queryParams}

	currentUser, err := user.GetCurrentUser(getTokenFromHeader(ctx))
	if err != nil {
		// TODO: Implement forbiden
		ctxWrapper.ReturnErrorResponse(err, 401)
		return nil, err
	}

	ctxWrapper.user = &currentUser
	return ctxWrapper, nil
}

func getQueryParams[Q any](ctx *gin.Context) (*Q, error) {
	var queryParams Q

	if !lib.IsNone(queryParams) {
		err := ctx.BindQuery(&queryParams)
		if err != nil {
			return nil, err
		}
	}

	return &queryParams, nil
}

func (ctx *contextWrapper[Q, B]) ReturnJSON(statusCode int, data interface{}) {
	ctx.ctx.JSON(statusCode, models.NewSuccessApiResponse(data))
}
