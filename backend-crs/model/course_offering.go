package model


type CourseOffering struct {
	CourseID     uint `gorm:"primaryKey"`
	DepartmentID uint `gorm:"primaryKey"`
	Semester     string
	IsMandatory  bool

	Course     Course     `gorm:"foreignKey:CourseID"`
	Department Department `gorm:"foreignKey:DepartmentID"`
}
