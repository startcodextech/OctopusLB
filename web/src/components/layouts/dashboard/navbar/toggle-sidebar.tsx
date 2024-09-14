import React, { FC } from 'react';

type Props = {
  sidebarOpen: boolean;
  setSidebarOpen: (open: boolean) => void;
};

const ToggleSidebar: FC<Props> = (props) => {
  const { sidebarOpen, setSidebarOpen } = props;

  const handleToggle = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.stopPropagation();
    setSidebarOpen(!sidebarOpen);
  };

  return (
    <>
      <button
        aria-controls='sidebar'
        onClick={handleToggle}
        className='z-50 block rounded border border-base-dark bg-transparent p-1 shadow-none ring-0 xl:hidden dark:border-base'
      >
        <span className='relative block h-5 w-5 cursor-pointer'>
          <span className='absolute right-0 h-full w-full'>
            <span
              className={`relative left-0 top-0 my-1 block h-0.5 w-0 rounded-sm bg-frame-dark delay-[0] duration-200 ease-in-out dark:bg-frame ${!sidebarOpen && '!w-full delay-300'}`}
            />
            <span
              className={`relative left-0 top-0 my-1 block h-0.5 w-0 rounded-sm bg-frame-dark delay-150 duration-200 ease-in-out dark:bg-frame ${!sidebarOpen && 'delay-400 !w-full'}`}
            />
            <span
              className={`relative left-0 top-0 my-1 block h-0.5 w-0 rounded-sm bg-frame-dark delay-200 duration-200 ease-in-out dark:bg-frame ${!sidebarOpen && '!w-full delay-500'}`}
            />
          </span>
          <span className='absolute right-0 h-full w-full rotate-45'>
            <span
              className={`absolute left-[10px] top-0 block h-full w-0.5 rounded-sm bg-frame-dark delay-300 duration-200 ease-in-out dark:bg-frame ${!sidebarOpen && '!h-0 !delay-[0]'}`}
            />
            <span
              className={`delay-400 absolute left-0 top-[9px] block h-0.5 w-full rounded-sm bg-frame-dark duration-200 ease-in-out dark:bg-frame ${!sidebarOpen && '!h-0 !delay-200'}`}
            />
          </span>
        </span>
      </button>
    </>
  );
};

export default ToggleSidebar;
