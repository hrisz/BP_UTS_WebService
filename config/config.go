package config

import "github.com/gofiber/fiber/v2"

var Iteung = fiber.Config{
	Prefork:       false,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "UTS WebService",
	AppName:       "Message Router",
}
