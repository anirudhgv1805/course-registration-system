import { useEffect, useState } from "react";
import type { Department } from "../model/course";
import type { StudentFormData } from "../model/user";
import {
    getListOfDepartments,
    registerStudent,
} from "../services/auth.service";

export const StudentRegistration: React.FC = () => {
    const [formData, setFormData] = useState<StudentFormData>({
        username: "",
        password: "",
        registerno: "",
        email: "",
        section: "",
        batch: new Date().getFullYear(),
        departmentId: -1,
    });
    const [departments, setDepartments] = useState<Department[]>([]);

    useEffect(() => {
        const fetchDepartments = async () => {
            const response = await getListOfDepartments();
            setDepartments(response);
        };
        fetchDepartments();
    }, []);

    const handleChange = (e: any) => {
        const { name, value, type, checked } = e.target;
        setFormData((prev) => ({
            ...prev,
            [name]: type === "checkbox" ? checked : value,
        }));
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            formData.registerno = formData.registerno.toLowerCase();
            registerStudent(formData);
        } catch (error) {
            alert("Something went wrong.");
            console.error("Error:", error);
        }
    };

    return (
        <div className="min-h-screen bg-gray-100 flex items-center justify-center px-4">
            <div className="bg-white shadow-md rounded-lg w-full max-w-xl p-8">
                <h2 className="text-2xl font-semibold mb-6 text-center text-gray-800">
                    Student Registration
                </h2>
                <form onSubmit={handleSubmit} className="space-y-4">
                    <div>
                        <label
                            className="block text-sm font-medium text-gray-700"
                            htmlFor="name"
                        >
                            Username
                        </label>
                        <input
                            type="text"
                            name="username"
                            value={formData.username}
                            onChange={handleChange}
                            required
                            className="mt-1 w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring focus:ring-blue-200"
                            autoComplete="name"
                            id="name"
                        />
                    </div>

                    <div>
                        <label
                            className="block text-sm font-medium text-gray-700"
                            htmlFor="registerno"
                        >
                            Register No
                        </label>
                        <input
                            type="text"
                            name="registerno"
                            value={formData.registerno}
                            onChange={handleChange}
                            required
                            className="mt-1 w-full border border-gray-300 rounded px-3 py-2"
                            id="registerno"
                        />
                    </div>

                    <div>
                        <label
                            className="block text-sm font-medium text-gray-700"
                            htmlFor="password"
                        >
                            Password
                        </label>
                        <input
                            type="password"
                            name="password"
                            value={formData.password}
                            onChange={handleChange}
                            required
                            className="mt-1 w-full border border-gray-300 rounded px-3 py-2"
                            id="password"
                        />
                    </div>

                    <div className="grid grid-cols-2 gap-4">
                        <div>
                            <label
                                className="block text-sm font-medium text-gray-700"
                                htmlFor="department"
                            >
                                Department
                            </label>
                            <select
                                name="departmentId"
                                value={formData.departmentId}
                                onChange={handleChange}
                                required
                                disabled={departments.length === 0}
                                className="mt-1 w-full border border-gray-300 rounded px-3 py-2"
                                id="department"
                            >
                                <option value={-1} disabled>
                                    {departments.length > 0
                                        ? "Select a department"
                                        : "No Departments Available"}
                                </option>

                                {departments.map((dept) => (
                                    <option key={dept.id} value={dept.id}>
                                        {dept.name}
                                    </option>
                                ))}
                            </select>
                        </div>

                        <div>
                            <label
                                className="block text-sm font-medium text-gray-700"
                                htmlFor="batch"
                            >
                                Batch
                            </label>
                            <input
                                type="number"
                                name="batch"
                                value={formData.batch}
                                onChange={handleChange}
                                required
                                className="mt-1 w-full border border-gray-300 rounded px-3 py-2"
                                id="batch"
                            />
                        </div>
                    </div>

                    <div>
                        <label
                            className="block text-sm font-medium text-gray-700"
                            htmlFor="email"
                        >
                            Email
                        </label>
                        <input
                            type="email"
                            name="email"
                            value={formData.email}
                            onChange={handleChange}
                            required
                            className="mt-1 w-full border border-gray-300 rounded px-3 py-2"
                            autoComplete="email"
                            id="email"
                        />
                    </div>

                    <div>
                        <label
                            className="block text-sm font-medium text-gray-700"
                            htmlFor="section"
                        >
                            Section
                        </label>
                        <input
                            type="text"
                            name="section"
                            value={formData.section}
                            onChange={handleChange}
                            required
                            className="mt-1 w-full border border-gray-300 rounded px-3 py-2"
                            id="section"
                        />
                    </div>

                    <button
                        type="submit"
                        className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition"
                    >
                        Register
                    </button>
                </form>
            </div>
        </div>
    );
};
