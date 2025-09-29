package mapper

import (
	"backend-crs/dto"
	"backend-crs/model"
)

func DepartmentToDepartmentResponseForListOfDepartments(department *model.Department, departmentResponse *dto.DepartmentResponse) {

	departmentResponse.ID = department.ID
	departmentResponse.Name = department.Name
	departmentResponse.Block = department.Block
	departmentResponse.Code = department.Code
}

func DepartmentToDepartmentResponse(department *model.Department, departmentResponse *dto.DepartmentResponse) {
	departmentResponse.ID = department.ID
	departmentResponse.Name = department.Name
	departmentResponse.Code = department.Code
	departmentResponse.Block = department.Block

	for _, courseApplicable := range department.ApplicableCourses {
		var courseApplicableDTO dto.CourseApplicableDTO
		courseApplicableDTO.CourseID = courseApplicable.CourseID
		courseApplicableDTO.DepartmentID = courseApplicable.DepartmentID
		departmentResponse.ApplicableCourses = append(departmentResponse.ApplicableCourses, courseApplicableDTO)
	}
}
