package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/eduwr/fingerprint/backend/models"
	"github.com/ua-parser/uap-go/uaparser"

	"github.com/eduwr/fingerprint/backend/lib"
	"github.com/gofiber/fiber/v2"
)

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

	parser, err := uaparser.New("./regexes.yaml")
	if err != nil {
		println("error getting the parser")
		log.Fatal(err)
	}
	headers := c.GetReqHeaders()

	println("Security Headers:")

	// User Agent
	client := parser.Parse(headers["User-Agent"][0])

	fmt.Println(client.UserAgent.Family)
	fmt.Println(client.UserAgent.Major)
	fmt.Println(client.UserAgent.Minor)
	fmt.Println(client.UserAgent.Patch)
	fmt.Println(client.Os.Family)
	fmt.Println(client.Os.Major)
	fmt.Println(client.Os.Minor)
	fmt.Println(client.Os.Patch)
	fmt.Println(client.Os.PatchMinor)
	fmt.Println(client.Device.Family)

	// Is Mobile
	h, hExists := headers["Sec-Ch-Ua-Mobile"]
	p, pExists := headers["Sec-Ch-Ua-Platform"]

	if hExists {
		// If the header exists, join its values
		println("sec-ch-ua-mobile:", strings.Join(h, ", "))
	} else {
		println("sec-ch-ua-mobile header not found")
	}

	if pExists {
		// If the header exists, join its values
		println("sec-ch-ua-platform:", strings.Join(p, ", "))
	} else {
		println("sec-ch-ua-platform header not found")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Fingerprint data received successfully",
		"data":    data,
	})
}
