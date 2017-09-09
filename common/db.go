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
        panic(err)
    }

    return con
}




