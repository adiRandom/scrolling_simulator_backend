package routes

import (
	"backend_scrolling_simulator/lib"
	"backend_scrolling_simulator/models"
	"backend_scrolling_simulator/services/user"
	"fmt"
	"github.com/gin-gonic/gin"
)

type contextWrapper[Q, B, P any] struct {
	ctx         *gin.Context
	user        *models.User
	body        B
	queryParams Q
	pathParams  P
}

func (ctx *contextWrapper[Q, B, P]) GetQueryParams() Q {
	return ctx.queryParams
}

type GinContextWrapper[Q, B, P any] interface {
	ReturnErrorResponse(err error, statusCode int)
	GetCurrentUser() models.User
	GetBody() B
	GetQueryParams() Q
	ReturnJSON(statusCode int, data interface{})
	GetPathParams() P
}

func getTokenFromHeader(ctx *gin.Context) string {
	// TODO: Get token
	return ""
	header := ctx.GetHeader("Authorization")
	return header[7:]
}

func (ctx *contextWrapper[Q, B, P]) GetCurrentUser() models.User {
	return *ctx.user
}

func (ctx *contextWrapper[Q, B, P]) ReturnErrorResponse(err error, statusCode int) {
	if err != nil {
		ctx.ctx.JSON(statusCode, models.NewErrorApiResponse(lib.Error{Msg: err.Error(), Reason: ""}))
	}
}

func (ctx *contextWrapper[Q, B, P]) GetBody() B {
	return ctx.body
}

// GetContextWrapper
// This function is used to get the current user and the body of the request.
// * It returns a contextWrapper that implements the GinContextWrapper interface.
// * Any errors that occur during the process are handled by the contextWrapper, returning a JSON response with the error.
func GetContextWrapper[Q, B, P any](ctx *gin.Context) (GinContextWrapper[Q, B, P], error) {
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

	pathParams, err := extractPathParams[P](ctx)
	if err != nil {
		fmt.Print(err.Error())
		ctx.JSON(400, models.NewErrorApiResponse(lib.Error{Msg: err.Error(), Reason: ""}))
		return nil, err
	}

	ctxWrapper := &contextWrapper[Q, B, P]{ctx: ctx, body: body, queryParams: *queryParams, pathParams: *pathParams}

	currentUser, err := user.GetCurrentUser(getTokenFromHeader(ctx))
	if err != nil {
		// TODO: Implement forbiden
		ctxWrapper.ReturnErrorResponse(err, 401)
		return nil, err
	}

	ctxWrapper.user = &currentUser
	return ctxWrapper, nil
}

func GetCtxWithBody[B any](ctx *gin.Context) (GinContextWrapper[lib.None, B, lib.None], error) {
	return GetContextWrapper[lib.None, B, lib.None](ctx)
}

func GetCtxWithQuery[Q any](ctx *gin.Context) (GinContextWrapper[Q, lib.None, lib.None], error) {
	return GetContextWrapper[Q, lib.None, lib.None](ctx)
}

func GetCtxWithBodyAndQuery[B, Q any](ctx *gin.Context) (GinContextWrapper[Q, B, lib.None], error) {
	return GetContextWrapper[Q, B, lib.None](ctx)
}

func GetCtxWithPath[P any](ctx *gin.Context) (GinContextWrapper[lib.None, lib.None, P], error) {
	return GetContextWrapper[lib.None, lib.None, P](ctx)
}

func GetCtxWithBodyAndPath[B, P any](ctx *gin.Context) (GinContextWrapper[lib.None, B, P], error) {
	return GetContextWrapper[lib.None, B, P](ctx)
}

func GetCtxWithQueryAndPathParams[Q, P any](ctx *gin.Context) (GinContextWrapper[Q, lib.None, P], error) {
	return GetContextWrapper[Q, lib.None, P](ctx)
}

func GetEmptyCtx(ctx *gin.Context) (GinContextWrapper[lib.None, lib.None, lib.None], error) {
	return GetContextWrapper[lib.None, lib.None, lib.None](ctx)
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

func (ctx *contextWrapper[Q, B, P]) GetPathParams() P {
	return ctx.pathParams
}

func extractPathParams[P any](ctx *gin.Context) (*P, error) {
	var pathParams P

	if !lib.IsNone(pathParams) {
		err := ctx.BindUri(&pathParams)
		if err != nil {
			return nil, err
		}
	}

	return &pathParams, nil
}

func (ctx *contextWrapper[Q, B, P]) ReturnJSON(statusCode int, data interface{}) {
	ctx.ctx.JSON(statusCode, models.NewSuccessApiResponse(data))
}
