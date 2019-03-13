package errors

import (
	"fmt"
)

type HttpError struct {  
    Err    string
    StatusCode int
}

func (e *HttpError) Error() string {  
    return fmt.Sprintf("http status: %d error message: %s", e.StatusCode, e.Err)
}