package berr

// Json Errors
const (
	Marshal   = Code("MarshalError")
	Unmarshal = Code("UnmarshalError")
)

// Others
const (
	InternalServer  = Code("InternalServerError")
	Database        = Code("DatabaseError")
	Timeout         = Code("Timeout")
	NotFound        = Code("NotFound")
	InvalidArgument = Code("InvalidArgument")
)
