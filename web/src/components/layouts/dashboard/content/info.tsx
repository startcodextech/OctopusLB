'use client';

import React, { FC, PropsWithChildren } from 'react';

type Props = PropsWithChildren<{}>;

const HeaderInfo: FC<Props> = (props) => {
  const { children } = props;

  return (
    <>
      <div className='flex flex-row items-center gap-4 px-0 py-8'>
        <div className='flex w-full flex-wrap content-center items-center justify-between gap-y-4'>
          {children}
        </div>
        <button className='hover:bg-neutral-background flex max-h-11 max-w-11 items-center justify-between rounded-[0.5rem] p-2.5 text-3xl'>
          <i className='icon-setting' />
        </button>
      </div>
    </>
  );
};

export default HeaderInfo;
