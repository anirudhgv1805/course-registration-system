import { useEffect, useState } from "react";
import { getAllCoursesForStudent } from "../services/course.service";
import type { Course } from "../model/course";

interface CourseCardProps {
    course: Course;
}

export const CourseRegistration: React.FC = () => {
    const [courses, setCourses] = useState<Course[]>([]);

    useEffect(() => {
        fetchCourses();
    }, []);

    const fetchCourses = async () => {
        const response = await getAllCoursesForStudent();
        setCourses(response.data.courses);
    };

    return (
        <div className="bg-black pt-16 min-h-screen">
            <div className="grid grid-cols-2 px-8">
                {courses.map((course) => (
                    <CourseCard key={course.courseId} course={course} />
                ))}
            </div>
        </div>
    );
};

const CourseCard: React.FC<CourseCardProps> = ({ course }) => {
    const filledPercentage =
        (course.currentlyFilled / course.totalCapacity) * 100;
    return (
        <div className="w-full bg-white border-b border-gray-200 hover:bg-gray-50 transition-all px-4 py-3 select-none">
            <div className="grid md:grid-cols-[1fr_auto] items-center gap-6">
                <div className="flex flex-col sm:flex-row sm:items-center sm:gap-6">
                    <div className="text-sm font-medium text-gray-600 w-24">
                        {course.courseCode}
                    </div>
                    <div className="text-gray-800 font-semibold">
                        {course.name}
                    </div>
                    <div className="flex flex-wrap gap-4 text-sm text-gray-600 mt-2 sm:mt-0">
                        <div>
                            <span className="font-medium">Staff:</span>{" "}
                            {course.handledByStaff.username}
                        </div>
                        <div>
                            <span className="font-medium">Dept:</span>{" "}
                            {course.department.code}
                        </div>
                        <div>
                            <span className="font-medium">Sem:</span>{" "}
                            {course.semester}
                        </div>
                        <div>
                            <span className="font-medium">Credit:</span>{" "}
                            {course.credit}
                        </div>
                        <div className="flex items-center gap-2">
                            <span className="font-medium">Seats:</span>
                            <span>
                                {course.currentlyFilled} /{" "}
                                {course.totalCapacity}
                            </span>
                            <div className="w-32 h-2 bg-gray-200 rounded-full overflow-hidden">
                                <div
                                    className="h-full bg-green-500"
                                    style={{
                                        width: `${filledPercentage}%`,
                                    }}
                                />
                            </div>
                        </div>
                    </div>
                </div>
                <div className="flex gap-2 text-sm text-gray-600">
                    <button className="font-medium px-3 py-1 rounded bg-gray-100 hover:bg-gray-200">
                        Block
                    </button>
                    <button className="font-medium px-3 py-1 rounded bg-blue-100 hover:bg-blue-200">
                        Register
                    </button>
                </div>
            </div>
        </div>
    );
};
