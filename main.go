package main

import (
	"fmt"
	"github.com/valterlobo/imaginary.json/schema"
)

func main() {

	//schema.JsonSchemaValidator2{}
	schemaRaw := `{
    "$id": "https://qri.io/schema/",
    "$comment" : "sample comment",
    "title": "Person",
    "type": "object",
    "properties": {
        "firstName": {
            "type": "string"
        },
        "lastName": {
            "type": "string"
        },
        "age": {
            "description": "Age in years",
            "type": "integer",
            "minimum": 0
        },
        "friends": {
          "type" : "array",
          "items" : { "title" : "REFERENCE", "$ref" : "#" }
        }
    },
    "required": ["firstName", "lastName"]
  }`

	jsonRow := `{
    "firstName" : "George" ,
     "lastName"  : "Lucas"
    }`

	fmt.Println("TESTE")
	fmt.Println("----------------------------------------")
	jsonValidator := schema.JsonSchemaValidatorQri{}
	valido, erros := jsonValidator.Validator(schemaRaw, jsonRow)

	if valido {
		println(" VALIDO ")
	} else {
		for _, v := range erros {

			println(" " + v.Error())
		}
	}

	fmt.Println("----------------------------------------")


	var invalid = `{
    "firstName" : "Prince"  ,
    }`

	valido2, erros2 := jsonValidator.Validator(schemaRaw, invalid)

	if valido2 {
		println(" VALIDO ")
	} else {
		for _, v := range erros2 {

			println(" " + v.Error())
		}
	}
	fmt.Println("----------------------------------------")

	var invalidFriend = `{
    "firstName" : "Jay",
    "lastName" : "Z",
    "friends" : [{
      "firstName" : "Nas"
      }]
    }`

	valido3, erros3 := jsonValidator.Validator(schemaRaw, invalidFriend)

	if valido3 {
		println(" VALIDO ")
	} else {
		for _, v := range erros3 {

			println(" " + v.Error())
		}
	}
	fmt.Println("----------------------------------------")
	fmt.Println("/FIM TESTE")



}
