'use client';
import React, { FC, PropsWithChildren, useContext } from 'react';
import Image from 'next/image';
import Search from './search';
import ToggleSidebar from './toggle-sidebar';
import DropdownUser from './user';
import DropdownNotifications from './notifications';
import { Context } from '@components/layouts/dashboard/dashboard';

type Props = PropsWithChildren<{}>;

const Navbar: FC<Props> = (props) => {
  const { children } = props;
  const { sidebarOpen, setSidebarOpen, ...rest } = useContext(Context);

  console.log('rest', rest);
  return (
    <>
      <header className='sticky top-0 z-40 flex w-full drop-shadow'>
        <div className='shadow-2 flex flex-grow items-center justify-between px-4 py-5 md:px-6 2xl:px-11'>
          <div className='flex items-center gap-2 sm:gap-4 xl:hidden'>
            <ToggleSidebar
              sidebarOpen={sidebarOpen}
              setSidebarOpen={setSidebarOpen}
            />
            {!children && (
              <a href='/' className='block flex-shrink-0 xl:hidden'>
                <Image
                  src='/images/logo.png'
                  alt={''}
                  className='w-8'
                  width={24}
                  height={24}
                />
              </a>
            )}
            {children}
          </div>
          <Search />
          <div className='flex items-center gap-3 2xsm:gap-7'>
            <ul className='flex items-center gap-2 2xsm:gap-4'>
              <DropdownNotifications />
            </ul>
            <DropdownUser />
          </div>
        </div>
      </header>
    </>
  );
};

export default Navbar;
