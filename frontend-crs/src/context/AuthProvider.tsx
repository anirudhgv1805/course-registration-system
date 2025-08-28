import { useEffect, useState, type ReactNode } from "react";
import { AuthContext } from "./AuthContext";
import { type User } from "../model/user";
import { useAxiosInstance } from "../utils/axiosInstance";

interface AuthProvividerProps {
    children: ReactNode;
}

export const AuthProvider: React.FC<AuthProvividerProps> = ({ children }) => {
    const [jwtToken, setJwtToken] = useState<string | null>(
        localStorage.getItem("jwtToken")
    );

    const [user, setUser] = useState<User | undefined>(undefined);
    const [error, setError] = useState<string | null>("");

    const login = (jwtToken: string) => {
        localStorage.setItem("jwtToken", jwtToken);
        setJwtToken(jwtToken);
    };

    const logout = () => {
        localStorage.removeItem("jwtToken");
        localStorage.removeItem("user");
        setJwtToken(null);
        setUser(undefined);
    };

    const setUserData = async (user : User) => {
        setUser(user);
        localStorage.setItem("user", JSON.stringify(user));
    };

    return (
        <AuthContext.Provider
            value={{
                isAuthenticated: !!jwtToken,
                jwtToken,
                user,
                login,
                logout,
                error,
                setError,
                setUserData,
            }}
        >
            {children}
        </AuthContext.Provider>
    );
};
