package responsehelper

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
These are the possible responses from the API.
*/
type ResponseHelper interface {
	// BadRequest sends a 400 Bad Request response
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - message: A brief message describing the error.
	//   - details: Additional details about the error.
	//
	// Example:
	//  responseHelper.BadRequest(c, "Invalid input", "The 'name' field is required.")
	//
	// Example Response Body:
	// {
	//	"success": false,
	//	"error": {
	//		"code":    400,
	//		"status":  "BAD_REQUEST",
	//		"message": "Invalid input",
	//		"details": "The 'name' field is required."
	//	}
	// }
	BadRequest(c *gin.Context, message string, details string)

	// NotFound sends a 404 Not Found response
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - message: A brief message describing the error.
	//
	// Example:
	//  responseHelper.NotFound(c, "Resource not found")
	//
	// Example Response Body:
	// {
	//	"success": false,
	//	"error": {
	//		"code":    404,
	//		"status":  "NOT_FOUND",
	//		"message": "Resource not found"
	//	}
	// }
	NotFound(c *gin.Context, message string)

	// Unauthorized sends a 401 Unauthorized response
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - message: A brief message describing the error.
	//
	// Example:
	//  responseHelper.Unauthorized(c, "Unauthorized access")
	//
	// Example Response Body:
	// {
	//	"success": false,
	//	"error": {
	//		"code":    401,
	//		"status":  "UNAUTHORIZED",
	//		"message": "Unauthorized access"
	//	}
	// }
	Unauthorized(c *gin.Context, message string)

	// InternalError sends a 500 Internal Server Error response
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - message: A brief message describing the error.
	//   - err: The error that occurred.
	//
	// Example:
	//  responseHelper.InternalError(c, "An unexpected error occurred", err)
	//
	// Example Response Body:
	// {
	//	"success": false,
	//	"error": {
	//		"code":    500,
	//		"status":  "INTERNAL_SERVER_ERROR",
	//		"message": "An unexpected error occurred",
	//		"details": "Error details here"
	//	}
	// }
	InternalError(c *gin.Context, message string, err error)

	// Success sends a 200 OK response
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - data: The data to include in the response.
	//
	// Example:
	//  responseHelper.Success(c, data)
	//
	// Example Response Body:
	// {
	//	"success": true,
	//	"data": {
	//		// response data here
	//	},
	//	"meta": "2023-01-01T00:00:00Z"
	// }
	Success(c *gin.Context, data interface{})

	// SuccessWithPagination sends a 200 OK response with pagination metadata
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - data: The data to include in the response.
	//   - meta: The pagination metadata.
	//
	// Example:
	//  responseHelper.SuccessWithPagination(c, data, meta)
	//
	// Example Response Body:
	// {
	//	"success": true,
	//	"data": {
	//		// response data here
	//	},
	//	"pagination": {
	//		"currentPage": 3,
	//		"pageSize": 10,
	//		"totalPages": 3,
	//		"totalRecords": 27
	//	}
	// }
	SuccessWithPagination(c *gin.Context, data interface{}, meta interface{})

	// Created sends a 201 Created response
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - data: The data to include in the response.
	//
	// Example:
	//  responseHelper.Created(c, data)
	//
	// Example Response Body:
	// {
	//	"success": true,
	//	"data": {
	//		// response data here
	//	},
	//	"meta": "2023-01-01T00:00:00Z"
	// }
	Created(c *gin.Context, data interface{})

	// Deleted sends a 204 No Content response
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - message: what you have deleted eg: qualification
	//
	// Example:
	//  responseHelper.Deleted(c, "qualification")
	//
	// Example Response Body:
	// {
	//	"success": true,
	//	"message": "qualification deleted successfully",
	//	"meta": "2023-01-01T00:00:00Z"
	// }
	Deleted(c *gin.Context, message string)
}

// Response helper - centralizes response logic
// The context is same in the case of all the responses , but there is no need to , group it in a struct
// only one response per request , so there is no reuse for context.
type responseHelper struct{}

func NewResponseHelper() ResponseHelper {
	return &responseHelper{}
}

func (r *responseHelper) BadRequest(c *gin.Context, message string, details string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error": gin.H{
			"code":    400,
			"status":  "BAD_REQUEST",
			"message": message,
			"details": details,
		},
	})
}

func (r *responseHelper) NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"error": gin.H{
			"code":    404,
			"status":  "NOT_FOUND",
			"message": message,
		},
	})
}

func (r *responseHelper) Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"error": gin.H{
			"code":    401,
			"status":  "UNAUTHORIZED",
			"message": message,
		},
	})
}

func (r *responseHelper) InternalError(c *gin.Context, message string, err error) {
	log.Printf("Internal error: %v", err)
	// Check if sanitization of error is needed,
	/*
		1. There is a possibility of leaking information through error messages.
	*/
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"error": gin.H{
			"code":    500,
			"status":  "INTERNAL_SERVER_ERROR",
			"message": message,
			"details": err.Error(), // sanitizing this in production
		},
		"data": nil,
	})
}

func (r *responseHelper) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
		"meta":    time.Now().Format(time.RFC3339),
	})
}

func (r *responseHelper) SuccessWithPagination(c *gin.Context, data interface{}, meta interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"data":       data,
		"pagination": meta,
	})
}

func (r *responseHelper) Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    data,
		"meta":    time.Now().Format(time.RFC3339),
	})
}

func (r *responseHelper) Deleted(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message + " deleted successfully",
		"meta":    time.Now().Format(time.RFC3339),
	})
}
