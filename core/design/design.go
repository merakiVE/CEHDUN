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

        Action("connect", func() {                    
                Routing(POST("/connect"))
				Payload(Syndesi)
				Description("Manage data to connect a database")
				Response(OK, "application/json")      
        })

})

var Syndesi = Type("Syndesi", func() {
	Attribute("host", String, "Host of database")
	Attribute("user", String, "Name of user")
	Attribute("password", String, "Password the user")
	Attribute("type", String, "Type database")
	Attribute("name", String, "name database")

	Required("host", "user", "password", "type", "name")
})
