package server


import
"net/http"

"github.com/gofiber/fiber/v2"
"github.com/Atarod61/middleware-go-fiber1/middleware"

func Run(){

//Initialize Fiber
	app := fiber.New()

	//Global Middleware
	app.use(middleware.RequestLogger)

//Define Routes.
	app.Get("/",func(ctx*fiber.ctx) error{
		return ctx.status(http.statusok).Jason(fiber.Map{
			"message": "hello I am working",
		   )} 
		)}
		app.Get("/health",func (ctx*fiber.ctx) error {
			
			return ctx.status(http.statusok).Jason(fiber.Map{
				"healthy": true,
			})
		})
		app.Listen(":3000")

}