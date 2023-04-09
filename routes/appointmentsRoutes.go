package routes

func (routes Routes) defineAppointmentsRoutes() {
	app := routes.app
	controller := routes.appointmentsController

	appointmentsRoutes := app.Group("/api/appointments")
	appointmentsRoutes.Get("/", controller.GetAppointments)
	appointmentsRoutes.Post("/", controller.AddAppointment)
	appointmentsRoutes.Get("/:id", controller.GetAppointmentById)
}
