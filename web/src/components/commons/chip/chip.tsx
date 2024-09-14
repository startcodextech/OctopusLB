'use client';
import React, { FC, PropsWithChildren } from 'react';

type Props = PropsWithChildren<{
  children?: React.ReactNode;
  type?: 'success' | 'warning' | 'error';
  showDot?: boolean;
}>;

const Chip: FC<Props> = (props) => {
  const { children, type = 'success', showDot = false } = props;

  return (
    <>
      <div
        className={`flex items-center gap-2 px-4 py-0 bg-chip-${type}-background text-chip-${type}-text rounded-lg text-sm font-bold leading-6`}
      >
        {showDot && (
          <div className={`h-2.5 w-2.5 rounded-full bg-chip-${type}-dot`} />
        )}
        {children}
      </div>
    </>
  );
};

export default Chip;
