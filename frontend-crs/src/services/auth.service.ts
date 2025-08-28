import type { LoginDTO } from "../model/auth.dto";
import { axiosInstance } from "../utils/axiosInstance";

const loginUser = async (payload: LoginDTO) => {
    console.log("Logging in");
    try {
        const role = extractRoleFromUserId(payload.userId);
        console.log("role is" + role);
        const response = await axiosInstance.post(
            `/${role}/auth/login`,
            payload
        );
        console.log(response);
        return response;
    } catch (err: any) {
        throw new Error(err.response?.data?.error || "Invalid Password");
    }
};

export { loginUser };

function extractRoleFromUserId(userId: string): "staff" | "student" {
    const match = userId.match(/\d+([a-zA-Z]+)\d+/);
    const role = match?.[1];

    if (!role) {
        throw new Error("Entered userId is wrong");
    }

    switch (role.length) {
        case 2:
            return "staff";
        case 3:
            return "student";
        default:
            throw new Error("Entered userId is wrong");
    }
}
