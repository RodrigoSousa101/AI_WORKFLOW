import { useEffect, useState } from 'react';
import '../App.css'
import { LuBrain, LuWorkflow } from "react-icons/lu";
import { IoPeopleOutline } from "react-icons/io5";
import { FiCheckCircle } from "react-icons/fi";
import axios from 'axios';
import { useNavigate } from 'react-router-dom';



const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [mensagem, setMensagem] = useState('');
  const [active, setactive] = useState('login');

  const navigate = useNavigate();  

   useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      navigate('/home');
    }
  }, [navigate]);

  
  const handleLogin = async () => {
    try {
      const response = await axios.post('http://localhost:8080/api/auth/login', {
        email: email,
        password: password,
      });

      console.log(response.data);
      setMensagem('Login bem-sucedido!');
      setError('');
      localStorage.setItem('token', response.data.token);
      navigate('/home');

      
    } catch (err) {
      console.error(err);
      if (axios.isAxiosError(err) && err.response) {
        setError(err.response.data?.error || 'Erro ao fazer login');
      } else {
        setError('Erro ao fazer login');
      }
      setMensagem('');
    }
  };

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
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Your@email.com"
            className=' mt-2 w-full h-10 rounded-lg border border-gray-400 p-4'
          />
        </div>
        <div className='mt-4'>
          <p>Password</p>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Your password"
            className=' mt-2 w-full h-10 rounded-lg border border-gray-400 p-4'
          />
        </div>
        {error && <p className="text-red-500 mt-2">{error}</p>}
        {mensagem && <p className="text-green-500 mt-2">{mensagem}</p>}
        <button
          className='w-full h-10 mt-6 bg-black text-white rounded-lg cursor-pointer'
          onClick={handleLogin}
        >
          Sign In
        </button>
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
};

export default Login;
