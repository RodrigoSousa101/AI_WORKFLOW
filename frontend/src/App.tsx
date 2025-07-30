
import { useState } from 'react';
import './App.css'
import { LuBrain } from "react-icons/lu";

function App() {

  const[active, setactive] = useState("login");

  return (
    <div className="min-h-screen bg-gradient-to-r from-blue-50 to-indigo-100 flex flex-col justify-center items-center">
      <div className='flex items-center gap-3'>
        <LuBrain size={50} className='text-indigo-600'/>
        <p className='font-bold text-4xl'>AI WorkFlow</p>
      </div>
      <p className='pt-4 text-gray-700 text-lg'>Sistema inteligente de gestão de tarefas</p>
      <div className='bg-white h-100 w-120 mt-6 rounded-2xl p-6'>
        <div className='grid grid-cols-2 items-center justify-items-center bg-gray-100 h-10'>
          <div>
            <button>Login</button>
          </div>
          <div>
            <button>Registar</button>
          </div>
        </div>
        <p className='text-2xl font-bold mt-2'>Bem-vindo de volta</p>
        <p className='text-gray-700'>Entre na sua conta para continuar</p>
      </div>
    </div>
  )
}

export default App
