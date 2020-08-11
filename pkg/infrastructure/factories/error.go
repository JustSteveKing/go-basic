package factories

import (
	"net/http"
	"strconv"

	problem "github.com/JustSteveKing/go-api-problem"
	response "github.com/JustSteveKing/go-http-response"
)

// APIError is a general error to be returned from the API
func APIError(rw http.ResponseWriter, title string, detail string, status int, code string, contentType string) {
	response.Send(
		rw,
		status,
		&problem.APIProblem{
			Title:  title,
			Detail: detail,
			Status: strconv.Itoa(status),
			Code:   code,
		},
		contentType,
	)
}
