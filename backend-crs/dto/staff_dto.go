package dto

type StaffResponse struct {
	Username       string `json:"username"`
	StaffId        string `json:"staffId"`
	Department     string `json:"department"`
	Email          string `json:"email"`
	IsClassAdvisor bool   `json:"isClassAdvisor"`
	Section        string `json:"section"`
	Batch          int    `json:"batch"`
	Role           string `json:"role"`
}

type StaffRegisterRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	StaffId        string `json:"staffId"`
	DepartmentID   string `json:"departmentId"`
	Batch          int    `json:"batch"`
	Email          string `json:"email"`
	IsClassAdvisor bool   `json:"isClassAdvisor"`
	Section        string `json:"section"`
}
