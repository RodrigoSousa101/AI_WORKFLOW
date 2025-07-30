
import { useState } from 'react';
import './App.css'
import { LuBrain, LuWorkflow } from "react-icons/lu";
import { IoPeopleOutline } from "react-icons/io5";
import { FiCheckCircle } from "react-icons/fi";



function App() {

  const[active, setactive] = useState('login');

  return (
    <div className="min-h-screen bg-gradient-to-r from-blue-50 to-indigo-100 flex flex-col justify-center items-center">
      <div className='flex items-center gap-3'>
        <LuBrain size={50} className='text-indigo-600'/>
        <p className='font-bold text-4xl'>AI WorkFlow</p>
      </div>
      <p className='pt-4 text-gray-700 text-lg'>Intelligent task management system</p>
      <div className='bg-white  w-110 mt-6 rounded-2xl shadow p-6'>
        <div className='grid grid-cols-2 items-center justify-items-center bg-gray-100 w-full h-12 rounded-lg p-1'>
          <div className='w-full'>
            <button onClick={() => setactive('login')} 
              className={`py-2 px-4 rounded w-full cursor-pointer ${
              active === 'login' ? 'bg-white text-black '  : 'bg-gray-100 text-gray-500'
              }`}>Sign In
            </button>
          </div>
          <div className='w-full'>
            <button onClick={() => setactive('register')} 
              className={`px-4 py-2 w-full rounded cursor-pointer ${
              active === 'register' ? 'bg-white text-black ' : 'bg-gray-100 text-gray-500'
              }`}>Sign Up
            </button>
          </div>
        </div>
        <p className='text-2xl font-bold mt-2'>Welcome back</p>
        <p className='text-gray-700'>Sign in to your account to continue</p>
        <div className='mt-4'>
          <p>Email</p>
          <input type="email" placeholder="Your@email.com" className=' mt-2 w-full h-10 rounded-lg border border-gray-400 p-4'></input>
        </div>
        <div className='mt-4'>
          <p>Password</p>
          <input type="password" placeholder="Your password" className=' mt-2 w-full h-10 rounded-lg border border-gray-400 p-4'></input>
        </div>
        <button className='w-full h-10 mt-6 bg-black text-white rounded-lg cursor-pointer'>Sign In</button>
      </div>
      <div className='flex items-center justify-center gap-20  mt-6'>
        <div className='flex flex-col items-center gap-2'> 
          <LuWorkflow size={35} className='text-indigo-600'></LuWorkflow>
          <p className='text-gray-700 text-sm'>AI Workflows</p>
        </div>
        <div className='flex flex-col items-center'>
          <IoPeopleOutline size={35} className='text-indigo-600'></IoPeopleOutline>
          <p className='text-gray-700 text-sm'>Collaborate</p>
        </div>
        <div className='flex flex-col items-center'>
           <FiCheckCircle size={35} className='text-indigo-600'></FiCheckCircle>
           <p className='text-gray-700 text-sm'>Automation</p>
        </div>
       
        </div>
    </div>
  )
}

export default App
