package helper

import "crud-postgres-orm/data/response"

func WebResponseFormatter(code int, status string, data interface{}, message string, errors []string) response.WebResponse {
	return response.WebResponse{
		Code:    code,
		Status:  status,
		Data:    data,
		Message: message,
		Errors:  errors,
	}
}
