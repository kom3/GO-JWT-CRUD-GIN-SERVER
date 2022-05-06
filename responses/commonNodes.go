package responses

type ErrorNode struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type EmployeeDetailsNode struct {
	Name    string `json:"name"`
	Age     int32  `json:"age"`
	Skills  string `json:"skills"`
	Address string `json:"address"`
}
