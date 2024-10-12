'use client';

export default function myImageLoader({ src }: { src: string }) {
    return `http://localhost:8080/static/${src}`;
}
