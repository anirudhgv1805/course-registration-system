export const Loading: React.FC = () => {
    return (
        <div className="absolute flex items-center bg-black z-10 opacity-50 justify-center w-full h-screen inset-0">
            <div className="">
                <h1 className="text-white font-semibold text-2xl animate-pulse">
                    Loading... Please wait!
                </h1>
            </div>
        </div>
    );
};
