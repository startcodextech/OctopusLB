'use client';

import React, { FC, PropsWithChildren } from 'react';

type Props = PropsWithChildren<{
  title: string;
  text: string;
}>;

const InfoItem: FC<Props> = (props) => {
  const { title, text } = props;
  return (
    <>
      <div className='flex flex-row gap-2 lg:flex-col'>
        <p className='text-text-secondary text-base font-bold'>{title}</p>
        <span className='block text-base font-medium text-black'>{text}</span>
      </div>
    </>
  );
};

export default InfoItem;
