package testutil

// testStruct a structure for tests
type testStruct struct {
	ID      int64
	Name    string
	IsValid bool
}

// Equal the current testStruct is equal to target testStruct or not
func (ts *testStruct) Equal(target *testStruct) bool {
	if ts == nil || target == nil {
		return ts == target
	}
	return ts.Name == target.Name && ts.ID == target.ID && ts.IsValid == target.IsValid
}

type testCycleStruct struct {
	Self *testCycleStruct
}
