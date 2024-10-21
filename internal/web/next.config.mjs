/** @type {import('next').NextConfig} */
const nextConfig = {
    output: 'standalone',
    images: {
        loader: 'custom',
        loaderFile: './src/image/loader.ts',
    },
};

export default nextConfig;
