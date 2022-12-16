package model

type Employee struct {
	EmployeeNo        int     `json:"employee_no"`
	EmployeeName      string  `json:"employee_name"`
	EmployeeAge       int     `json:"employee_age"`
	EmployeeGender    string  `json:"employee_gender"`
	EmployeeEnterTime string  `json:"employee_enterTime"`
	EmployeeAddress   string  `json:"employee_address"`
	EmployeePost      string  `json:"employee_post"`
	EmployeeSalary    float64 `json:"employee_salary"`
	EmployeePower     int     `json:"employee_power"`
}

type PostEmployee struct {
	PostId   string     `json:"post_id"`
	Employee []Employee `json:"data"`
}
