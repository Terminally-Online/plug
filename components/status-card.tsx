import { motion } from "framer-motion"
import { CheckCircle } from "lucide-react"

interface StatusCardProps {
  title: string
  tag?: string
  description?: string
  status: "operational" | "degraded" | "down"
}

export const StatusCard = ({ title, tag = "", description, status = "operational" }: StatusCardProps) => {
  return (
    <motion.div 
      className="relative flex flex-col gap-2 rounded-lg border border-grayscale-100 bg-white p-6"
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.3 }}
    >
      <div className="absolute right-4 top-4">
        <CheckCircle 
          size={20} 
          className={status === "operational" ? "text-plug-yellow" : "text-red-500"}
        />
      </div>

      <div className="flex flex-col gap-1">
        <div className="flex items-center gap-2">
          <h2 className="text-2xl font-bold">{title}</h2>
          {tag && (
            <span className="rounded-md bg-plug-yellow/10 px-2 py-1 text-xs font-medium text-plug-green">
              {tag}
            </span>
          )}
        </div>
        {description && (
          <p className="text-sm text-gray-500">{description}</p>
        )}
      </div>
    </motion.div>
  )
}