package update_api_1c

type RequestError struct {
	Timestamp int64  `json:"timestamp"`
	Status    int    `json:"status"`
	Err       string `json:"error"`
	Exception string `json:"exception"`
	Errors    []struct {
		Codes     []string `json:"codes"`
		Arguments []struct {
			Codes          []string    `json:"codes"`
			Arguments      interface{} `json:"arguments"`
			DefaultMessage string      `json:"defaultMessage"`
			Code           string      `json:"code"`
		} `json:"arguments"`
		DefaultMessage string `json:"defaultMessage"`
		ObjectName     string `json:"objectName"`
		Field          string `json:"field"`
		RejectedValue  string `json:"rejectedValue"`
		BindingFailure bool   `json:"bindingFailure"`
		Code           string `json:"code"`
	} `json:"errors"`
	Message string `json:"message"`
	Path    string `json:"path"`
}

func (e RequestError) Error() string {

	return e.Message

}
