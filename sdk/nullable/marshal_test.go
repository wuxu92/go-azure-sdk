package nullable_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

type TestStruct struct {
	Name    *string  `json:"name,omitempty"`
	Age     *float64 `json:"age,omitempty"`
	Address *string  `json:"address,omitempty"`
}

// MarshalJSON implements json.Marshaler.
func (t TestStruct) MarshalJSON() ([]byte, error) {
	return nullable.MarshalNullableStruct(t)
}

var _ json.Marshaler = TestStruct{}

func TestMarshalNullable(t *testing.T) {
	name := "John Doe"
	age := 30.0
	obj := TestStruct{
		Name: &name,
		Age:  &age,
	}
	obj.Address = nullable.NullValue[*string]()

	expected := map[string]interface{}{
		"name":    "John Doe",
		"age":     30.0,
		"address": nil,
	}

	data, err := json.Marshal(obj)
	if err != nil {
		t.Fatalf("MarshalNullable returned an error: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("json.Unmarshal returned an error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
