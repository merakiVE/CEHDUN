package design



import (
    . "github.com/goadesign/goa/design"
    . "github.com/goadesign/goa/design/apidsl"
)

var _ = API("CEHDUN", func() {                     
    Title("Standardized Set of Programs")           
    Description("A test of shit")       
    Scheme("http")                             
    Host("localhost:8089")

    Consumes("application/json")
    Produces("application/json")

    Origin("*", func() {
            Methods("GET", "POST", "PUT", "PATCH", "DELETE")
            Headers("content-type")
            MaxAge(600)
            Credentials()
    })
})



var _ = Resource("routine", func() {
        BasePath("/routines")

        

        Action("show", func() {
                Description("Get routine by id")
                Routing(GET("/:cedula"))
                
                
                Params(func(){
                    
                    Param("cedula", String, "Routine ID")
                    
                })
                

                
                
                
                Response(OK, "application/json")
                
        })
        

        Action("submit", func() {
                Description("Post accepts a form encoded request and returns a form encoded response")
                Routing(POST("/"))
                
                

                
                Payload(Routine)
                
                
                
                Response(OK, "application/json")
                
        })
        

})

// Payloads







var Routine = Type("Routine", func() {

    
    Attribute("routineID", Integer, "Routine ID")
    
    Attribute("name", String, "Name Routine name")
    
    Attribute("activities", Integer, "Number of activities")
    
    
    Required(
    
        "routineID",
    
        "name",
    
        "activities",
    
    )
    
})




var RoutineMedia = MediaType("application/json", func() {
    Description("A routine of a CEHDUN")
    TypeName("RoutineMedia")

    
    Attribute("nombre", String, "Unique routine ID")
    
    Attribute("apellido", String, "API href for making requests on the routine")
    
    Attribute("cedula", String, "Name of a routine")
    
    
    Required(
    
        "nombre",
    
        "apellido",
    
        "cedula",
    
    )

    
    View("default", func() {
        
        Attribute("nombre")
        
        Attribute("apellido")
        
        Attribute("cedula")
        
    })
    
})

