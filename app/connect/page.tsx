import { getServerSession } from 'next-auth'
import dynamic from 'next/dynamic'
import { redirect } from 'next/navigation'
import { authOptions } from '../api/auth/[...nextauth]/route'

const Button = dynamic(() => import('../../components/auth/button'), { ssr: false })

export default async function Page() { 
  const session = await getServerSession(authOptions)

  const username = session?.user?.name

  if(username) redirect('/canvas')

  return <>
    <h1>Log in to proceed</h1>

    <Button />
  </>
}
