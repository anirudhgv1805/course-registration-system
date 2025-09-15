import { useState } from "react";
import { getAllCoursesForStudent } from "../services/course.service";

export const CourseRegistration: React.FC = () => {
    useState(() => {
        const fetchCourses = async () => {
            const response = await getAllCoursesForStudent();
            console.log(response.data);
        };
        fetchCourses();
    });

    return (
        <>
            <h1 className="flex items-center justify-center text-7xl text-white bg-black">
                Course registration page
            </h1>
        </>
    );
};
