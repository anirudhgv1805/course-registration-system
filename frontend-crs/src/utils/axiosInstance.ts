import axios from "axios";
import { useEffect } from "react";

export const axiosInstance = axios.create({
  baseURL: import.meta.env.VITE_BACKEND_API_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

export const useAxiosInstance = () => {
  useEffect(() => {
    const requestInterceptor = axiosInstance.interceptors.request.use(
      (config) => {
        const jwtToken = localStorage.getItem("jwtToken");
        if (jwtToken) {
          config.headers["Authorization"] = `Bearer ${jwtToken}`;
        }
        return config;
      },
      (error) => Promise.reject(error)
    );

    const responseInterceptor = axiosInstance.interceptors.response.use(
      (response) => response,
      (error) => {
        if (error.response?.status === 401) {
          console.log("Unauthorized access, redirecting to login...");
          window.location.href = "/";
        }
        if (error.response?.status === 500) {
          console.log("Token Expired Redirecting to Login page");
          localStorage.removeItem("jwtToken");
          window.location.href = "/";
        }
        return Promise.reject(error);
      }
    );

    return () => {
      axiosInstance.interceptors.request.eject(requestInterceptor);
      axiosInstance.interceptors.response.eject(responseInterceptor);
    };
  }, []);

  return axiosInstance;
};
