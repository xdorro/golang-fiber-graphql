package util

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ResponseSuccess : returning json structure for success request
func ResponseSuccess(message string, data interface{}) error {
	status := fiber.StatusOK

	return jsonResponse(status, message, data)
}

// ResponseNotFound : returning json structure for notfound request
func ResponseNotFound(message string) error {
	status := fiber.StatusNotFound

	return jsonResponse(status, message, nil)
}

// ResponseError : returning json structure for error request
func ResponseError(message string, data interface{}) error {
	status := fiber.StatusInternalServerError

	return jsonResponse(status, message, data)
}

// ResponseUnauthorized : returning json structure for validator error request
func ResponseUnauthorized(message string, data interface{}) error {
	status := fiber.StatusUnauthorized

	return jsonResponse(status, message, data)
}

// ResponseBadRequest : returning json structure for validator error request
func ResponseBadRequest(message string, data interface{}) error {
	status := fiber.StatusBadRequest

	return jsonResponse(status, message, data)
}

func jsonResponse(status int, message string, data interface{}) error {
	msg, _ := JSONMarshal(&Response{
		Status:  status,
		Message: message,
		Data:    data,
	})

	return fiber.NewError(status, string(msg))
}
