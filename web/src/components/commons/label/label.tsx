'use client';
import React, { FC, LabelHTMLAttributes } from 'react';

const Label: FC<LabelHTMLAttributes<HTMLLabelElement>> = (props) => {
  return (
    <>
      <label
        {...props}
        className='text-grey-900 mb-2 inline-block text-start text-base font-medium'
      />
    </>
  );
};

export default Label;
