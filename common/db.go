package common

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/merakiVE/CEHDUN/core/types"
)

func Connect(db types.DataBase) *gorm.DB{
    
    con, err := gorm.Open(
        "postgres", 
        "host="+db.Host+" user="+db.User+" dbname="+db.Name+" sslmode=disable password="+db.Password)

    if err != nil {
        return nil
    }

    return con
}

func GetData(db types.DataBase) interface{}{
    
    var err types.Error 
    var name_tables []types.Name_table  
    var tables []types.Table

    con := Connect(db) 

    if con == nil {
        err.Message = "Datos incorrectos y/o Error de conexion"
        return err
    }

    data_base := types.BaseData{Tables: nil} 

    switch(db.Type){
        case "postgresql":

            con.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema != 'pg_catalog' AND table_schema != 'information_schema'").Scan(&name_tables)

            for i := range name_tables {
                values := types.Table{Table_name: name_tables[i].Table_name, Columns: nil}
                con.Raw("SELECT column_name, data_type FROM information_schema.columns WHERE table_name = ?", name_tables[i].Table_name).Scan(&values.Columns)
                tables = append(tables, values)
                data_base.Tables = append(tables)
            }
            
            data_base.Tables = append(tables)

            return data_base

        default:

            err.Message = "Datos incorrectos"

            return err
    }
    
}




