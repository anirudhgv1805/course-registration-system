import { useContext } from "react";
import { AuthContext, type AuthContextType } from "../context/authContext";

export const useAuth = () : AuthContextType => {
  const context = useContext(AuthContext);
  if (!context) throw new Error("useAuth is defined out of the AuthProvider");
  return context;
};
