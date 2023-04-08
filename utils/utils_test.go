package utils

import (
	"testing"
	"reflect"
	"github.com/wrewsama/MyFitnessTomodachi-api/models"
)

func TestGetMatches(t *testing.T) {
	inputs := []models.Food{
		{
			ID: 1, 
			Name: "cartwheel", 
			Unit: "", 
			Calories: 0, 
			Protein: 0, 
			Carbohydrates: 0, 
			Fat: 0,
		},
		{
			ID: 2, 
			Name: "foobar", 
			Unit: "", 
			Calories: 0, 
			Protein: 0, 
			Carbohydrates: 0, 
			Fat: 0,
		},
		{
			ID: 3, 
			Name: "wheel", 
			Unit: "", 
			Calories: 0, 
			Protein: 0, 
			Carbohydrates: 0, 
			Fat: 0,
		},
		{
			ID: 4, 
			Name: "baz", 
			Unit: "", 
			Calories: 0, 
			Protein: 0, 
			Carbohydrates: 0, 
			Fat: 0,
		},
	}

	outputs := GetMatches("whl", inputs)

	expected := []models.Food{
		{
			ID: 1, 
			Name: "cartwheel", 
			Unit: "", Calories: 0, 
			Protein: 0, 
			Carbohydrates: 0, 
			Fat: 0,
		},
		{
			ID: 3, 
			Name: "wheel", 
			Unit: "", 
			Calories: 0, 
			Protein: 0, 
			Carbohydrates: 0, 
			Fat: 0,
		},
	}

	if len(outputs) != len(expected) {
		t.Errorf("Result has wrong len, got: %v, want: %v.", outputs, expected)
	} else {
		for i := 0; i < len(outputs); i++ {
			if !reflect.DeepEqual(outputs[i], expected[i]) {
				t.Errorf("Result was incorrect, got: %v, want: %v.", outputs, expected)
			}
		}
	}
}