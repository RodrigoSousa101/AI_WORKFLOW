import { RiVipCrownLine } from "react-icons/ri";
import { GoClock } from "react-icons/go";
import { FiCheckCircle } from "react-icons/fi";

type Props = {
  name : string;
  description: string;
  status: string;
};

const ManagerWorkFlowCard = ({ status, name, description }: Props) => {
  const renderStatus = (status: string) => {
    switch (status) {
      case "In Progress":
        return (
          <div className="w-24 h-5 bg-yellow-200 text-yellow-800 text-xs rounded-2xl flex items-center justify-center">
            <GoClock size={15} className="mr-1" />
            <span>In Progress</span>
          </div>
        );
      case "Completed":
        return (
          <div className="w-24 h-5 bg-green-200 text-green-800 text-xs rounded-2xl flex items-center justify-center">
            <FiCheckCircle size={15} className="mr-1" />
            <span>Completed</span>
          </div>
        );
      default:
        return null;
    }
  };

  return (
    <div className="border-l-4 border-indigo-600 w-100 h-70 rounded-xl shadow p-4">
    <div className="flex gap-4 items-center justify-between min-w-0">
        <p className="font-semibold text-xl flex-1 line-clamp-2">{name}</p>

        <div className="w-20 h-5 bg-indigo-200 text-indigo-800 font-semibold text-xs rounded-2xl flex items-center justify-center">
          <RiVipCrownLine size={16} className="mr-1" />
          <p>Manager</p>
        </div>

        {renderStatus(status)}
      </div>

      <p className="flex-1 max-w-70 break-words text-sm text-gray-700 mt-2 line-clamp-3">{description}</p>  
      <div className="flex justify-between mt-4">
        <p>Progress</p>
        <p>8/12 tasks</p>
      </div>
      
    </div>

  );
};

export default ManagerWorkFlowCard;
