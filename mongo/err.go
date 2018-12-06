package mongo



import (
	"strings"
)

const (
	ErrNoDefaultConnection    = "No default connection"
	ErrExistConnectionAlias   = "Exist connection alias"
	ErrNoDefaultDatabase      = "No default database"
	ErrNoConnection           = "No connection"
	ErrCannotSwitchCollection = "Can not switch collection"
	ErrMongoObjDestroyed      = "The mongo object has been destoryed"
	ErrCollectionDuplicateKey = "duplicate key error"
)

func EqualError(err error, str string) bool {
	return str == err.Error() || strings.HasPrefix(err.Error(), str) || strings.Index(err.Error(), str) > -1
}
