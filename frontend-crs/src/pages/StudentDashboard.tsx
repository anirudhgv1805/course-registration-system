import { useAuth } from "../hooks/useAuth";

export const StudentDashboard : React.FC = () => {
    const { user } = useAuth();

    return (
        <>
            <div className="pt-16 bg-[#262626] h-screen">
                <h1>hello</h1>
                <h1>Welcome, {user?.Username}</h1>
            </div>
        </>
    );
};
