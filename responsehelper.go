package responsehelper

/*
Author: Arun CS
Date: 2025-10-16
Last Modified: 2026-08-06

*/

import (
	"fmt"
	"net/http"

	"github.com/aruncs31s/gologger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

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

	// AlreadyExists sends a 409 Conflict response indicating resource already exists
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - resource: The name of the resource that already exists.
	//   - err: The error that occurred.
	//
	// Example:
	//  responseHelper.AlreadyExists(c, "User", err)
	//
	// Example Response Body:
	// {
	//	"success": false,
	//	"error": {
	//		"code":    409,
	//		"status":  "CONFLICT",
	//		"message": "User already exists",
	//		"details": "Error details here"
	//	}
	// }
	AlreadyExists(c *gin.Context, resource string, err error)

	// Conflict sends a 409 Conflict response
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - message: A brief message describing the error.
	//   - err: The error that occurred.
	//
	// Example:
	//  h.responseHelper.Conflict(c, "Resource conflict", err)
	//
	// Example Response Body:
	// {
	//	"success": false,
	//	"error": {
	//		"code":    409,
	//		"status":  "CONFLICT",
	//		"message": "Resource conflict",
	//		"details": "Error details here"
	//	}
	// }
	Conflict(c *gin.Context, message string, err error)
	// NotFound sends a 404 Not Found response
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - message: A brief message describing the error.
	//
	// Example:
	//  h.responseHelper.NotFound(c, "Resource not found")
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
	// h.responseHelper.Unauthorized(c, "Unauthorized access")
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
	// Forbidden sends a 403 Forbidden response
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - message: A brief message describing the error.
	//
	// Example:
	// h.responseHelper.Forbidden(c, "Forbidden access")
	//
	// Example Response Body:
	// {
	//	"success": false,
	//	"error": {
	//		"code":    403,
	//		"status":  "FORBIDDEN",
	//		"message": "This User does not have access to the resource"
	//	}
	// }
	Forbidden(c *gin.Context, message string)
	// InternalError sends a 500 Internal Server Error response
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - message: A brief message describing the error.
	//   - err: The error that occurred.
	//
	// Example:
	//  h.responseHelper.InternalError(c, "An unexpected error occurred", err)
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
	//  h.responseHelper.Success(c, data)
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
	List(
		c *gin.Context,
		data interface{},
		totalCount ...int,
	)

	// ListWithMessage sends a 200 OK response with a list of resources, a count, and a message
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - data: The list data to include in the response.
	//   - count: The total count of resources.
	//   - message: A brief message to include in the response.
	//
	// Example:
	//  responseHelper.ListWithMessage(c, users, 42, "Users retrieved successfully")
	//
	// Example Response Body:
	// {
	//	"success": true,
	//	"list": [
	//		// response data here
	//	],
	//	"meta": "2023-01-01T00:00:00Z",
	//	"total_count": 42,
	//	"message": "Users retrieved successfully"
	// }
	ListWithMessage(
		c *gin.Context,
		data interface{},
		count int,
		message string,
	)
	// SuccessWithPagination sends a 200 OK response with pagination metadata
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - data: The data to include in the response.
	//   - meta: The pagination metadata.
	//
	// Example:
	//  h.responseHelper.SuccessWithPagination(c, data, meta)
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

	// Extended version of the [Success] method that includes a custom message in the response
	SuccessWithMessage(c *gin.Context, data interface{}, message string)
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

	// NoContent sends a 204 No Content response
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//
	// Example:
	//  responseHelper.NoContent(c)
	//
	// Example Response Body:
	// {
	//	"success": true,
	//	"data":    null,
	//	"meta":    "2023-01-01T00:00:00Z"
	// }
	NoContent(c *gin.Context)

	// SendDoc sends a 200 OK response with a file attachment
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - contentType: The MIME type of the document (e.g. "application/pdf", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet").
	//   - filename: The name of the file to be sent as an attachment.
	//   - documentBytes: The raw bytes of the document to send.
	//
	// Example:
	//  responseHelper.SendDoc(c, "application/pdf", "report.pdf", pdfBytes)
	//
	// Example Response Headers:
	// Content-Type: application/pdf
	// Content-Disposition: attachment; filename=report.pdf
	// Content-Length: <byte length>
	SendDoc(
		c *gin.Context,
		contentType string,
		filename string,
		documentBytes []byte,
	)

	// Ok sends a 200 OK response with a request ID and data
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - data: The data to include in the response.
	//
	// Example:
	//  responseHelper.Ok(c, data)
	//
	// Example Response Body:
	// {
	//	"req_id":  "abc123",
	//	"success": true,
	//	"data": {
	//		// response data here
	//	},
	//	"meta": "2023-01-01T00:00:00Z"
	// }
	Ok(c *gin.Context, data interface{})

	// Filters sends a 200 OK response with available filter and sort options
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - filters: The available filter options to include in the response.
	//   - sort: Optional variadic sort options to include in the response.
	//
	// Example:
	//  responseHelper.Filters(c, availableFilters, sortOptions)
	//
	// Example Response Body:
	// {
	//	"success": true,
	//	"available_filters": {
	//		// filter options here
	//	},
	//	"available_sort": [
	//		// sort options here
	//	],
	//	"meta": "2023-01-01T00:00:00Z"
	// }
	Filters(
		c *gin.Context,
		filters any,
		sort ...any,
	)

	// FilterDropdown sends a 200 OK response with structured dropdown filter and sort options
	//
	// Parameters:
	//   - c: The Gin context to send the response to.
	//   - filters: A slice of FilterDropdown items, each with an ID, Name, and optional ParentID.
	//   - sorts: Optional variadic sort options to include in the response.
	//
	// Example:
	//  responseHelper.FilterDropdown(c, []responseHelper.FilterDropdown{{ID: 1, Name: "Active"}}, sortOptions)
	//
	// Example Response Body:
	// {
	//	"success": true,
	//	"available_filters": [
	//		{"id": 1, "name": "Active"},
	//		{"id": 2, "name": "Inactive", "parent_id": 1}
	//	],
	//	"available_sorts": [
	//		// sort options here
	//	],
	//	"meta": "2023-01-01T00:00:00Z"
	// }
	FilterDropdown(c *gin.Context, filters []FilterDropdown, sorts ...any)
	// TODO: Document
	SuccessList(c *gin.Context, data interface{})
}

// Response helper - centralizes response logic
// The context is same in the case of all the responses , but there is no need to , group it in a struct
// only one response per request , so there is no reuse for context.
type responseHelper struct{}

func NewResponseHelper() ResponseHelper {
	return &responseHelper{}
}

func (r *responseHelper) BadRequest(c *gin.Context, message string, details string) {
	fields, reqID, meta := getReqIDAndMeta(c)
	fields = append(fields, zap.String("message", message))
	fields = append(fields, zap.String("details", details))
	gologger.Warn("BAD_REQUEST", fields...)
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error": gin.H{
			"code":    http.StatusBadRequest,
			"status":  "BAD_REQUEST",
			"message": message,
			"details": details,
		},
		"meta":   meta,
		"req_id": reqID,
	})
}

func (r *responseHelper) AlreadyExists(c *gin.Context, resource string, err error) {
	r.Conflict(c, resource+" already exists", err)
}

func (r *responseHelper) Conflict(c *gin.Context, message string, err error) {
	fields, reqID, meta := getReqIDAndMeta(c)
	fields = append(fields, zap.String("message", message))
	fields = append(fields, zap.Error(err))
	gologger.Warn("CONFLICT", fields...)

	c.JSON(http.StatusConflict, gin.H{
		"success": false,
		"error": gin.H{
			"code":    http.StatusConflict,
			"status":  "CONFLICT",
			"message": message,
			"details": err.Error(),
		},
		"meta":   meta,
		"req_id": reqID,
	})
}

func (r *responseHelper) NotFound(c *gin.Context, message string) {
	fields, reqID, meta := getReqIDAndMeta(c)
	fields = append(fields, zap.String("message", message))
	gologger.Warn("NOT_FOUND", fields...)

	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"error": gin.H{
			"code":    http.StatusNotFound,
			"status":  "NOT_FOUND",
			"message": message,
		},
		"meta":   meta,
		"req_id": reqID,
	})
}

func (r *responseHelper) Unauthorized(c *gin.Context, message string) {
	fields, reqID, meta := getReqIDAndMeta(c)
	fields = append(fields, zap.String("message", message))
	gologger.Warn("UNAUTHORIZED", fields...)
	c.JSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"error": gin.H{
			"code":    http.StatusUnauthorized,
			"status":  "UNAUTHORIZED",
			"message": message,
		},
		"meta":   meta,
		"req_id": reqID,
	})
}

func (r *responseHelper) InternalError(c *gin.Context, message string, err error) {
	// Check if sanitization of error is needed,
	/*
		1. There is a possibility of leaking information through error messages.
	*/

	fields, reqID, meta := getReqIDAndMeta(c)
	fields = append(fields, zap.String("message", message))
	fields = append(fields, zap.Error(err))
	gologger.Error("INTERNAL_SERVER_ERROR", fields...)

	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"error": gin.H{
			"code":    http.StatusInternalServerError,
			"status":  "INTERNAL_SERVER_ERROR",
			"message": message,
			"details": err.Error(), // sanitizing this in production
		},
		"data":   nil,
		"meta":   meta,
		"req_id": reqID,
	})
}

func (r *responseHelper) Success(c *gin.Context, data interface{}) {
	fields, reqID, meta := getReqIDAndMeta(c)
	gologger.Info("SUCCESS", fields...)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
		"meta":    meta,
		"req_id":  reqID,
	})
}

func (r *responseHelper) SuccessWithPagination(c *gin.Context, data interface{}, paginationMeta interface{}) {
	fields, reqID, meta := getReqIDAndMeta(c)
	gologger.Info("SUCCESS", fields...)

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"data":       data,
		"pagination": paginationMeta,
		"meta":       meta,
		"req_id":     reqID,
	})
}

func (r *responseHelper) Created(c *gin.Context, data interface{}) {
	fields, reqID, meta := getReqIDAndMeta(c)
	gologger.Info("SUCCESS", fields...)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    data,
		"meta":    meta,
		"req_id":  reqID,
	})
}

func (r *responseHelper) Deleted(c *gin.Context, message string) {
	fields, reqID, meta := getReqIDAndMeta(c)
	fields = append(fields, zap.String("message", message))
	gologger.Info("SUCCESS", fields...)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message + " deleted successfully",
		"meta":    meta,
		"req_id":  reqID,
	})
}

func (r *responseHelper) Forbidden(c *gin.Context, message string) {
	fields, reqID, meta := getReqIDAndMeta(c)
	fields = append(fields, zap.String("message", message))
	gologger.Warn("FORBIDDEN", fields...)
	c.JSON(http.StatusForbidden, gin.H{
		"success": false,
		"error": gin.H{
			"code":    http.StatusForbidden,
			"status":  "FORBIDDEN",
			"message": message,
		},
		"meta":   meta,
		"req_id": reqID,
	})
}

func (r *responseHelper) NoContent(c *gin.Context) {
	fields, _, meta := getReqIDAndMeta(c)
	gologger.Info("SUCCESS", fields...)
	c.JSON(http.StatusNoContent, gin.H{
		"success": true,
		"data":    nil,
		"meta":    meta,
	})
}

type ListResponse struct {
	Success    bool        `json:"success"`
	List       interface{} `json:"list"`
	Meta       interface{} `json:"meta,omitempty"`
	TotalCount int         `json:"total_count,omitempty"`
	Message    string      `json:"message,omitempty"`
	ReqUI      string      `json:"req_id,omitempty"`
}

func (r *responseHelper) List(
	c *gin.Context,
	data interface{},
	totalCount ...int,
) {
	fields, reqID, meta := getReqIDAndMeta(c)
	gologger.Info("SUCCESS", fields...)

	resp := ListResponse{
		Success: true,
		List:    data,
		Meta:    meta,
		// Forgive me :)
		ReqUI: reqID,
	}

	if len(totalCount) > 0 {
		resp.TotalCount = totalCount[0]
	}

	c.JSON(http.StatusOK, resp)
}

func (r *responseHelper) ListWithMessage(
	c *gin.Context,
	data interface{},
	count int,
	message string,
) {
	fields, reqID, meta := getReqIDAndMeta(c)
	fields = append(fields, zap.String("message", message))
	fields = append(fields, zap.Int("total_count", count))
	gologger.Info("SUCCESS", fields...)

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"list":        data,
		"meta":        meta,
		"total_count": count,
		"message":     message,
		"req_id":      reqID,
	})
}

func (r *responseHelper) SendDoc(
	c *gin.Context,
	contentType string,
	filename string,
	documentBytes []byte,
) {
	c.Header("Content-Type", contentType)
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Length", fmt.Sprintf("%d", len(documentBytes)))

	// Send file
	c.Data(http.StatusOK, contentType, documentBytes)
}

func (r *responseHelper) Ok(c *gin.Context, data interface{}) {
	fields, reqID, meta := getReqIDAndMeta(c)
	gologger.Info("SUCCESS", fields...)

	c.JSON(http.StatusOK, gin.H{
		"req_id":  reqID,
		"success": true,
		"data":    data,
		"meta":    meta,
	})
}

type FilterDropdown struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID int    `json:"parent_id,omitempty"`
}

func (r *responseHelper) Filters(
	c *gin.Context,
	filters any,
	sort ...any,
) {
	fields, _, meta := getReqIDAndMeta(c)
	gologger.Info("SUCCESS", fields...)

	c.JSON(http.StatusOK, gin.H{
		"success":           true,
		"available_filters": filters,
		"available_sorts":   sort,
		"meta":              meta,
	})
}

func (r *responseHelper) FilterDropdown(c *gin.Context, filters []FilterDropdown, sorts ...any) {
	fields, _, meta := getReqIDAndMeta(c)
	gologger.Info("SUCCESS", fields...)

	c.JSON(http.StatusOK, gin.H{
		"success":           true,
		"available_filters": filters,
		"available_sorts":   sorts,
		"meta":              meta,
	})
}

func (r *responseHelper) SuccessList(c *gin.Context, data interface{}) {
	fields, _, _ := getReqIDAndMeta(c)
	gologger.Info("SUCCESS", fields...)

	r.List(
		c,
		data,
	)
}

func (r *responseHelper) SuccessWithMessage(c *gin.Context, data interface{}, message string) {
	fields, reqID, meta := getReqIDAndMeta(c)

	fields = append(fields, zap.String("message", message))
	gologger.Info("SUCCESS", fields...)
	c.JSON(http.StatusOK, gin.H{
		"req_id":  reqID,
		"success": true,
		"data":    data,
		"message": message,
		"meta":    meta,
	})
}

func getReqIDAndMeta(c *gin.Context) ([]zap.Field, string, any) {
	reqID := c.GetString("req_id")
	// Try Fallback
	if reqID == "" {
		reqID = c.GetString("request_id")
	}
	fields := []zap.Field{}
	if reqID != "" {
		fields = append(fields, zap.String("req_id", reqID))
	}

	meta, exists := c.Get("meta")
	if exists {
		fields = append(fields, zap.Any("meta", meta))
	}
	return fields, reqID, meta
}
