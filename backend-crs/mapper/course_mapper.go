package mapper

import (
	"backend-crs/dto"
	"backend-crs/model"
)

func CourseToCourseResponseDTO(course *model.Course, courseDTO *dto.CourseResponse) {
	courseDTO.CourseID = course.ID
	courseDTO.CourseCode = course.CourseCode
	courseDTO.Name = course.Name
	courseDTO.TotalCapacity = course.TotalCapacity
	courseDTO.CurrentlyFilled = course.CurrentlyFilled
	courseDTO.DepartmentID = course.DepartmentID

	// Mapping the offered by department to DepartmentDTO
	courseDTO.OfferedBy.Name = course.OfferedBy.Name
	courseDTO.OfferedBy.Code = course.OfferedBy.Code
	courseDTO.OfferedBy.Block = course.OfferedBy.Block

	courseDTO.StaffID = course.StaffID

	// Mapping the handledByStaff to StaffResponse
	courseDTO.HandledByStaff.Username = course.HandledByStaff.Username
	courseDTO.HandledByStaff.StaffId = course.HandledByStaff.StaffId

	// Department to DepartmentResponse
	courseDTO.HandledByStaff.Department.ID = course.HandledByStaff.Department.ID
	courseDTO.HandledByStaff.Department.Name = course.HandledByStaff.Department.Name
	courseDTO.HandledByStaff.Department.Block = course.HandledByStaff.Department.Block
	courseDTO.HandledByStaff.Department.Code = course.HandledByStaff.Department.Code

	// CourseApplicable to CourseApplicableDTO
	// courseDTO.HandledByStaff.Department.ApplicableCourses = course.HandledByStaff.Department.ApplicableCourses

	courseDTO.HandledByStaff.Email = course.HandledByStaff.Email
	courseDTO.HandledByStaff.IsClassAdvisor = course.HandledByStaff.IsClassAdvisor
	courseDTO.HandledByStaff.Section = course.HandledByStaff.Section
	courseDTO.HandledByStaff.Batch = course.HandledByStaff.Batch

	courseDTO.Semester = course.Semester
	courseDTO.Credit = course.Credit
	courseDTO.Batch = course.Batch
}
