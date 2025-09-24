import type React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Header } from "./components/Header";
import { StaffLayout } from "./layout/StaffLayout";
import { StudentLayout } from "./layout/StudentLayout";
import { CourseRegistration } from "./pages/CourseRegistration";
import Login from "./pages/Login";
import { StaffDashboard } from "./pages/StaffDashboard";
import { StudentDashboard } from "./pages/StudentDashboard";
import { StaffRegistration } from "./pages/StaffRegistration";

const App: React.FC = () => {
    return (
        <BrowserRouter>
            <Header />
            <Routes>
                <Route path="/" element={<Login />} />
                <Route path="/staff-register" element={<StaffRegistration />} />
                {/* <Route
                    path="/student-register"
                    element={<StudentRegistration />}
                /> */}
                <Route path="/student" element={<StudentLayout />}>
                    <Route path="dashboard" element={<StudentDashboard />} />
                    <Route
                        path="course-registration"
                        element={<CourseRegistration />}
                    />
                </Route>
                <Route path="/staff" element={<StaffLayout />}>
                    <Route path="dashboard" element={<StaffDashboard />} />
                </Route>
            </Routes>
        </BrowserRouter>
    );
};

export default App;
