package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("LTime API", func() {
	Title("The LTime API")
	Description("An API for the LTime backend")
	Contact(func() {
		Name("Lucas Vikström")
		Email("lucas.vikstrom@gmail.com")
	})
	Host("localhost:8080")
	Scheme("http")
	BasePath("/api/")
	Origin("*", func() {
		Headers("Content-Type")
		Methods("GET", "POST", "PATCH", "DELETE", "PUT", "OPTION")
	})
	Consumes("application/json")
	Produces("application/json")
})
