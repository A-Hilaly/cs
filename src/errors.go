package gocs

import (
    "errors"
)


func Error(message string) error {
    // Return New error type
    return errors.New(message)
}
