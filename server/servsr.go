package server


import
"net/http"

"github.com/gofiber/fiber/v2"

func Run(){

	app := fiber.New()
	app.use(middleware.RequestLogger)

	app.Get("/",func(ctx*fiber.ctx) error{
		return ctx.status(http.statusok).Jason(fiber.Map{
			"message": "hello I am working",
		   )} 
		)}
		app.Get("/health",func (ctx*fiber.ctx) error {
			return ctx.status(http.statusok).Jason(fiber.Map{
				"healthy": true
			})
		})
		app.Listen(":3000")

}