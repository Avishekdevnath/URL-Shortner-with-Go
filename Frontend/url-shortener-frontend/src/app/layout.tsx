// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Frontend\url-shortener-frontend\src\app\layout.tsx
import type { Metadata } from "next";
import "@/app/globals.css";
import Navbar from "@/components/Navbar";

export const metadata: Metadata = {
    title: "URL Shortener",
    description: "Shorten your URLs easily",
};

export default function RootLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <html lang="en">
            <body className="min-h-screen bg-gray-100">
                <Navbar />
                {children}
            </body>
        </html>
    );
}