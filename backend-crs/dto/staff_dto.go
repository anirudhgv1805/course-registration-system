package dto

type StaffResponse struct {
	Username       string `json:"username"`
	StaffId        string `json:"staffId"`
	Department     string `json:"department"`
	Email          string `json:"email"`
	IsClassAdvisor bool   `json:"isClassAdvisor"`
	Section        string `json:"section"`
	Batch          string `json:"batch"`
}
