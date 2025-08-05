import { useEffect, useState } from 'react';
import { LuBrain } from "react-icons/lu";
import { useNavigate } from 'react-router-dom';
import Topbar from '../components/Topbar';
import { MdLogout } from "react-icons/md";
import api, { logout } from '../utils/api';
import WorkFlowCard from '../components/WorkFlowCard';

const Home = () => {
  const navigate = useNavigate();
  const [userName, setUserName] = useState('');
  const [userEmail, setUserEmail] = useState('');

  const handleLogout = () => {
    logout();
    navigate('/');
  };

  useEffect(() => {
    const fetchCurrentUser = async () => {
      try {
        const response = await api.get('/users/current');
        setUserName(response.data.user.name);
        setUserEmail(response.data.user.email);
      } catch (error) {
        console.error('Erro ao buscar usuário atual:', error);
        navigate('/'); // Redireciona para login se não autenticado
      }
    };

    fetchCurrentUser();
  }, [navigate]);

  return (
    <div>
      <Topbar
        children={
          <>
            <div className='gap-2 flex items-center'>
              <LuBrain size={35} className='text-indigo-600' />
              <p className='font-semibold text-2xl'>AI WorkFlow</p>
            </div>

            <div className='flex gap-6 items-center'>
              <div className='flex flex-col text-xs'>
                <p className='font-semibold'>{userName}</p>
                <p className='text-gray-700'>{userEmail}</p>
              </div>
              <button
                onClick={handleLogout}
                className='cursor-pointer hover:bg-gray-100 h-10 w-10 flex items-center justify-center rounded-lg'
              >
                <MdLogout size={20} />
              </button>
            </div>
          </>
        }
      />
      <div className='px-40 flex flex-col mt-8'>
        <p className='text-2xl font-bold'>My WorkFlows</p>
        <p className='text-gray-700 mb-6'>Workflows you created and manage</p>

        <WorkFlowCard />
      </div>
     
    </div>
  );
};

export default Home;
