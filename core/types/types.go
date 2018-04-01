package types

import (
	"time"
)

type JsonObject map[string]interface{}

type JsonArray []JsonObject

type Timestamps struct {
	CreatedAt time.Time `json:"created_at" on_create:"set,auto_now"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseAPI struct {
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
    Errors  interface{} `json:"errors"`
}

type Contact struct {
    Name string  `json:"name"`
    Email string `json:"email"`
}

type Api struct {
    Name string `json:"name, omitempty"`
    Title string `json:"title, omitempty"` 
    Description string `json:"description,omitempty"`
    Contact Contact `json:"contact"`
    Host string `json:"host, omitempty"`
    Port int `json:"port"`
    BasePath string `json:"basePath, omitempty"`
}

type Response struct {
    Name string `json:"name"`
}

type Param struct {
    Name string `json:"name"`
    Type string `json:"type"`
    Description string `json:"description,omitempty"` 
}

type Attribute struct {
    Name string `json:"name"`
    Type string `json:"type"`
    Description string `json:"description,omitempty"` 
    Required bool `json:"required"`
}

type Payload struct {
    Name string `json:"name"` 
    Attributes []Attribute `json:"attributes"`
}

type Action struct {
    Name string `json:"name"` 
    Description string `json:"description"` 
    Method string `json:"method"` 
    Route string `json:"route"`
    Params []Param `json:"params,omitempty"`
    Responses []Response `json:"responses"`
    Payload Payload `json:"payload,omitempty"`
}

type View struct {
    Name string `json:"name"`
    Attributes []Attribute `json:"attributes"`
}

type MediaType struct {
    Name string `json:"name"`
    Type string `json:"type"`
    Description string `json:"description"` 
    Attributes []Attribute `json:"attributes"` 
    Views []View `json:"views"` 
}

type Resource struct {
    Namegroup string `json:"namegroup"` 
    Basepath string `json:"basepath"`
    MediaType MediaType `json:"mediatype"` 
    Actions []Action `json:"actions"`
}

type Template struct {
    Api Api `json:"api"`
    Resources []Resource `json:"resources"`
}

type DataBase struct {
    Host     string `json:"host"`
    User     string `json:"user"`
    Password string `json:"password"`
    Type     string `json:"type"`
    Name     string `json:"name"`
    Port     int    `json:"port"`
}


type ResultSQLJson struct {
    ResultJson string 
}

func (this *ResultSQLJson) GetResultInBytes()([]byte){
    return []byte(this.ResultJson)
}


type DatabaseSchema struct {
    Tables []Table `json:"tables"`
}

type Table struct {
    Name string `json:"name"`
    Columns []Column `json:"columns"`
}

type Column struct {
    Name string `json:"name"`
    DataType string `json:"data_type"`
}

type Error struct {
    Message string `json:"message"`
}
