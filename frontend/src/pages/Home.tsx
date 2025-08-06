import { useEffect, useState } from 'react';
import { LuBrain } from "react-icons/lu";
import { useNavigate } from 'react-router-dom';
import Topbar from '../components/Topbar';
import { MdLogout } from "react-icons/md";
import api, { logout } from '../utils/api';
import ManagerWorkFlowCard from '../components/Manager_WorkFlowCard';
import { Carousel } from 'primereact/carousel';

type Workflow = {
  id: string;
  name: string;
  description: string;
  status: string;
};
        

const Home = () => {
  const navigate = useNavigate();
  const [userName, setUserName] = useState('');
  const [userEmail, setUserEmail] = useState('');
  const [userID, setUserID] = useState('');
  const [workflowsCreated, setWorkflowsCreated] = useState<Workflow[]>([]);
  const [workflowsWorker, setWorkflowsWorker] = useState<Workflow[]>([]);


  const handleLogout = () => {
    logout();
    navigate('/');
  };

  useEffect(() => {
  const fetchCurrentUserAndWorkflows = async () => {
    try {
      const response = await api.get('/users/current');
      const user = response.data.user;

      setUserName(user.name);
      setUserEmail(user.email);
      setUserID(user.id);

      const wfResponse = await api.get(`/workflowuser/user/${user.id}`);
      setWorkflowsCreated(wfResponse.data.WorkflowCreated);
    setWorkflowsWorker(wfResponse.data["Workflows worker"]);
      console.log('Workflows:', wfResponse.data);
    } catch (error) {
      console.error('Erro ao buscar dados:', error);
      navigate('/');
    }
  };

  fetchCurrentUserAndWorkflows();
}, [navigate]);

const workflowTemplate = (workflow: Workflow) => (
 
    <ManagerWorkFlowCard
      key={workflow.id}
      name={workflow.name}
      description={workflow.description}
      status={workflow.status}
    />
);


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
      <div className='px-40 flex flex-col mt-8 '>
        <p className='text-2xl font-bold'>My WorkFlows</p>
        <p className='text-gray-700 mb-6'>Workflows you created and manage</p>
        <Carousel
          value={workflowsCreated}
          itemTemplate={workflowTemplate} 
          numVisible={3}       
          numScroll={1}        
          className="custom-carousel"
        />
      </div>
     
    </div>
  );
};

export default Home;
