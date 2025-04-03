import axios from "axios";

const API_BASE_URL = "http://localhost:8080";

export interface ShortenResponse {
  short_url: string;
  original_url: string;
}

export interface UrlListItem {
  short_code: string;
  original_url: string;
}

export async function shortenUrl(url: string): Promise<ShortenResponse> {
  const response = await axios.post(`${API_BASE_URL}/api/shorten`, { url });
  return response.data;
}

export async function getAllUrls(): Promise<UrlListItem[]> {
  const response = await axios.get(`${API_BASE_URL}/urls`);
  return response.data;
}