package models

import "encoding/json"

type JsonResponse struct {
	Data      any    `json:"data,omitempty"`
	TotalData *int64 `json:"total_data,omitempty"`
	Message   string `json:"message,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
	Success   bool   `json:"success"`
}

func NewJsonResponse(success bool) *JsonResponse {
	return &JsonResponse{Success: success}
}

func NewError(code, message string) *JsonResponse {
	return &JsonResponse{Success: false, ErrorCode: code, Message: message}
}

func (r *JsonResponse) SetList(data any, total int64) *JsonResponse {
	r.Data = data
	r.TotalData = &total
	return r
}

func (r *JsonResponse) SetData(data any) *JsonResponse {
	r.Data = data
	return r
}

func (r *JsonResponse) SetMessage(message string) *JsonResponse {
	r.Message = message
	return r
}

func (r *JsonResponse) SetError(code string, message string) *JsonResponse {
	r.ErrorCode = code
	r.Message = message
	return r
}

func (r *JsonResponse) Error() string {
	errBytes, _ := json.Marshal(r)
	return string(errBytes)
}
