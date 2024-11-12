import { motion, useScroll, useTransform } from "framer-motion";
import { useRef } from "react";
import { InfoCard } from "./cards";
import { CalendarClock } from "lucide-react";
import { ActionStaking } from "./actions";

export const Curve3D = () => {
	const containerRef = useRef<HTMLDivElement>(null);
	const { scrollYProgress } = useScroll({
		target: containerRef,
		offset: ["start end", "end start"],
	});

	const pathLength = useTransform(scrollYProgress, [0, 0.35], [0, 1]);

	const diagonal = Math.sqrt(Math.pow(2800, 2) + Math.pow(1400, 2));
	const maxRadius = diagonal / 2;

	const circleRadius = useTransform(
		scrollYProgress,
		[0.34, 0.35, 0.6],
		[0, 30, maxRadius]
	);

	const textOpacity = useTransform(
		scrollYProgress,
		[0.5, 0.55],
		[0, 1]
	);

	return (
		<div className="h-screen w-full relative" ref={containerRef}>
			<div className="absolute inset-0 overflow-visible z-[99999]">
				<svg
					viewBox="0 0 1827 976"
					fill="none"
					xmlns="http://www.w3.org/2000/svg"
					className="w-full h-full overflow-visible"
				>
					<motion.path
						style={{ pathLength }}
						d="M0 251C234 385 219.145 251 508.814 251C732.341 251 718.957 473.761 913 486.5"
						stroke="url(#paint0_linear_4614_166)"
						strokeWidth="60"
						strokeLinecap="round"
					/>
					<motion.circle
						cx="913"
						cy="486.5"
						fill="#D2F38A"
						r={circleRadius}
						className="overflow-visible"
					/>
					<defs>
						<linearGradient
							id="paint0_linear_4614_166"
							x1="774.577"
							y1="456.812"
							x2="61.8676"
							y2="240.295"
							gradientUnits="userSpaceOnUse"
						>
							<stop stopColor="#D2F38A" />
							<stop offset="1" stopColor="#385842" />
						</linearGradient>
					</defs>
				</svg>
			</div>

			<motion.div
				className="absolute inset-0 flex items-center justify-center z-[100000] pt-[60%] flex flex-col gap-12"
				style={{
					opacity: textOpacity,
				}}
			>
				<div className="text-center max-w-[720px]">
					<h2 className="text-[52px] font-black mb-4 text-[#385842]">Every common crypto usecase at your fingertips.</h2>
					<p className="text-xl font-bold text-plug-green/40">Scroll down to explore more</p>
				</div>

				<div className="grid grid-cols-4 gap-4 max-w-[1200px]">
					<InfoCard
						icon={<CalendarClock size={24} className="opacity-40" />}
						text="Staking."
						description="Define timeframes for your transactions."
						className="relative z-[99999] h-[320px] sm:h-[320px] 2xl:h-[300px]"
					>
						<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
						<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
					</InfoCard>
					<InfoCard
						icon={<CalendarClock size={24} className="opacity-40" />}
						text="Borrow & Lend."
						description="Boost your crypto earnings by lending and borrowing against your collateral."
						className="relative z-[99999] col-span-2 h-[320px] sm:h-[320px] 2xl:h-[300px]"
					>

						<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
						<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
					</InfoCard>
					<ActionStaking />
					<InfoCard
						icon={<CalendarClock size={24} className="opacity-40" />}
						text="Discover Opportunities."
						description="Deposit your crypto into liquidity pools to earn swap fees and yield."
						className="relative z-[99999] col-span-2 h-[320px] sm:h-[320px] 2xl:h-[300px]"
					>

						<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
						<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
					</InfoCard>
					<InfoCard
						icon={<CalendarClock size={24} className="opacity-40" />}
						text="Bridge."
						description="Move your crypto quickly between chains."
						className="relative z-[99999] h-[320px] sm:h-[320px] 2xl:h-[300px]"
					>

						<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
						<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
					</InfoCard>
					<InfoCard
						icon={<CalendarClock size={24} className="opacity-40" />}
						text="Bridge."
						description="Move your crypto quickly between chains."
						className="relative z-[99999] h-[320px] sm:h-[320px] 2xl:h-[300px]"
					>

						<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
						<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
					</InfoCard>
					<InfoCard
						icon={<CalendarClock size={24} className="opacity-40" />}
						text="Provide Liquidity."
						description="Deposit your crypto into liquidity pools to earn swap fees and yield."
						className="relative z-[99999] col-span-2 h-[320px] sm:h-[320px] 2xl:h-[300px]"
					>
						<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
						<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
					</InfoCard>
					<InfoCard
						icon={<CalendarClock size={24} className="opacity-40" />}
						text="Bridge."
						description="Move your crypto quickly between chains."
						className="relative z-[99999] h-[320px] sm:h-[320px] 2xl:h-[300px]"
					>

						<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
						<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
					</InfoCard>
				</div>
			</motion.div>
		</div>
	);
};

export default Curve3D;
