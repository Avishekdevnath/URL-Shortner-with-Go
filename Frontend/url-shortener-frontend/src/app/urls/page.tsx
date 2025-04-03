import { getAllUrls, UrlListItem } from "@/lib/api";

export default async function UrlsPage() {
  const urls: UrlListItem[] = await getAllUrls();

  return (
    <main className="max-w-2xl mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">All Shortened URLs</h1>
      {urls.length === 0 ? (
        <p className="text-gray-500">No URLs shortened yet.</p>
      ) : (
        <ul className="space-y-2">
          {urls.map((item) => (
            <li key={item.short_code} className="p-2 bg-white rounded shadow">
              <a
                href={`http://localhost:3000/${item.short_code}`}
                className="text-blue-500 underline"
              >
                {`http://localhost:3000/${item.short_code}`}
              </a>
              {" → "}
              <span>{item.original_url}</span>
            </li>
          ))}
        </ul>
      )}
      <a href="/" className="mt-4 inline-block text-blue-500 underline">
        Back to Home
      </a>
    </main>
  );
}