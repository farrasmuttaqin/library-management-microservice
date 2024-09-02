package models

type MapErrorResponse struct {
	Path         string `json:"path"`
	Err          string `json:"err"`
	FunctionName string `json:"function_name"`
	File         string `json:"file"`
	Line         int    `json:"line"`
	StackTrace   string `json:"stack_trace"`
}

type MapErrorRequest struct {
	QueryParams string `json:"query_params"`
	RequestBody string `json:"request_body"`
}
