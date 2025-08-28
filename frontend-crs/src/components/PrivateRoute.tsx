import { Navigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

interface Props {
    children: React.ReactNode;
}
const PrivateRoute: React.FC<Props> = ({ children }) => {
    const { isAuthenticated } = useAuth();

    if (!isAuthenticated) {
        return <Navigate to={"/"} replace />;
    }

    return <>{children}</>;
};

const StudentPrivateRoute: React.FC<Props> = ({ children }) => {
    const { isAuthenticated, user } = useAuth();

    if (!isAuthenticated || user?.role !== "student") {
        console.log("triggered private route student");
        return <Navigate to={"/"} replace />;
    }
    return <>{children}</>;
};

const StaffPrivateRoute: React.FC<Props> = ({ children }) => {
    const { isAuthenticated, user } = useAuth();

    if (!isAuthenticated || user?.role !== "staff") {
        return <Navigate to={"/"} replace />;
    }
    return <>{children}</>;
};

export { PrivateRoute, StudentPrivateRoute, StaffPrivateRoute };
