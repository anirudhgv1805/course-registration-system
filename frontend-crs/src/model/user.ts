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
