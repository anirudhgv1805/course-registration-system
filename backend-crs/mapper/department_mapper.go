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
