package update_api_1c

import "fmt"

type ErrorResponse struct {
	ErrorName    string `json:"errorName"`
	ErrorMessage string `json:"errorMessage"`
}

func (c *ErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", c.ErrorName, c.ErrorMessage)
}

type ProgramVersion struct {
	ProgramName   string `json:"programName"`
	VersionNumber string `json:"versionNumber"`
}
