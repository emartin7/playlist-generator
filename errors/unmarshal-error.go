package errors

import (
	"fmt"
)

type UnmarshalError struct {  
    Err    string
}

func (e *UnmarshalError) Error() string {  
    return fmt.Sprintf("error message: %s", e.Err)
}