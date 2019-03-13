package errors

import (
	"fmt"
)

type MarshalError struct {  
    err    string
}

func (e *MarshalError) Error() string {  
    return fmt.Sprintf("error message: %s", e.err)
}