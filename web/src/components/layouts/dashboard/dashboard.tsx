'use client';
import React, { FC, useState, PropsWithChildren, createContext } from 'react';
import { Navbar } from '@components/layouts/dashboard';
import { Sidebar } from '@components/layouts/dashboard/sidebar';

export const Context = createContext({
  lng: 'es',
});

type Props = PropsWithChildren<{
  lng: string;
}>;

const Dashboard: FC<Props> = (props) => {
  const { children, lng } = props;

  const [sidebarOpen, setSidebarOpen] = useState(false);

  return (
    <>
      <Context.Provider value={{ lng }}>
        <div className='flex h-screen overflow-hidden'>
          <Sidebar open={sidebarOpen} toggle={setSidebarOpen} />
          <div className='bg-neutral-background relative flex min-h-lvh flex-1 flex-col overflow-x-hidden overflow-y-hidden'>
            <Navbar sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen} />
            <main className='h-screen overflow-y-hidden lg:p-4'>
              <div className='mx-auto max-h-full min-h-full max-w-screen-2xl overflow-scroll bg-white p-4 md:p-6 lg:rounded-2xl 2xl:p-10'>
                {children}
              </div>
            </main>
          </div>
        </div>
      </Context.Provider>
    </>
  );
};

export default Dashboard;
