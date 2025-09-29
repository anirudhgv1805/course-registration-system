import { useEffect, useRef } from "react";
import { Link } from "react-router-dom";

interface SelectRegisterUserTypeProps {
    setSelectUserRegisterType: React.Dispatch<React.SetStateAction<boolean>>;
}

const SelectRegisterUserType: React.FC<SelectRegisterUserTypeProps> = ({
    setSelectUserRegisterType,
}) => {
    const popupRef = useRef<HTMLDivElement>(null);
    useEffect(() => {
        const handleOutsideClick = (event: MouseEvent) => {
            if (
                popupRef.current &&
                !popupRef.current.contains(event.target as Node)
            ) {
                setSelectUserRegisterType(() => false);
            }
        };
        document.addEventListener("mousedown", handleOutsideClick);

        return () => {
            document.removeEventListener("mousedown", handleOutsideClick);
        };
    }, [setSelectUserRegisterType]);
    return (
        <>
            <div className="absolute flex items-center bg-black z-10 opacity-80 justify-center w-full h-screen inset-0 gap-4 flex-col">
                <p className="text-3xl text-white pt-10 pb-10">Who are you?</p>
                <div ref={popupRef} className="relative flex flex-row gap-20">
                    <Link
                        className="text-xl text-white border-2 p-2 rounded-lg"
                        to={"student-register"}
                    >
                        Student
                    </Link>
                    <Link
                        to={"staff-register"}
                        className="text-xl text-white border-2 p-2 rounded-lg"
                    >
                        Staff
                    </Link>
                </div>
            </div>
        </>
    );
};


export default SelectRegisterUserType;