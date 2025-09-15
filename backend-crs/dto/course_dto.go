package dto

type CreateCourseDTO struct {
	CourseCode              string `json:"courseCode" binding:"required"`
	Name                    string `json:"name" binding:"required"`
	TotalCapacity           int    `json:"totalCapacity" binding:"required"`
	DepartmentID            uint   `json:"departmentId" binding:"required"`
	StaffID                 uint   `json:"staffId" binding:"required"`
	Semester                int    `json:"semester" binding:"required"`
	Credit                  int    `json:"credit" binding:"required"`
	Batch                   int    `json:"batch" binding:"required"`
	ApplicableDepartmentIDs []uint `json:"applicableDepartmentIds"`
}

type CourseResponse struct {
	CourseID                uint               `json:"courseId"`
	CourseCode              string             `json:"courseCode"`
	Name                    string             `json:"name"`
	TotalCapacity           int                `json:"totalCapacity"`
	CurrentlyFilled         int                `json:"currentlyFilled"`
	DepartmentID            uint               `json:"departmentId"`
	OfferedBy               DepartmentResponse `json:"department"`
	StaffID                 uint               `gorm:"uniqueIndex"`
	HandledByStaff          StaffResponse      `json:"handledByStaff"`
	Semester                int                `json:"semester"`
	Credit                  int                `json:"credit"`
	Batch                   int                `json:"batch"`
	ApplicableDepartmentIDs []uint             `json:"applicableDepartmentIds"`
}

type CourseApplicableDTO struct {
	CourseID     uint
	DepartmentID uint

	Course     CourseResponse
	Department DepartmentResponse
}
