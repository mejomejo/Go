package common

type Student struct {
	Name  string `json:"name,omitempty"`
	Sex   string `json:"sex,omitempty"`
	Id    string `json:"id,omitempty"`
	Grade string `json:"grade,omitempty"`
}
