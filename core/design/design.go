package design

import (
        . "github.com/goadesign/goa/design"
        . "github.com/goadesign/goa/design/apidsl"
)

var _ = API("CEHDUN", func() {
        Title("Standardized set of development tools for the union and interconnection of Neurones.")
        Description("Core that process information to create a API")
        Scheme("http")
        Host("localhost:8080")

        Consumes("application/json")
        Produces("application/json")
})

var _ = Resource("syndesi", func() {                
        BasePath("/syndesi")                       
        DefaultMedia(SyndesiMedia)

        Action("connect", func() {                    
                Routing(POST("/connect"))
				Payload(Syndesi)
				Description("Manage data to connect a database")
				Response(OK, SyndesiMedia)      
        })

})

var Syndesi = Type("Syndesi", func() {
	Attribute("name", String, "Name of user")
	Attribute("password", String, "Password of the user")
	Attribute("typeDB", String, "Type of database")

	Required("name", "password", "typeDB")
})

var SyndesiMedia = MediaType("application/json", func() {
	TypeName("SyndesiMedia")
	Attributes(func() {
		Attribute("message", String, "TSucces or Error")
		Required("message")
	})
	View("default", func() {
		Attribute("message")
	})
})
