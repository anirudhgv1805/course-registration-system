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

  const axiosInstance = useAxiosInstance();

  useEffect(() => {
    if (jwtToken) {
      const getUserData = async () => {
        try {
          const response = await axiosInstance.post("/me");
          console.log(response?.data);
          setUser(response?.data);
          localStorage.setItem("user", JSON.stringify(user));
        } catch (err: unknown) {
          console.log("Error while hitting '/me' ", err);
        } finally {
        }
      };
      getUserData();
    }
  }, [jwtToken]);

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

  return (
    <AuthContext.Provider
      value={{
        isAuthenticated: !!jwtToken,
        jwtToken,
        user,
        login,
        logout,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};
