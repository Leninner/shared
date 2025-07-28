package exception

import "errors"

func NewDomainException(message string) error {
	return errors.New(message)
}