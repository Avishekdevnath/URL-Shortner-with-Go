import ShortenForm from "@/components/ShortenForm";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-4">
      <h1 className="text-4xl font-bold mb-8 text-gray-900">URL Shortener</h1>
      <div className="w-full max-w-5xl">
        <ShortenForm />
      </div>
      <a href="/urls" className="mt-6 text-blue-600 hover:underline text-lg font-medium">
        View All URLs
      </a>
    </main>
  );
}