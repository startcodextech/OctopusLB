'use client';
import React, { FC, PropsWithChildren } from 'react';

type ActionProps = PropsWithChildren<{
  show?: boolean;
  onClick?: () => void;
  disabled?: boolean;
}>;

const Action: FC<ActionProps> = (props) => {
  const { children, show = true, disabled = false, onClick } = props;
  return (
    <button
      onClick={onClick}
      disabled={disabled}
      className={`${show ? '' : 'hidden'} border-neutral-background hover:bg-neutral-background flex h-11 w-11 items-center justify-center rounded-[0.5rem] border-2 bg-white p-2.5 text-2xl`.trim()}
    >
      {children}
    </button>
  );
};

export default Action;
