"use client";

import { useState } from "react";
import { shortenUrl, ShortenResponse } from "@/lib/api";

export default function ShortenForm() {
  const [url, setUrl] = useState("");
  const [result, setResult] = useState<ShortenResponse | null>(null);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setResult(null);

    try {
      const data = await shortenUrl(url);
      setResult(data);
      setUrl("");
    } catch (err) {
      setError("Failed to shorten URL. Please try again.");
    }
  };

  return (
    <div className="w-full max-w-4xl mx-auto p-6 bg-white rounded-lg shadow-md border border-gray-100">
      <form onSubmit={handleSubmit} className="space-y-6">
        <div>
          <label
            htmlFor="url"
            className="block text-2xl font-semibold text-gray-800"
          >
            Create a Short URL
          </label>
          <input
            type="url"
            id="url"
            value={url}
            onChange={(e) => setUrl(e.target.value)}
            placeholder="Enter your URL (e.g., https://example.com)"
            required
            className="mt-4 block w-full p-4 text-lg border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all"
          />
        </div>
        <button
          type="submit"
          className="w-full bg-blue-600 text-white text-lg font-medium p-4 rounded-md hover:bg-blue-700 active:bg-blue-800 transition-all shadow-sm"
        >
          Shorten URL
        </button>
      </form>
      {result && (
        <div className="mt-6 p-4 bg-green-50 border border-green-200 rounded-md text-green-800 text-base font-medium">
          Shortened URL:{" "}
          <a
            href={`http://localhost:3000/${result.short_url}`}
            className="underline hover:text-green-600"
          >
            {`http://localhost:3000/${result.short_url}`}
          </a>
        </div>
      )}
      {error && (
        <div className="mt-6 p-4 bg-red-50 border border-red-200 rounded-md text-red-800 text-base font-medium">
          {error}
        </div>
      )}
    </div>
  );
}