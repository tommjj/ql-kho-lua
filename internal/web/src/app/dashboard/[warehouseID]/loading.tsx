import { Box, LoaderCircle } from 'lucide-react';

export const LoadingPage = () => {
    return (
        <div className="flex items-center justify-center top-0 left-0 w-full h-screen bg-main-bg-color dark:bg-main-bg-color-dark z-[999]">
            <div>
                <div className="mt-5">
                    <Box className="text-primary size-20" />
                </div>
                <div className="flex justify-center mt-6 ">
                    <div className="flex-grow-0 animate-spin rounded-full">
                        <LoaderCircle className="text-primary w-6 h-6 text-nav-bg-color-dark" />
                    </div>
                </div>
            </div>
        </div>
    );
};

export default LoadingPage;
