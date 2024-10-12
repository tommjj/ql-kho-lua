/** @type {import('next').NextConfig} */
const nextConfig = {
    images: {
        loader: 'custom',
        loaderFile: './src/image/loader.ts',
    },
};

export default nextConfig;
