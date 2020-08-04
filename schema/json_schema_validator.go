package schema

import (
	"context"
	"encoding/json"
	"github.com/qri-io/jsonschema"
)

type JsonSchemaValidator interface {
	validator(schemaRaw string, jsonRaw string) (bool, []string)
}

type JsonSchemaValidatorQri struct {
}

func (jsonValidator JsonSchemaValidatorQri) ValidatorStr(schemaRaw string, jsonRaw string) (bool, []string) {

	schemaRawBytes := []byte(schemaRaw)
	var jsonBytes = []byte(jsonRaw)
	return jsonValidator.ValidatorBytes (schemaRawBytes , jsonBytes)
}


func (jsonValidator JsonSchemaValidatorQri) ValidatorBytes(schemaBytes []byte, jsonBytes  []byte) (bool, []string) {

	var listErrors = make([]string, 0)
	rs := &jsonschema.Schema{}

	if err := json.Unmarshal(schemaBytes, rs); err != nil {
		listErrors = append(listErrors, err.Error())
		return false, listErrors
	}

	ctx := context.TODO()
	errs, err := rs.ValidateBytes(ctx, jsonBytes)
	if err != nil {
		listErrors = append(listErrors, err.Error())
		return false, listErrors
	}

	if len(errs) > 0 {
		for _, e := range errs {
			listErrors = append(listErrors, e.Error())
		}
		return false, listErrors
	}

	return true, nil
}
