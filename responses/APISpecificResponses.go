package responses

type GetAllEmployeesResp struct {
	Error   ErrorNode   `json:"error"`
	Results interface{} `json:"results"` // interface acts as type 'any'
}

type GetEmployeeByIdResp struct {
	Error ErrorNode
	// Results EmployeeDetailsNode
	Results EmployeeDetailsNode // interface acts as type 'any'
}

type AddEmployee struct {
	Error ErrorNode
}

type AddEmployees struct {
	Error ErrorNode
}

type UpdateEmployee struct {
	Error ErrorNode
}
