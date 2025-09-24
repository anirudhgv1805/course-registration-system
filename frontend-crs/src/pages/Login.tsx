import { useState } from "react";
import { useAuth } from "../hooks/useAuth";
import { Link, useNavigate } from "react-router-dom";
import { Loading } from "../components/Loading";
import { loginUser } from "../services/auth.service";

const Login: React.FC = () => {
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        userId: "",
        password: "",
    });

    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState<boolean>(false);
    const { login, setUserData } = useAuth();

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

        setError("");
        setLoading(true);

        const validatedData = {
            ...formData,
            userId: formData.userId.toLowerCase(),
        };

        try {
            const response = await loginUser(validatedData);
            login(response?.data?.token);
            const user = response.data.body;
            setUserData({
                Username: user.username,
                userId: user.registerno ? user.registerno : user.staffId,
                Department: user.department,
                Email: user.email,
                Section: user.section,
                Batch: user.batch,
                isClassAdvisor: user?.isClassAdvisor,
                role: user.role,
            });
            if (user?.role === "student") navigate("/student/dashboard");
            else navigate("/staff/dashboard");
        } catch (err: any) {
            setError(err?.message);
        } finally {
            setLoading(false);
        }
    };
    return (
        <>
            <div className="relative flex items-center justify-center h-screen bg-[url('/loginbg.jpg')] bg-cover bg-center scroll-pt-16">
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
                            name="userId"
                            autoComplete="username"
                        />
                        <input
                            className="border border-gray-300 px-4 py-2 rounded-sm w-full"
                            type="password"
                            placeholder="Password"
                            onChange={handleChange}
                            name="password"
                            autoComplete="current-password"
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
                        {error && <p>{error}</p>}
                    </form>
                </div>
                {loading && <Loading />}
            </div>
        </>
    );
};

export default Login;
