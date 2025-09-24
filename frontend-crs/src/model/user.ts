export interface User {
    Username: string;
    userId: string;
    Department: string;
    Email: string;
    Section: string | null;
    Batch: number;
    isClassAdvisor: boolean;
    role : string;
}

export interface StaffFormData {
    username: string;
    staffId: string;
    password: string;
    departmentId: number;
    batch: number;
    email: string;
    isClassAdvisor: boolean;
    section: string;
}