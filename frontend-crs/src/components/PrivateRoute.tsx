import { Navigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import { Loading } from "./Loading";

interface Props {
    children: React.ReactNode;
}
const PrivateRoute: React.FC<Props> = ({ children }) => {
    const { isAuthenticated, loading } = useAuth();

    if (loading) return <Loading />;

    if (!isAuthenticated) {
        return <Navigate to={"/"} replace />;
    }

    return <>{children}</>;
};

const StudentPrivateRoute: React.FC<Props> = ({ children }) => {
    const { isAuthenticated, user, loading } = useAuth();
    if (loading) return <Loading />;

    if (!isAuthenticated || user?.role !== "student") {
        return <Navigate to={"/"} replace />;
    }
    return <>{children}</>;
};

const StaffPrivateRoute: React.FC<Props> = ({ children }) => {
    const { isAuthenticated, user, loading } = useAuth();
    if (loading) return <Loading />;

    if (!isAuthenticated || user?.role !== "staff") {
        return <Navigate to={"/"} replace />;
    }
    return <>{children}</>;
};

export { PrivateRoute, StudentPrivateRoute, StaffPrivateRoute };
