import { Outlet } from "react-router-dom";
import { StaffPrivateRoute } from "../components/PrivateRoute";

export const StaffLayout: React.FC = () => {
    return <StaffPrivateRoute children={<Outlet />} />;
};
