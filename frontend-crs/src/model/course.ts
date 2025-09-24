export interface CourseResponse {
    courses: Course[];
    status: string;
}

export interface Course {
    courseId: number;
    courseCode: string;
    name: string;
    totalCapacity: number;
    currentlyFilled: number;
    departmentId: number;
    department: Department;
    StaffID: number;
    handledByStaff: Staff;
    semester: number;
    credit: number;
    batch: number;
    applicableDepartmentIds: number[] | null;
}

export interface Department {
    id: number;
    name: string;
    code: string;
    block: string;
    ApplicableCourses: any | null;
}

export interface Staff {
    username: string;
    staffId: string;
    department: string;
    email: string;
    isClassAdvisor: boolean;
    section: string;
    batch: number;
    role: string;
}
