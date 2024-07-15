import { fileURLToPath } from 'url';
import { dirname, join } from 'path';

/** @type {import('next').NextConfig} */
const nextConfig = {
    output: "export",
    distDir: "build",
    reactStrictMode: true,
    cleanDistDir: true,
    images: {
        unoptimized: true,
    },
    sassOptions: {
        includePaths: [join(dirname(fileURLToPath(import.meta.url)), "styles")]
    },
};

export default nextConfig;
