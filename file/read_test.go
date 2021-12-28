package file

import "testing"

func TestReadFullData(t *testing.T) {
	t.Run("should pass", func(t *testing.T) {
		data, err := ReadFullData("testdata/input.txt")
		if err != nil {
			t.Fatal(err)
		}
		if data[0] != "1" {
			t.Error("Data 0 doesn't match")
		}
		if data[1] != "2" {
			t.Error("Data 1 doesn't match")
		}
		if data[2] != "3" {
			t.Error("Data 2 doesn't match")
		}
	})
	t.Run("should have error", func(t *testing.T) {
		_, err := ReadFullData("input.txt")
		if err == nil {
			t.Error("did not receive error")
		}
	})
}
