package mapi

//BasicResponse : basic response
type BasicResponse struct {
	Code    int
	Message string
}

//SuccessMessage : return success message
func SuccessMessage() BasicResponse {
	var response = BasicResponse{
		Code:    2000,
		Message: "Succss",
	}

	return response
}

//DataNotFound : return bad request message data not found
func DataNotFound() BasicResponse {
	var response = BasicResponse{
		Code:    4001,
		Message: "Data not found",
	}

	return response
}

//WrongHTTPMethod : invalid http method
func WrongHTTPMethod() BasicResponse {
	var response = BasicResponse{
		Code:    4002,
		Message: "Invalid HTTP method",
	}

	return response
}

//MongoExecutionFailed : failed when try toexecute query/command
func MongoExecutionFailed() BasicResponse {
	var response = BasicResponse{
		Code:    5001,
		Message: "Database execution failed during operation.",
	}
	return response
}

//InternalServerError : something wrong in server or error code
func InternalServerError() BasicResponse {
	var response = BasicResponse{
		Code:    5002,
		Message: "Internal server error",
	}
	return response
}

//ErrorParsingBody : error when decode body to struct
func ErrorParsingBody() BasicResponse {
	var response = BasicResponse{
		Code:    5003,
		Message: "Error when parsing body",
	}
	return response
}
