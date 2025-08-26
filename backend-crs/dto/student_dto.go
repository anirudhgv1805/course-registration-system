package dto

type StudentResponse struct {
	Username   string `json:"username" gorm:"unique"`
	RegisterNo string `json:"registerno" gorm:"unique"`
	Department string `json:"department"`
	Email      string `json:"email"`
	Section    string `json:"section"`
	Batch      int    `json:"batch"`
}