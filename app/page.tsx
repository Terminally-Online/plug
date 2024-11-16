import { motion } from "framer-motion"
import { StatusCard } from "@/components/status-card"

export default function Home() {
  return (
    <main className="min-h-screen bg-[#FDFFF7] p-8">
      <div className="mx-auto max-w-6xl">
        <motion.div 
          className="mb-12"
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
        >
          <h1 className="text-4xl font-bold text-plug-green">System Status</h1>
          <p className="mt-2 text-gray-600">Check the current status of our services</p>
        </motion.div>

        <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
          <StatusCard 
            title="Website" 
            tag="Up!"
            status="operational"
          />
          <StatusCard 
            title="Docs" 
            tag="Up!"
            status="operational"
          />
          <StatusCard 
            title="WebApp" 
            tag="Up!"
            description="Error messages and more information about what's going on"
            status="operational"
          />
          <StatusCard 
            title="Solver" 
            tag="Learn More"
            status="operational"
          />
        </div>
      </div>
    </main>
  )
}