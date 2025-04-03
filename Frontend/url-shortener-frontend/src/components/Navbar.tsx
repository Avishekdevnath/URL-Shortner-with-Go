"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

export default function Navbar() {
  const pathname = usePathname();

  const navLinks = [
    { name: "Home", href: "/" },
    { name: "History", href: "/urls" },
    { name: "Features", href: "/features" },
    { name: "About", href: "/about" },
  ];

  return (
    <nav className="sticky top-0 z-50 bg-blue-600 text-white shadow-md">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          <div className="flex-shrink-0">
            <Link href="/" className="text-2xl font-bold tracking-tight">
              URL Shortener
            </Link>
          </div>
          <div className="hidden sm:flex sm:space-x-8">
            {navLinks.map((link) => {
              const isActive = pathname === link.href;
              return (
                <Link
                  key={link.name}
                  href={link.href}
                  className={`text-sm font-medium transition-colors duration-200 ${
                    isActive
                      ? "text-white border-b-2 border-white"
                      : "text-blue-100 hover:text-white hover:border-b-2 hover:border-blue-200"
                  }`}
                >
                  {link.name}
                </Link>
              );
            })}
          </div>
          <div className="hidden sm:block">
            <button className="text-sm font-medium text-blue-100 hover:text-white bg-blue-700 px-4 py-2 rounded-md hover:bg-blue-800 transition-colors duration-200">
              Login
            </button>
          </div>
          <div className="sm:hidden">
            <button className="text-blue-100 hover:text-white focus:outline-none">
              <svg
                className="h-6 w-6"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M4 6h16M4 12h16m-7 6h7"
                />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </nav>
  );
}