package dto

type CreateDepartment struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Block string `json:"block"`
}

type UpdateDepartment struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Block string `json:"block"`
}

type DepartmentResponse struct {
	Name              string                `json:"name"`
	Code              string                `json:"code"`
	Block             string                `json:"block"`
	ApplicableCourses []CourseApplicableDTO `gorm:"foreignKey:DepartmentID"`
}
