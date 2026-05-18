package errs

import "net/http"

// AppError โครงสร้าง Error ของเราเองที่จะเก็บ Status Code และ Message ไว้คู่กัน
type AppError struct {
	Code    int
	Message string
}

// ทำให้ AppError เป็น error type ในภาษา Go (implement error interface)
func (e AppError) Error() string {
	return e.Message
}

// --- ฟังก์ชันช่วยเหลือสำหรับสร้าง Error แต่ละประเภท ---

func NewNotFoundError(message string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "Unexpected error occurred",
	}
}

func NewBadRequestError(message string) error {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewConflictError(message string) error {
	return AppError{
		Code:    http.StatusConflict,
		Message: message,
	}
}
