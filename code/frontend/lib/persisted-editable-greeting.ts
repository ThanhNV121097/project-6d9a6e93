export type Greeting = {
  text: string;
};

const browserApiBase = process.env.NEXT_PUBLIC_API_URL ?? "/api";
const serverApiBase = process.env.API_ORIGIN ?? "http://backend:8080";

async function parseGreeting(response: Response): Promise<Greeting> {
  if (!response.ok) {
    throw new Error("Greeting request failed.");
  }
  return response.json();
}

export async function getGreeting(): Promise<Greeting> {
  return parseGreeting(
    await fetch(`${serverApiBase}/v1/greeting`, { cache: "no-store" }),
  );
}

export async function saveGreeting(text: string): Promise<Greeting> {
  return parseGreeting(
    await fetch(`${browserApiBase}/v1/greeting`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ text }),
    }),
  );
}
