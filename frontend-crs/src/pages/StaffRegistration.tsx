import { useEffect, useState } from "react";
import type { StaffFormData } from "../model/user";
import { getListOfDepartments, registerStaff } from "../services/auth.service";
import type { Department } from "../model/course";

export const StaffRegistration: React.FC = () => {
    const [formData, setFormData] = useState<StaffFormData>({
        username: "",
        staffId: "",
        password: "",
        departmentId: -1,
        batch: new Date().getFullYear(),
        email: "",
        isClassAdvisor: false,
        section: "",
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
            formData.staffId = formData.staffId.toLowerCase();
            registerStaff(formData);
        } catch (error) {
            alert("Something went wrong.");
            console.error("Error:", error);
        }
    };

    return (
        <div className="min-h-screen bg-gray-100 flex items-center justify-center px-4">
            <div className="bg-white shadow-md rounded-lg w-full max-w-xl p-8">
                <h2 className="text-2xl font-semibold mb-6 text-center text-gray-800">
                    Staff Registration
                </h2>
                <form onSubmit={handleSubmit} className="space-y-4">
                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Username
                        </label>
                        <input
                            type="text"
                            name="username"
                            value={formData.username}
                            onChange={handleChange}
                            required
                            className="mt-1 w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring focus:ring-blue-200"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Staff ID
                        </label>
                        <input
                            type="text"
                            name="staffId"
                            value={formData.staffId}
                            onChange={handleChange}
                            required
                            className="mt-1 w-full border border-gray-300 rounded px-3 py-2"
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Password
                        </label>
                        <input
                            type="password"
                            name="password"
                            value={formData.password}
                            onChange={handleChange}
                            required
                            className="mt-1 w-full border border-gray-300 rounded px-3 py-2"
                        />
                    </div>

                    <div className="grid grid-cols-2 gap-4">
                        <div>
                            <label className="block text-sm font-medium text-gray-700">
                                Department
                            </label>
                            <select
                                name="departmentId"
                                value={formData.departmentId}
                                onChange={handleChange}
                                required
                                disabled={departments.length === 0}
                                className="mt-1 w-full border border-gray-300 rounded px-3 py-2"
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
                            <label className="block text-sm font-medium text-gray-700">
                                Batch
                            </label>
                            <input
                                type="number"
                                name="batch"
                                value={formData.batch}
                                onChange={handleChange}
                                required
                                className="mt-1 w-full border border-gray-300 rounded px-3 py-2"
                            />
                        </div>
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Email
                        </label>
                        <input
                            type="email"
                            name="email"
                            value={formData.email}
                            onChange={handleChange}
                            required
                            className="mt-1 w-full border border-gray-300 rounded px-3 py-2"
                        />
                    </div>

                    <div className="flex items-center gap-3">
                        <input
                            type="checkbox"
                            name="isClassAdvisor"
                            checked={formData.isClassAdvisor}
                            onChange={handleChange}
                            className="h-4 w-4 text-blue-600 border-gray-300 rounded"
                        />
                        <label className="text-sm text-gray-700">
                            Is Class Advisor
                        </label>
                    </div>

                    <div>
                        <label className="block text-sm font-medium text-gray-700">
                            Section
                        </label>
                        <input
                            type="text"
                            name="section"
                            value={formData.section}
                            onChange={handleChange}
                            required
                            className="mt-1 w-full border border-gray-300 rounded px-3 py-2"
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
