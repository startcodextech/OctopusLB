'use client';
import React, {
  FC,
  useState,
  PropsWithChildren,
  createContext,
  ReactNode, useEffect,
} from 'react';
import { Navbar } from '@components/layouts/dashboard';
import { Sidebar } from '@components/layouts/dashboard/sidebar';

export const Context = createContext<{
  lng: string;
  sidebarOpen: boolean;
  setSidebarOpen: (open: boolean) => void;
  tabNavbar?: ReactNode | null;
  setTabNavbar: (children: ReactNode) => void;
}>({
  lng: 'es',
  sidebarOpen: false,
  setSidebarOpen: (open: boolean) => {},
  setTabNavbar: (children?: ReactNode) => {},
});

type Props = PropsWithChildren<{
  lng: string;
}>;

const Dashboard: FC<Props> = (props) => {
  const { children, lng } = props;

  const [sidebarOpen, setSidebarOpen] = useState(false);
  const [tabNavbar, setTabNavbar] = useState<ReactNode>();

  return (
    <>
      <Context.Provider
        value={{ lng, sidebarOpen, setSidebarOpen, setTabNavbar, tabNavbar }}
      >
        <div className='relative flex h-screen overflow-hidden'>
          <Sidebar />
          <div className='relative flex min-h-lvh flex-1 flex-col overflow-x-hidden overflow-y-hidden'>
            <Navbar>{tabNavbar}</Navbar>
            <main className='h-screen overflow-y-hidden lg:p-4'>
              <div className='mx-auto max-h-full min-h-full overflow-scroll p-4 md:p-6 lg:rounded-2xl 2xl:p-10'>
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
