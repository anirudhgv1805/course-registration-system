import type React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Login from "./pages/Login";
import Register from "./pages/Register";
import {
    StaffPrivateRoute,
    StudentPrivateRoute,
} from "./components/PrivateRoute";
import { StudentDashboard } from "./pages/StudentDashboard";
import { StaffDashboard } from "./pages/StaffDashboard";

const App: React.FC = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Login />} />
                <Route path="/register" element={<Register />} />
                <Route
                    path="/student/dashboard"
                    element={
                        <StudentPrivateRoute children={<StudentDashboard />} />
                    }
                />
                <Route
                    path="/staff/dashboard"
                    element={
                        <StaffPrivateRoute children={ <StaffDashboard />} />
                    }
                />
            </Routes>
        </BrowserRouter>
    );
};

export default App;
