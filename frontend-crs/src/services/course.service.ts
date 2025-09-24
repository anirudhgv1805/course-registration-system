import { axiosInstance } from "../utils/axiosInstance";

const getAllCoursesForStudent = async () => {
    try {
        const response = await axiosInstance.get(`/students/courses/`);
        return response;
    } catch (err: any) {
        throw new Error(
            err.response?.data?.message || "error in fetching courses"
        );
    }
};

export { getAllCoursesForStudent };
