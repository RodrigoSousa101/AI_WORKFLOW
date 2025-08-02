import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom';

const Home: React.FC = () => {
    const navigate = useNavigate();  
    useEffect(() => {
    const token = localStorage.getItem('token');
    if (!token) {
      navigate('/');
    }
  }, [navigate]);
  
    return(
        <div>a</div>
    )
    
}

export default Home; 