package handler

import (
	e "crud/error"
	"crud/model"

	"github.com/gofiber/fiber/v2"
	log "crud/logger"
)

func ErrorResponseHandler(c *fiber.Ctx, err error) error {
	// DefinedError
	if definedError, ok := err.(e.DefinedError); ok {
		log.Logger().Warn("Handle DefinedError.")
		errorResponse := model.ErrorResponse(definedError.Code, definedError.Message, definedError.Detail)
		return c.Status(definedError.HttpStatus).JSON(errorResponse)
	}

	// FieldValidationError
	if validationError, ok := err.(e.FieldValidationError); ok {
		log.Logger().Warn("Handle FieldValidationError.")
		errorResponse := model.ErrorResponse(validationError.Code, validationError.Message, validationError.FieldErrors)
		return c.Status(400).JSON(errorResponse)
	}

	// undefined error
	log.Logger().Errorf("Handle UndefinedError. Caused: %v.", err)
	errorResponse := model.ErrorResponse("500", err.Error(), nil)
	return c.Status(500).JSON(errorResponse)
}
