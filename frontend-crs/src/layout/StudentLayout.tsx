import { Outlet } from "react-router-dom";
import { StudentPrivateRoute } from "../components/PrivateRoute";

export const StudentLayout: React.FC = () => {
    return <StudentPrivateRoute children={<Outlet />} />;
};
