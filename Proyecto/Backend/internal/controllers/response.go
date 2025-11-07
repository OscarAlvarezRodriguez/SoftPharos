package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data,omitempty"`
	Error     *ErrorInfo  `json:"error,omitempty"`
	Timestamp string      `json:"timestamp"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const (
	ErrCodeInternalError  = "INTERNAL_ERROR"
	ErrCodeNotFound       = "NOT_FOUND"
	ErrCodeBadRequest     = "BAD_REQUEST"
	ErrCodeInvalidID      = "INVALID_ID"
	ErrCodeInvalidRequest = "INVALID_REQUEST"
	ErrCodeUnauthorized   = "UNAUTHORIZED"
	ErrCodeForbidden      = "FORBIDDEN"
)

var Response = ResponseBuilder{}

type ResponseBuilder struct{}

func (ResponseBuilder) Success(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, APIResponse{
		Success:   true,
		Data:      data,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func (ResponseBuilder) Error(ctx *gin.Context, statusCode int, errorCode, message string) {
	ctx.JSON(statusCode, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    errorCode,
			Message: message,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func (ResponseBuilder) BadRequest(ctx *gin.Context, message string) {
	ctx.JSON(400, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    ErrCodeBadRequest,
			Message: message,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func (ResponseBuilder) NotFound(ctx *gin.Context, message string) {
	ctx.JSON(404, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    ErrCodeNotFound,
			Message: message,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func (ResponseBuilder) InternalError(ctx *gin.Context, message string) {
	ctx.JSON(500, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    ErrCodeInternalError,
			Message: message,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func (ResponseBuilder) InvalidID(ctx *gin.Context, message string) {
	ctx.JSON(400, APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    ErrCodeInvalidID,
			Message: message,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
