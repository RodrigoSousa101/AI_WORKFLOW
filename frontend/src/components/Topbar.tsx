import { type ReactNode } from 'react';

type Props = {
  children: ReactNode;
};

const Topbar = ({ children }: Props) => {
    return (
        <div className="w-full h-16 flex items-center px-40 border border-gray-200">
            {children}
        </div>
    )
}

export default Topbar;