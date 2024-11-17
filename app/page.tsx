import { Endpoint } from "@/components/endpoint";
import { ExperiencingIssues } from "@/components/experiencing-issues";
import { Operational } from "@/components/operational";
import { Endpoints } from "@/lib/types";
import Image from "next/image";
import Link from "next/link";

const BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export const revalidate = 300;

export default async function Home() {
  const response = await fetch(`${BASE_URL}/domain/history?domain=plug`);
  const data: Endpoints = await response.json();

  const operational = data.endpoints.every(
    (endpoint) =>
      endpoint.history[endpoint.history.length - 1].status ===
      endpoint.history[endpoint.history.length - 1].expected
  );

  return (
    <>
      <main>
        <div className="flex flex-row gap-2 p-8 max-w-[720px] items-center mx-auto justify-between">
          <Link className="flex flex-row items-center gap-2" href="https://onplug.io" target="_blank">          
            <Image className="w-8 h-8" src="/plug-logo-green.svg" alt="Logo" height={32} width={32} />
            <Image className="h-8" src="/plug-word-green.svg" alt="Wordmark" height={32} width={96} />
          </Link>

          <Link
            className="bg-plug-yellow text-plug-green font-bold px-3 py-1 rounded-md"
            href="https://twitter.com/onplug_io" target="_blank">
            Follow
          </Link>
        </div>

        <div className="h-52 flex items-center justify-center flex-col gap-2">
          {operational ? (
            <Operational className="w-24 h-24" />
          ) : (
            <ExperiencingIssues className="w-24 h-24" />
          )}
          <p className="text-2xl font-black">
            {operational ? "Operational" : "Experiencing Issues"}
          </p>
        </div>

        <div className="flex flex-col gap-2 p-8 max-w-[720px] items-center mx-auto">
          {data.endpoints.map((endpoint, endpointIndex) => (
            <Endpoint key={endpointIndex} endpoint={endpoint} />
          ))}

          <div className="bg-plug-green/10 h-[2px] w-full my-2" />

          <p className="text-sm font-bold">
            <span className="opacity-60">
              If you believe the system is down, but it is not reflected here
              please notify us by sending a DM on Twitter to
            </span>{" "}
            <Link
              href="https://twitter.com/onplug_io"
              target="_blank"
              className="opacity-100"
            >
              @onplug_io
            </Link>
            <span className="opacity-60">.</span>
          </p>

          <p className="text-sm font-bold opacity-60">
            We monitor the status of all our systems very closely. We check the
            status of all systems every 30 minutes. This page utilizes a global
            cache that clears every 5 minutes.
          </p>
        </div>
      </main>
    </>
  );
}
