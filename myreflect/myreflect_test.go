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

func BenchmarkUpdateStruct(b *testing.B) {
	user := &User{
		Age:    uint8(1),
		Gender: uint8(0),
		Email:  "default@mail.ru",
		Name:   "baseUser",
	}
	m := map[string]interface{}{
		"Age": uint8(20), "Gender": uint8(1), "Email": "1@gmail.com", "Name": "NewUser",
	}

	for i := 0; i < b.N; i++ {
		_ = UpdateStruct(user, m)
	}
}

func BenchmarkUpdateStruct_WithError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = UpdateStruct(nil, nil)
	}
}
