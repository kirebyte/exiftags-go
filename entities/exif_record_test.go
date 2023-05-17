package entities

import (
	"testing"
)

func TestExifRecord_String(t *testing.T) {
	r := ExifRecord{"test.jpg", "12.3456", "65.4321"}

	expected := "test.jpg 12.3456 65.4321"
	result := r.String()

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestExifRecord_Json(t *testing.T) {
	r := ExifRecord{"test.jpg", "12.3456", "65.4321"}

	expected := `{"filename":"test.jpg","latitude":"12.3456","longitude":"65.4321"}`
	result, err := r.Json()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestExifRecord_ToArray(t *testing.T) {
	r := ExifRecord{"test.jpg", "12.3456", "65.4321"}

	expected := []string{"test.jpg", "12.3456", "65.4321"}
	result := r.ToArray()

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected '%s', got '%s'", expected[i], v)
		}
	}
}
