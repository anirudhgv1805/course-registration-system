import type { LoginDTO } from "../model/auth.dto";
import type { Department } from "../model/course";
import type { StaffFormData, StudentFormData } from "../model/user";
import { axiosInstance } from "../utils/axiosInstance";

const loginUser = async (payload: LoginDTO) => {
    try {
        const role = extractRoleFromUserId(payload.userId);
        const response = await axiosInstance.post(
            `/${role}/auth/login`,
            payload
        );
        return response;
    } catch (err: any) {
        throw new Error(
            err.response?.data?.error || "Invalid UserId or Password"
        );
    }
};

const registerStaff = async (payload: StaffFormData) => {
    try {
        const response = await axiosInstance.post(
            `/staff/auth/register`,
            payload
        );
        return response;
    } catch (err: any) {
        throw new Error(
            err.response?.data?.message || "unable to register staff"
        );
    }
};

const registerStudent = async (payload: StudentFormData) => {
    try {
        const response = await axiosInstance.post(
            `/student/auth/register`,
            payload
        );
        return response;
    } catch (err: any) {
        throw new Error(
            err.response?.data?.message || "unable to register student"
        );
    }
};

const getListOfDepartments = async (): Promise<Department[]> => {
    try {
        const response = await axiosInstance.get("/departments");
        return response.data?.departments;
    } catch (err: any) {
        throw new Error(
            err.response?.data?.message ||
                "cannot fetch the list of departments"
        );
    }
};

export { loginUser, getListOfDepartments, registerStaff, registerStudent };

// helpers

function extractRoleFromUserId(userId: string): "staff" | "student" {
    const match = userId.match(/\d+([a-zA-Z]{3})\d+/);
    const code = match?.[1];

    if (!code) {
        throw new Error("Entered userId is wrong");
    }

    const lastChar = code.charAt(2).toLowerCase();

    if (lastChar === "l" || lastChar === "r") {
        return "student";
    }

    return "staff";
}
