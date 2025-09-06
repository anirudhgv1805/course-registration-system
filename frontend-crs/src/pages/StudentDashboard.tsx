import { useNavigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

export const StudentDashboard: React.FC = () => {
    const { user } = useAuth();
    const navigate = useNavigate();

    function handleViewCourses(): void {
        navigate("/student/course-registration")
    }

    return (
        <div className="pt-20 px-8 bg-[#262626] min-h-screen text-white">
            <h1 className="text-3xl font-bold mb-2">Hello 👋</h1>
            <h2 className="text-2xl mb-6">Welcome, {user?.Username.toUpperCase()}</h2>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <DashboardCard title="📖 Available Courses">
                    <p className="text-lg">
                        View all open courses, including real-time seat availability and registration options.
                    </p>
                    <button onClick={() => handleViewCourses()} className="mt-4 px-6 py-3 bg-blue-600 hover:bg-blue-700 rounded text-lg">
                        View Courses
                    </button>
                </DashboardCard>

                <DashboardCard title="📚 My Registered Courses">
                    <p className="text-lg">
                        These are the courses you're currently enrolled in. You can drop or view details.
                    </p>
                    <button className="mt-4 px-6 py-3 bg-green-600 hover:bg-green-700 rounded text-lg">
                        My Courses
                    </button>
                </DashboardCard>

                <DashboardCard title="🔔 Notifications">
                    <ul className="list-disc list-inside text-lg space-y-2">
                        <li>Registration closes: Sep 10</li>
                        <li>Elective slots open for Semester 5</li>
                        <li>New course: AI & Ethics</li>
                    </ul>
                </DashboardCard>

                <DashboardCard title="📆 Academic Calendar">
                    <ul className="list-disc list-inside text-lg space-y-2">
                        <li>Classes begin: Sep 15</li>
                        <li>Mid-sem exams: Oct 20-25</li>
                        <li>Finals: Dec 10-20</li>
                    </ul>
                </DashboardCard>
            </div>
        </div>
    );
};

const DashboardCard: React.FC<{ title: string; children: React.ReactNode }> = ({ title, children }) => {
    return (
        <div className="bg-[#1f1f1f] rounded-lg shadow-lg p-6 min-h-[300px] flex flex-col justify-between hover:shadow-2xl transition duration-300">
            <h3 className="text-2xl font-semibold mb-4">{title}</h3>
            <div className="flex flex-col flex-grow justify-between">{children}</div>
        </div>
    );
};
