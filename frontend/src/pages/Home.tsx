import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom';
import Topbar from '../components/Topbar'

const Home = () => {
    const navigate = useNavigate();  
    useEffect(() => {
    const token = localStorage.getItem('token');
    if (!token) {
      navigate('/');
    }
  }, [navigate]);
  
    return(
     
        <div>
          <Topbar 
            children={
              <div> ola </div>
            }
            
          />
        </div>
    )
    
}

export default Home; 