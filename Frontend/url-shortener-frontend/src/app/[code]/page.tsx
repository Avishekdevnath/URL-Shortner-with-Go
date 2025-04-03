import { redirect } from "next/navigation";

export default function ShortUrlPage({ params }: { params: { code: string } }) {
  redirect(`http://localhost:8080/${params.code}`);
}