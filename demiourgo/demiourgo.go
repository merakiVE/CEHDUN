package main

import (
    "os"
    "text/template"
    "encoding/json"
)

const TEMPLATE = `package design

{{$api := .Api}}{{$resources := .Resources}}

import (
    . "github.com/goadesign/goa/design"
    . "github.com/goadesign/goa/design/apidsl"
    . "github.com/goadesign/oauth2/design"
)

var _ = API('{{$api.Mainname}}', func() {                     
    Title('{{$api.Title}}')           
    Description('{{$api.Description}}')       
    Scheme('http')                             
    Host('{{$api.Host}}')

    Consumes('application/json')
    Produces('application/json')

    // OAuth2 requires form encoding
    Consumes('application/x-www-form-urlencoded', func() {
        Package('github.com/goadesign/goa/encoding/form')
    })
    Produces('application/x-www-form-urlencoded', func() {
        Package('github.com/goadesign/goa/encoding/form')
    })
})

{{range $resource := $resources}}

var _ = Resource("{{$resource.Namegroup}}", func() {
        BasePath("{{$resource.Basepath}}")

        {{range $action := $resource.Actions}}

        Action("{{$action.Name}}", func() {
                Description("{{$action.Description}}")
                Routing({{$action.Method}}("{{$action.Route}}"))
                
                {{if $action.Params}}
                Params(func(){
                    {{range $param := $action.Params}}
                    Param("{{$param.Name}}}", {{$param.Type}}, "{{$param.Description}}")
                    {{end}}
                })
                {{end}}

                {{if $action.Payload.Name }}
                Payload({{$action.Payload.Name}})
                {{end}}
                
                {{range $response := $action.Responses}}
                Response({{$response.Name}})
                {{end}}
        })
        {{end}}

})

// Payloads
{{range $action := $resource.Actions}}

{{if $action.Payload.Name }}

var {{$action.Payload.Name}} = Type("{{$action.Payload.Name}}", func() {

    {{range $attribute := $action.Payload.Attributes}}
    Attribute("{{$attribute.Name}}", {{$attribute.Type}}, "{{$attribute.Description}}")
    {{end}}
    
    Required(
    {{range $attribute := $action.Payload.Attributes}}
        "{{$attribute.Name}}",
    {{end}}
    )
    
})
{{end}}
{{end}}


var {{$resource.MediaType.Name}} = MediaType("{{$resource.MediaType.Type}}", func() {
    Description("{{$resource.MediaType.Description}}")
    TypeName("{{$resource.MediaType.Name}}")

    {{range $attribute := $resource.MediaType.Attributes}}
    Attribute("{{$attribute.Name}}", {{$attribute.Type}}, "{{$attribute.Description}}")
    {{end}}
    
    Required(
    {{range $attribute := $resource.MediaType.Attributes}}
        "{{$attribute.Name}}",
    {{end}}
    )

    {{range $view := $resource.MediaType.Views}}
    View("{{$view.Name}})", func() {
        {{range $attribute :=$view.Attributes}}
        Attribute("{{$attribute.Name}}")
        {{end}}
    })
    {{end}}
})
{{end}}
`

func main() {
    
    data := []byte(`
    {
        "api": 
        {
            "mainname": "CEHDUN",
            "title": "Standardized Set of Programs",
            "description": "A test of shit",
            "host": "localhost:8080"
        },

        "resources":
        [
            {
                "namegroup": "routine",
                "basepath": "/routines",
                "mediatype": 
                {
                    "name": "RoutineMedia",
                    "type": "application/json",
                    "description": "A routine of a CEHDUN",
                    "attributes":
                    [
                        {
                            "name": "id",
                            "type": "Integer",
                            "description": "Unique routine ID",
                            "required": true
                        },
                        {
                            "name": "href",
                            "type": "String",
                            "description": "API href for making requests on the routine",
                            "required": true
                        },
                        {
                            "name": "name",
                            "type": "String",
                            "description": "Name of a routine",
                            "required": true
                        },
                        {
                            "name": "message",
                            "type": "String",
                            "description": "Message",
                            "required": false
                        }

                    ],
                    "views":
                    [
                        {
                            "name": "default",
                            "attributes":
                            [
                                {
                                    "name": "id",
                                    "type": "Integer",
                                    "description": "Unique routine ID",
                                    "required": true
                                },
                                {
                                    "name": "href",
                                    "type": "String",
                                    "description": "API href for making requests on the routine",
                                    "required": true
                                },
                                {
                                    "name": "name",
                                    "type": "String",
                                    "description": "Name of a routine",
                                    "required": true
                                },
                                {
                                    "name": "message",
                                    "type": "String",
                                    "description": "Message",
                                    "required": false
                                }  
                            ]
                        }
                    ]

                },

                "actions": 
                [
                    {
                        "name": "show",
                        "description": "Get routine by id",
                        "method": "GET",
                        "Route": "/:routineID",
                        "params": 
                        [
                            {
                                "name": "routineID",
                                "type": "Integer",
                                "description": "Routine ID"
                            }
                        ],
                        "payload": {},
                        "responses": 
                        [
                            {
                                "name": "OK"
                            },
                            {
                                "name": "NotFound"
                            }

                        ]
                    },

                    {
                        "name": "submit",
                        "description": "Post accepts a form encoded request and returns a form encoded response",
                        "method": "POST",
                        "Route": "/",
                        "params": [],
                        "payload": 
                        {
                            "name": "Routine",
                            "attributes":
                            [
                                {
                                    "name": "routineID",
                                    "type": "Integer",
                                    "description": "Routine ID",
                                    "required": true
                                },
                                {
                                    "name": "name",
                                    "type": "String",
                                    "description": "Name Routine name",
                                    "required": true
                                },
                                {
                                    "name": "activities",
                                    "type": "Integer",
                                    "description": "Number of activities",
                                    "required": true
                                }
                            ]
                        },
                        "responses": 
                        [
                            {
                                "name": "OK"
                            },
                            {
                                "name": "NotFound"
                            }

                        ]
                    }
                ]
            }
        ]
    }
    `)
    
    var appTemplate Template
    
    err := json.Unmarshal(data, &appTemplate)

    t, err := template.New("Template-Design").Parse(TEMPLATE)

        if err != nil {
            panic(err)
        }

    w, err := os.Create("design.go")

        if err != nil {
            panic(err)
        }

        defer w.Close()

    err = t.Execute(w, appTemplate)
        
        if err != nil {
            panic(err)
        }

}