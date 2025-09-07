package dto

type CreateCourseDTO struct {
	CourseCode    string `json:"courseCode" binding:"required"`
	Name          string `json:"name" binding:"required"`
	TotalCapacity int    `json:"totalCapacity" binding:"required"`
	Department    string `json:"department" binding:"required"`
	StaffID       uint   `json:"staffId" binding:"required"`
	Semester      string `json:"semester" binding:"required"`
	Credit        int    `json:"credit" binding:"required"`
	Batch         int    `json:"batch" binding:"required"`
}
