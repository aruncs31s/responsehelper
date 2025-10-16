# Response Helper
This is a simple utility that helps to standardize api responses across the application.


## Installation

```bash
go get github.com/aruncs31s/responsehelper
``` 

```go
import "github.com/aruncs31s/responsehelper"
``` 

### Usage
I suggest you that you include this in your handler constructor like the following

```go
type Handler struct {
    responseHelper responsehelper.ResponseHelper
}
func NewHandler() *Handler {
    return &Handler{
        responseHelper: responsehelper.NewResponseHelper(),
    }
}
```

Then you can use it in your handler methods like the following
```go
h.responseHelper.Success(c, data)
```


```json
{
    "success": true,
    "data": {
    // response data here
    },
    "meta": "2025-10-01T00:00:00Z"
}
```

### Example used with Gin framework
```go
func (h *userHandler) Login(c *gin.Context) {
	var loginData dto.LoginRequest
	if err := c.ShouldBindJSON(&loginData); err != nil {
		h.responseHelper.BadRequest(c, utils.ErrBadRequest.Error(), utils.ErrDetailBadRequestJSONPayload.Error())
		return
	}

	token, err := h.userService.Login(loginData.Email, loginData.Password)
	if err != nil {
		reaction := utils.NewReaction(err)
		if strings.Contains(err.Error(), utils.ErrNotFound.Error()) {
			h.responseHelper.Unauthorized(c, utils.ErrEmailorPasswordEmpty.Error())
			return
		}
		h.responseHelper.InternalError(c, reaction.Reaction(), err)
		return
	}
	data := map[string]string{"token": token}
	
	h.responseHelper.Success(c, data)
```

## Features

Comes with intellisense support for VSCode and other IDEs.
