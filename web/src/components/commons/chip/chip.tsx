'use client';
import React, {FC, PropsWithChildren} from 'react';

type Props = PropsWithChildren<{
    children?: React.ReactNode;
    type?: 'success' | 'warning' | 'error';
    showDot?: boolean;
}>;

const Chip: FC<Props> = (props) => {
    const {children, type = 'success', showDot = false} = props;

    return (
        <>
            <div
                className={`flex items-center gap-2 py-0 px-4 bg-chip-${type}-background text-chip-${type}-text text-sm font-bold leading-6 rounded-lg`}>
                {showDot && <div className={`rounded-full w-2.5 h-2.5 bg-chip-${type}-dot`}/>}
                {children}
            </div>
        </>
    )
}

export default Chip;