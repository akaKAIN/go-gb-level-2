package myreflect

import "testing"

func TestUpdateStruct(t *testing.T) {
	testTable := []struct {
		name     string
		inStruct *User
		inMap    map[string]interface{}
		expect   bool
	}{
		{
			name:     "nil struct",
			inStruct: nil,
			inMap:    map[string]interface{}{"Name": "Ivan", "Age": uint8(13)},
			expect:   true,
		},
		{
			name:     "nil map",
			inStruct: new(User),
			inMap:    nil,
			expect:   true,
		},
		{
			name:     "nil map",
			inStruct: new(User),
			inMap:    map[string]interface{}{"Name": "Ivan", "Age": uint8(13)},
			expect:   false,
		},
	}

	for _, tc := range testTable {
		resultErr := UpdateStruct(tc.inStruct, tc.inMap)

		if resultErr != nil && !tc.expect {
			t.Fatalf("%s: expected error: %t, but got: %v", tc.name, tc.expect, resultErr)
		}
	}
}
