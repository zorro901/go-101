import { useMutateAuth } from '../hooks/useMutateAuth';
import { ShieldCheckIcon } from '@heroicons/react/20/solid';

const Todo = () => {
  const { logoutMutation } = useMutateAuth()
  const logout = async () => {
    await logoutMutation.mutateAsync()
  }
  return (
    <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
      <div className="flex items-center my-3">
        <ShieldCheckIcon className="h-8 w-8 mr-3 text-indigo-500 cursor-pointer"
        onClick={logout}/>
      </div>
    </div>
  )
}

export default Todo
