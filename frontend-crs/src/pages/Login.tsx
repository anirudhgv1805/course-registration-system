import { useState } from "react";
import { useAxiosInstance } from "../utils/axiosInstance";
import { useAuth } from "../hooks/useAuth";
import { Link } from "react-router-dom";

const Login: React.FC = () => {
    const [formData, setFormData] = useState({
        registerno: "",
        password: "",
    });

    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState<boolean>(false);

    const axiosInstance = useAxiosInstance();
    const { login } = useAuth();

    const handleChange = (e: { target: { name: any; value: any } }) => {
        const { name, value } = e.target;

        setFormData((prev) => ({
            ...prev,
            [name]: value,
        }));

        setError("");
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        setLoading(true);

        const validatedData = {
            ...formData,
            registerno: formData.registerno.toLowerCase(),
        };

        try {
            const response = await axiosInstance.post(
                "/student/auth/login",
                validatedData
            );
            const jwtToken = response?.data?.token;
            login(jwtToken);
        } catch (err) {
            console.log("An error occurred while logging in : ", err);
        } finally {
            setLoading(false);
        }
    };
    return (
        <>
            <div className="relative flex items-center justify-center h-screen bg-[url('/loginbg.jpg')] bg-cover bg-center">
                <div className="shadow-lg rounded w-full max-w-sm p-8 select-none backdrop-blur-2xl bg-white/30">
                    <form
                        onSubmit={handleSubmit}
                        className="flex flex-col gap-4 items-center"
                    >
                        <h2 className="text-2xl font-medium mb-3 bg-clip-text">
                            Login to your account
                        </h2>

                        <input
                            className="border border-gray-300 px-4 py-2 rounded-sm w-full"
                            type="text"
                            placeholder="Register No or Staff Id"
                            onChange={handleChange}
                            name="registerno"
                        />
                        <input
                            className="border border-gray-300 px-4 py-2 rounded-sm w-full"
                            type="password"
                            placeholder="Password"
                            onChange={handleChange}
                            name="password"
                        />
                        <button
                            className="bg-blue-600 text-white px-4 py-2 rounded w-full hover:bg-blue-700 transition"
                            type="submit"
                        >
                            Login
                        </button>

                        <Link
                            to="/register"
                            className="underline text-blue-100 text-sm"
                        >
                            Don't have an account? Register here
                        </Link>
                    </form>
                </div>
                {loading ? (
                    <div className="absolute flex items-center bg-black z-10 opacity-50 justify-center w-full h-screen inset-0">
                        <div className="">
                            <h1 className="text-white font-semibold text-2xl animate-pulse">Loading...</h1>
                        </div>
                    </div>
                ) : (
                    <></>
                )}
            </div>
        </>
    );
};

export default Login;
