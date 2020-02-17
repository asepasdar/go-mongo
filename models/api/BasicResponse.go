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


//MongoExecutionFailed : failed when try toexecute query/command
func MongoExecutionFailed() BasicResponse {
	var response = BasicResponse{
		Code:    5001,
		Message: "Database execution failed during operation.",
	}
	return response
}
