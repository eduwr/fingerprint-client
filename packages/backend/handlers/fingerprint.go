package handlers

import (
	"log"
	"strings"

	"github.com/eduwr/fingerprint/backend/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func HandleFingerprint(c *fiber.Ctx) error {
	var data models.FingerprintData
	validate := validator.New()

	if err := c.BodyParser(&data); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	if err := validate.Struct(data); err != nil {
		var errors []ValidationError
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, ValidationError{
				Field:   strings.ToLower(err.Field()),
				Message: err.Tag(),
			})
		}
		log.Printf("Validation errors: %+v", errors)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": errors,
		})
	}

	log.Printf("Received fingerprint data: %+v", data)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Fingerprint data received successfully",
		"data":    data,
	})
}
