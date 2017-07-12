package main

import (
    "os"
    "text/template"
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

        {{range $action := $resource.Action}}

        Action("{{$action.Name}}", func() {
                Description("{{$action.Description}}")
                Routing({{$action.Method}}("{{$action.Route}}"))
                
                {{range $response := $action.Response}}
                Response({{$response.Name}})
                {{end}}
        })
        {{end}}

})
{{end}}
`

func main() {

    type Api struct {
        Mainname, Title, Description, Host string
    }

    type Response struct {
        Name string
    }

    type Action struct {
        Name, Description, Method, Route string 
        Response []Response
    }
    
    type Resource struct {
        Namegroup, Basepath string
        Action []Action
    }

    type Payload struct {
        Name, Type, Label string
        Required bool
    }

    type Template struct {
        Api Api
        Resources []Resource
    }

    var api = Api{
        "CEHDUN", 
        "Standardized Set of Programs", 
        "sabeee",
        "localhost:8080",
    }

    var responses = []Response{
        {"Ok"},
        {"NotFound"},
    }

    var actions = []Action{
        {"show", "Get routine by id", "GET", "/:routineID", responses},
        {"view", "Get view by id", "GET", "/:viewID", responses},
    }

    var resources = []Resource{
        {"routine", "/routines", actions},
        {"view", "/views", actions},
    }

    var data = Template{
        api,
        resources,
    }

    t, err := template.New("Template-Design").Parse(TEMPLATE)

        if err != nil {
            panic(err)
        }

    w, err := os.Create("design.go")

        if err != nil {
            panic(err)
        }

        defer w.Close()

    err = t.Execute(w, data)
        
        if err != nil {
            panic(err)
        }

}