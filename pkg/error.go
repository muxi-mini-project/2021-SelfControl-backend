package error

type Error struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

var (
	BadRequest        = &Error{ErrorCode: "00001", Message: "Fail."}
	ParamBadRequest   = &Error{ErrorCode: "00002", Message: "Lack Param Or Param Not Satisfiable."}
	FrequentRequests1 = &Error{ErrorCode: "00003", Message: "Post over 2 requirements within 1 minute"}
	FrequentRequests2 = &Error{ErrorCode: "00004", Message: "Post over 15 requirements within 1 day"}

	TokenInvalid = &Error{ErrorCode: "10001", Message: "Token Invalid."}
	Unauthorized = &Error{ErrorCode: "10002", Message: "Unauthorized."}
	BlackList    = &Error{ErrorCode: "10003", Message: "You are in our blacklist."}

	LoginError = &Error{ErrorCode: "20001", Message: "Password or account wrong."}
	LoginAgain = &Error{ErrorCode: "20002", Message: "Please try again."}

	ServerError = &Error{ErrorCode: "30001", Message: "Fail."}
)
