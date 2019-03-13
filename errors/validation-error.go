package errors

import (
	"fmt"
)

type ValidationError struct {  
    Err string
    ProblemDiscriptors []string
}

func (e *ValidationError) Error() string {  
    return fmt.Sprintf("{error type: %s. error description: %v", e.Err, e.ProblemDiscriptors)
}



