'use client';
import React, {FC, PropsWithChildren} from "react";

type ActionProps = PropsWithChildren<{
    show?: boolean;
    onClick?: () => void;
    disabled?: boolean;
}>;

const Action: FC<ActionProps> = (props) => {
    const {children, show = true, disabled = false, onClick} = props;
    return (
        <button
            onClick={onClick}
            disabled={disabled}
            className={`${show ? '': 'hidden'} w-11 h-11 flex items-center justify-center p-2.5 rounded-[0.5rem] border-2 bg-white border-neutral-background hover:bg-neutral-background text-2xl`.trim()}>
            {children}
        </button>
    )
};

export default Action;