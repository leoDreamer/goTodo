package errors

import (
	pkgerr "github.com/pkg/errors"
)

func ErrFn(msg string, err error) {
	panic(pkgerr.Wrap(err, msg))
}

// DBERR 数据库错误
const DBERR string = "database error"

// PARAMERR 参数错误
const PARAMERR string = "params error"

// FORBIDDEN 参数错误
const FORBIDDEN string = "forbidden"
