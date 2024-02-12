package schema

import "errors"

var ErrUsernameAlreadyInUse = errors.New("the requested Username was already present in the database")
