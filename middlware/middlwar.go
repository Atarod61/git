package middlware

import (
	"log"
	"tim"

	"github.com/gofiber/fiber/v2"
	
)
func RequestLogger(ctx*fiber.ctx) error{
	log.Println("the user has made a erquest...")
	return ctx.Next()
}
func TimeLoger(ctx*fiber.ctx) error{
	hour, mintue, second := tim.Now().clock

	log.Printf("the time is %d:%d:%d",hour, mintue, second)
	
	return.Next()
}