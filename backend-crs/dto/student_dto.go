package dto

type StudentResponse struct {
	Username   string             `json:"username" gorm:"unique"`
	RegisterNo string             `json:"registerno" gorm:"unique"`
	Department DepartmentResponse `json:"department"`
	Email      string             `json:"email"`
	Section    string             `json:"section"`
	Batch      int                `json:"batch"`
	Role       string             `json:"role"`
}

type RegisterStudentDTO struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	RegisterNo   string `json:"registerno" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	DepartmentID uint   `json:"departmentId" binding:"required"`
	Section      string `json:"section" binding:"required"`
	Batch        int    `json:"batch" binding:"required"`
}
