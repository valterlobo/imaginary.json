package schema

import (
	"fmt"
	"testing"
)

func TestValidator(t *testing.T) {


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
	jsonValidator := JsonSchemaValidatorQri{}
	valido, erros := jsonValidator.ValidatorStr(schemaRaw, jsonRow)

	if valido {
		println(" VALIDO ")
	} else {
		for _, v := range erros {

			println(" " + v)
		}
	}

	fmt.Println("----------------------------------------")

	var invalid = `{
    "firstName" : "Prince"  ,
    }`

	valido2, erros2 := jsonValidator.ValidatorStr(schemaRaw, invalid)

	if valido2 {
		println(" VALIDO ")
	} else {
		for _, v := range erros2 {

			println(" " + v)
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

	valido3, erros3 := jsonValidator.ValidatorStr(schemaRaw, invalidFriend)

	if valido3 {
		println(" VALIDO ")
	} else {
		for _, v := range erros3 {

			println(" " + v)
		}
	}
	fmt.Println("----------------------------------------")
	fmt.Println("/FIM TESTE")

}

