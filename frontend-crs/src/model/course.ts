import type { User } from "./user";

export interface Course {
    ID: number;
    courseCode: string;
    name: string;
    totalCapacity: number;
    currentlyFilled: number;
    department: string;
    handledByStaff: User;
    semester: string;
    credit: number;
    batch: number;
}
