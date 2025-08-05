import { RiVipCrownLine } from "react-icons/ri";

const WorkFlowCard = () => {
    return(
        <div className="border border-gray-300 w-100 h-70 rounded-xl shadow-xs p-4">
            <div className="flex gap-4 items-center justify-between">
                <p className="font-semibold text-xl">App Development</p>
                <div className="w-20 h-5 bg-indigo-200 text-indigo-800 font-semibold text-xs rounded-2xl flex items-center justify-center">
                    <RiVipCrownLine size={16} className="mr-1" />
                    <p>Manager</p>
                </div>
                <p>Status</p>
            </div>
            <p className="max-w-70 break-words text-sm text-gray-700 mt-2">Workflow for mobile application development</p>
        </div>
    )
}

export default WorkFlowCard;