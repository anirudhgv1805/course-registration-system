import { createContext } from "react";
import type { User } from "../model/user";

export interface AuthContextType {

    jwtToken : string | null;
    user : User | undefined;
    isAuthenticated : boolean;
    login : (...arg0: any[]) => any;
    logout : () => any;
}


export const AuthContext = createContext< AuthContextType | undefined>(undefined);