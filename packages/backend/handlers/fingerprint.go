package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/eduwr/fingerprint/backend/models"

	"github.com/eduwr/fingerprint/backend/lib"
	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func HandleFingerprint(c *fiber.Ctx) error {
	var data models.FingerprintData
	validate := lib.NewValidator()

	if err := c.BodyParser(&data); err != nil {
		log.Printf("Error parsing request body")
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: fiber.ErrBadRequest.Message,
		}
	}

	if errs := validate.Validate(data); len(errs) > 0 {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"Field '%s' with value '%v' failed validation: %s (rule: %s)",
				err.Field,
				err.Value,
				err.Message,
				err.Tag,
			))
		}

		println(strings.Join(errMsgs, "\n"))

		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: fiber.ErrBadRequest.Message,
		}
	}

	log.Printf("Received fingerprint data: %+v", data)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Fingerprint data received successfully",
		"data":    data,
	})
}
