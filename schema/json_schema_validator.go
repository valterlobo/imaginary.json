package schema

import (
	"context"
	"encoding/json"
	"github.com/qri-io/jsonschema"
)

type JsonSchemaValidator interface {
	validator(schemaRaw string, jsonRaw string) (bool, []error)
}

type JsonSchemaValidatorQri struct {
}

func (jsonValidator JsonSchemaValidatorQri) Validator(schemaRaw string, jsonRaw string) (bool, []error) {

	schemaRawBytes := []byte(schemaRaw)
	var listErrors = make([]error, 0)
	rs := &jsonschema.Schema{}

	if err := json.Unmarshal(schemaRawBytes, rs); err != nil {
		listErrors = append(listErrors, err)
		return false, listErrors
	}

	var valid = []byte(jsonRaw)
	ctx := context.TODO()
	errs, err := rs.ValidateBytes(ctx, valid)
	if err != nil {
		listErrors = append(listErrors, err)
		return false, listErrors
	}

	if len(errs) > 0 {
		for _, e := range errs {
			listErrors = append(listErrors, e)
		}
		return false, listErrors
	}

	return true, nil
}
