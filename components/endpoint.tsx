"use client";

import { Endpoint as EndpointType } from "@/lib/types";
import { ChevronRight } from "lucide-react";
import { FC, useMemo, useState } from "react";
import { twMerge } from "tailwind-merge";
import { motion } from "framer-motion";
import Link from "next/link";
import { Operational } from "./operational";
import { ExperiencingIssues } from "./experiencing-issues";

export const Endpoint: FC<{ endpoint: EndpointType }> = ({ endpoint }) => {
  const [hoveredName, setHoveredName] = useState(false);
  const [hoveredIndex, setHoveredIndex] = useState<number | null>(null);

  const operational =
    endpoint.history[endpoint.history.length - 1].status ===
    endpoint.history[endpoint.history.length - 1].expected;

  const max = useMemo(
    () => Math.max(...endpoint.history.map((h) => h.duration)),
    [endpoint]
  );

  return (
    <div className="relative rounded-lg border-[1px] border-plug-green/10 p-4 gap-2 flex flex-col font-bold w-full overflow-hidden">
      <div className="absolute left-0 bottom-0 filter blur-[100px] -z-[1]">
        {operational ? (
          <Operational className="w-32 h-32 min-w-32" />
        ) : (
          <ExperiencingIssues className="w-32 h-32 min-w-32" />
        )}
      </div>

      <div className="flex flex-col relative">
        <div className="flex flex-row gap-4">
          {operational ? (
            <Operational className="w-12 h-12 min-w-12" />
          ) : (
            <ExperiencingIssues className="w-12 h-12 min-w-12" />
          )}

          <div className="flex flex-row justify-between items-center font-black w-full min-w-0">
            <div className="flex flex-col min-w-0 flex-1 mr-4 overflow-ellipsis">
              <Link
                href={endpoint.url}
                target="_blank"
                className="flex flex-row items-center cursor-pointer truncate max-w-full"
                onMouseEnter={() => setHoveredName(true)}
                onMouseLeave={() => setHoveredName(false)}
              >
                <motion.span
                  animate={{
                    marginRight: hoveredName ? "8px" : "4px",
                  }}
                >
                  {endpoint.url.replace("https://", "")}
                </motion.span>
                <motion.span animate={{ opacity: hoveredName ? 1 : 0.4 }}>
                  <ChevronRight size={18} className="mt-[2px]" />
                </motion.span>
              </Link>

              <p className="font-black text-sm opacity-60">
                {operational ? "Operational" : "Experiencing Issues"}
              </p>
            </div>

            <div className="flex flex-col text-right">
              <p className="opacity-60 text-sm">Uptime</p>
              <h3 className="opacity-60 tabular-nums text-sm">
                {endpoint.stats.uptime_percentage}%
              </h3>
            </div>
          </div>
        </div>
      </div>

      <div className="flex flex-row justify-between gap-[2px] h-16">
        {endpoint.history.map((history, historyIndex) => (
          <motion.div
            key={historyIndex}
            className={twMerge(
              "w-full h-12 rounded-lg cursor-pointer mt-auto",
              history.status === history.expected
                ? "bg-plug-yellow"
                : "bg-plug-red"
            )}
            initial={{ height: 0 }}
            animate={{
              height: `${(history.duration / max) * 100}%`,
            }}
            onMouseEnter={() => setHoveredIndex(historyIndex)}
            onMouseLeave={() => setHoveredIndex(null)}
          />
        ))}
      </div>

      <div className="flex flex-row justify-between text-sm">
        <p className="opacity-60 whitespace-nowrap truncate">
          {hoveredIndex !== null
            ? new Date(
                endpoint.history[hoveredIndex].timestamp
              ).toLocaleString()
            : new Date(endpoint.stats.last_check).toLocaleString()}
        </p>
        <p className="opacity-60">
          {hoveredIndex !== null
            ? `${(endpoint.history[hoveredIndex].duration / 1e9).toFixed(2)}s`
            : `${(endpoint.stats.average_response_ms / 1000).toFixed(2)}s`}
        </p>
      </div>
    </div>
  );
};
