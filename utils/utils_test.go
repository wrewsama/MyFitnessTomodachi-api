package utils

import (
	"testing"
	"reflect"
	"github.com/wrewsama/MyFitnessTomodachi-api/models"
)

func TestGetMatches(t *testing.T) {
	inputs := []models.Food{
		{1, "cartwheel", "", 0, 0, 0, 0},
		{2, "foobar", "", 0, 0, 0, 0},
		{3, "wheel", "", 0, 0, 0, 0},
		{4, "baz", "", 0, 0, 0, 0},
	}

	outputs := GetMatches("whl", inputs)

	expected := []models.Food{
		{1, "cartwheel", "", 0, 0, 0, 0},
		{3, "wheel", "", 0, 0, 0, 0},
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