package testutil

// TestStruct a structure for tests
type TestStruct struct {
	ID      int64
	Name    string
	IsValid bool
}

// Equal the current TestStruct is equal to target TestStruct or not
func (ts *TestStruct) Equal(target *TestStruct) bool {
	if ts == nil || target == nil {
		return ts == target
	}
	return ts.Name == target.Name && ts.ID == target.ID && ts.IsValid == target.IsValid
}
