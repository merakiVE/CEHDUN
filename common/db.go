package common

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/merakiVE/CEHDUN/core/types"
    "errors"
    "encoding/json"
)

const SQL_ALL_TABLES_POSTGRES = `
    SELECT row_to_json(object_result) AS result_json
    FROM (
        SELECT array_to_json(array_agg(row_to_json(array_tables))) AS tables
        FROM (
            SELECT table_schema || '.' || table_name AS name,
            (
                SELECT array_to_json(array_agg(row_to_json(array_columns)))
                FROM (
                    SELECT column_name as name, data_type 
                    FROM information_schema.columns WHERE table_name = tables.table_name
                ) AS array_columns
            ) AS columns 
            FROM information_schema.tables AS tables WHERE table_schema != 'pg_catalog' AND table_schema != 'information_schema'
        ) AS array_tables
    ) AS object_result
`

func Connect(db types.DataBase) *gorm.DB{
    
    con, err := gorm.Open(
        "postgres", 
        "host="+db.Host+" user="+db.User+" dbname="+db.Name+" sslmode=disable password="+db.Password)

    if err != nil {
        return nil
    }

    return con
}

func GetTables(db types.DataBase) (types.DatabaseSchema, error){
    var schema types.DatabaseSchema

    con := Connect(db) 

    if con == nil {
        return schema, errors.New("Datos incorrectos y/o Error de conexion")
    }

    switch(db.Type){
        case "postgresql":
            //declare var where save result json tables database
            result_json := new(types.ResultSQLJson)

            con.Raw(SQL_ALL_TABLES_POSTGRES).Scan(result_json)

            //Convert result json in struct DatabaseSchema
            if err := json.Unmarshal([]byte(result_json.GetResultInBytes()), &schema); err != nil {
                panic(err)
            }

            return schema, nil

        default:
            return schema, errors.New("Datos incorrectos")
    }
    
}




