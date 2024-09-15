import { FC, HTMLAttributes, PropsWithChildren } from "react"

export type ButtonProps = {
	callbackUrl?: string
	redirect?: boolean
}

export const AuthButton: FC<HTMLAttributes<HTMLButtonElement> & PropsWithChildren<ButtonProps>> = ({
	callbackUrl = "/app/",
	redirect = true,
	className
}) => {
	return <></>
}

export default AuthButton
