package utils

import "regexp"

const (
	SUCCESS  = "success"
	FAILURE  = "failure"
	CANCELED = "canceled"
	PENDING  = "pending"

	OK                 = 2000
	BAD_REQUEST        = 4000
	LOGIN_FAILED       = 4001
	SESSION_EXPIRED    = 4002
	PERMISSION_DENIED  = 4003
	NOT_FOUND          = 4004
	SERVER_ERROR       = 5000
	GRPC_SERVICE_ERROR = 6000
	MYSQL_ERROR        = 7000
)

var (
	YES_NO_OPTIONS = []string{"yes", "no"}
	MAIL_REGEX     = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)
