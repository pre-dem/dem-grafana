package tsdb

import (
	"fmt"
	"qiniu.com/pandora/base/reqerr"
)

const errCodePrefixLen = 5

type errBuilder struct{}

func (e errBuilder) Build(msg, text, reqId string, code int) error {

	err := reqerr.New(msg, text, reqId, code)
	if len(msg) <= errCodePrefixLen {
		return err
	}
	errId := msg[:errCodePrefixLen]

	switch errId {
	case "E7100":
		err.ErrorType = reqerr.NoSuchRepoError
	case "E6102":
		err.ErrorType = reqerr.RepoAlreadyExistsError
	case "E6205":
		err.ErrorType = reqerr.NoSuchRetentionError
	case "E6300":
		err.ErrorType = reqerr.InvalidSeriesNameError
	case "E6302":
		err.ErrorType = reqerr.SeriesAlreadyExistsError
	case "E6303":
		err.ErrorType = reqerr.NoSuchSeriesError
	case "E6400":
		err.ErrorType = reqerr.InvalidViewNameError
	case "E6403":
		err.ErrorType = reqerr.InvalidViewSqlError
	case "E6404", "E6405":
		err.ErrorType = reqerr.ViewFuncNotSupportError
	case "E6410":
		err.ErrorType = reqerr.NoSuchViewError
	case "E6411":
		err.ErrorType = reqerr.ViewAlreadyExistsError
	case "E6412":
		err.ErrorType = reqerr.InvalidViewStatementError
	case "E7102":
		err.ErrorType = reqerr.PointsNotInSameRetentionError
	case "E7103":
		err.ErrorType = reqerr.TimestampTooFarFromNowError
	case "E7200", "E7201", "E7205", "E7206":
		err.ErrorType = reqerr.InvalidQuerySql
	case "E9002", "E7204":
		err.ErrorType = reqerr.QueryInterruptError
	case "E7212":
		err.ErrorType = reqerr.ExecuteSqlError
	case "E9001":
		err.ErrorType = reqerr.InternalServerError
	default:
		if code == 401 {
			err.Message = fmt.Sprintf("unauthorized. Please check the local time to ensure the consistent with the server time, if you are using the token, please make sure that token has not expired.")
			err.ErrorType = reqerr.UnauthorizedError
		}
	}

	return err
}
