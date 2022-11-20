package encoding

import (
	"testing"
	"time"
)

func TestJson_Serialize(t *testing.T) {
	s := &jsonSerializer{}
	ms := getMyStructCase()
	data, err := s.Serialize(ms)
	if err != nil {
		t.Errorf("json serialize error, err=%v", err)
		return
	}
	expect := ms.json()
	actual := string(data)
	if actual != expect {
		t.Errorf("json serialize error, expect to get %s, but get %s", expect, actual)
	}
}

func TestJson_Deserialize(t *testing.T) {
	s := &jsonSerializer{}
	ms := getMyStructCase()
	var target MyStruct
	err := s.Deserialize([]byte(ms.json()), &target)
	if err != nil {
		t.Errorf("json deserialize error, err=%v", err)
		return
	}
	if !ms.equal(target) {
		t.Errorf("json deserialize error, current case is not equal to target")
	}
}

type MyStruct struct {
	Name       string        `json:"name"`
	ID         int64         `json:"id"`
	IsValid    bool          `json:"is_valid"`
	CreateTime time.Time     `json:"create_time"`
	Expires    time.Duration `json:"expires"`

	jsonData string
}

func (ms MyStruct) json() string {
	return ms.jsonData
}

func (ms MyStruct) equal(target MyStruct) bool {
	return ms.Name == target.Name && ms.ID == target.ID && ms.CreateTime == target.CreateTime && ms.Expires == target.Expires
}

func getMyStructCase() MyStruct {
	t, _ := time.Parse(time.RFC3339, "2022-11-25T05:06:07Z")
	return MyStruct{
		Name:       "go",
		ID:         1,
		IsValid:    true,
		CreateTime: t,
		Expires:    time.Minute * 3,
		jsonData:   `{"name":"go","id":1,"is_valid":true,"create_time":"2022-11-25T05:06:07Z","expires":180000000000}`,
	}
}
