package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type ErrorResponse struct {
	Msg       string `json:"msg"`
	Error     string `json:"error,omitempty"`
	ErrorCode string `json:"errorCode"`
}

func sendResponse(statusCode int, c *gin.Context, response *Response) {
	c.AbortWithStatusJSON(statusCode, *response)
}

func sendErrorResponse(statusCode int, c *gin.Context, response *ErrorResponse) {
	c.AbortWithStatusJSON(statusCode, *response)
}

func OkResponse(c *gin.Context, response *Response) {
	if response == nil {
		response = &Response{}
	}
	if response.Msg == "" {
		response.Msg = "Request completed successfully"
	}
	sendResponse(http.StatusOK, c, response)
}

func CreatedSuccessResponse(c *gin.Context, response *Response) {
	if response == nil {
		response = &Response{}
	}
	if response.Msg == "" {
		response.Msg = "Resource successfully created."
	}
	sendResponse(http.StatusCreated, c, response)
}

func BadRequestResponse(c *gin.Context, response *ErrorResponse) {
	if response == nil {
		response = &ErrorResponse{}
	}
	if response.Msg == "" {
		response.Msg = "Something went wrong with your request. Please check and try again."
	}
	sendErrorResponse(http.StatusBadRequest, c, response)
}

func UnauthorizedResponse(c *gin.Context, response *ErrorResponse) {
	if response == nil {
		response = &ErrorResponse{}
	}
	if response.Msg == "" {
		response.Msg = "It seems you're not authorized to access this. Please log in or check your credentials."
	}
	sendErrorResponse(http.StatusBadRequest, c, response)
}

func PaymentRequiredErrorResponse(c *gin.Context, response *ErrorResponse) {
	if response == nil {
		response = &ErrorResponse{}
	}
	if response.Msg == "" {
		response.Msg = "Payment is required to access this resource."
	}
	sendErrorResponse(http.StatusPaymentRequired, c, response)
}

func ForbiddenErrorResponse(c *gin.Context, response *ErrorResponse) {
	if response == nil {
		response = &ErrorResponse{}
	}
	if response.Msg == "" {
		response.Msg = "You do not have permission to access this resource."
	}
	sendErrorResponse(http.StatusForbidden, c, response)
}

func NotFoundErrorResponse(c *gin.Context, response *ErrorResponse) {
	if response == nil {
		response = &ErrorResponse{}
	}
	if response.Msg == "" {
		response.Msg = "The requested resource could not be found on the server."
	}
	sendErrorResponse(http.StatusNotFound, c, response)
}

func UnprocessableEntityErrorResponse(c *gin.Context, response *ErrorResponse) {
	if response == nil {
		response = &ErrorResponse{}
	}
	if response.Msg == "" {
		response.Msg = "There seems to be an issue with your request. Please review your input and try again."
	}
	sendErrorResponse(http.StatusUnprocessableEntity, c, response)
}

func ServerErrorErrorResponse(c *gin.Context, response *ErrorResponse) {
	if response == nil {
		response = &ErrorResponse{}
	}
	if response.Msg == "" {
		response.Msg = "An unexpected error occurred. Please contact support."
	}
	sendErrorResponse(http.StatusInternalServerError, c, response)
}

func TooManyRequestsErrorResponse(c *gin.Context, response *ErrorResponse) {
	if response == nil {
		response = &ErrorResponse{}
	}
	if response.Msg == "" {
		response.Msg = "Too many requests."
	}
	sendErrorResponse(http.StatusTooManyRequests, c, response)
}

func GeneralResponse(c *gin.Context, statusCode int, response interface{}) {
	if response != nil {
		c.JSON(statusCode, response)
	} else {
		c.AbortWithStatus(statusCode)
	}
}
