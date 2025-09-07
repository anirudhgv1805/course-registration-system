import { useState } from "react";
import { useAuth } from "../hooks/useAuth";

export const Header: React.FC = () => {
    const { isAuthenticated, logout } = useAuth();
    const [showLogoutOverlay, setShowLogoutOverlay] = useState<boolean>(false);

    const handleLogout = () => {
        if (!isAuthenticated) return;
        setShowLogoutOverlay(true);

        setTimeout(() => {
            logout();
            setShowLogoutOverlay(false);
        }, 1500);
    };

    return isAuthenticated ? (
        <header className="flex p-3 w-full justify-around items-center font-thin fixed z-10 backdrop:blur-2xl md:text-2xl text-white select-none">
            <div className="flex relative ml-auto gap-8">
                <div>Home</div>
                <div>Course Available</div>
                <div>Registered Courses</div>
            </div>
            <div
                className="ml-auto mr-3 cursor-pointer px-0.5 py-0.5 rounded-lg shadow-lg bg-gradient-to-r from-[#0F2027] via-[#0d485b] to-[#262626]"
                onClick={() => handleLogout()}
            >
                Logout
            </div>
            <div
                className={`fixed inset-0 z-50 flex items-center justify-center transition-all duration-2000 text-white
                ${
                    showLogoutOverlay
                        ? "opacity-100 bg-black bg-opacity-70"
                        : "opacity-0 pointer-events-none bg-opacity-0"
                }
            `}
            >
                <h1 className="text-white font-semibold text-2xl animate-pulse">
                    Logging out...
                </h1>
            </div>
        </header>
    ) : (
        <></>
    );
};
