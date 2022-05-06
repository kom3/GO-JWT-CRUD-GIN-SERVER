package models

type EmployeeDetails struct {
	Name    string `json:"name" validate:"required"`
	Age     int32  `json:"age,omitempty"`
	Skills  string `json:"skills" validate:"required"`
	Address string `json:"address" validate:"required"`
}
