'use client';
import React, {FC, PropsWithChildren, ButtonHTMLAttributes} from 'react';

type Props = PropsWithChildren<{
    variant?: 'solid' | 'outline' | 'text';
    color?: 'primary' | string;
    textColor?: string;
    full?: boolean | string;
}> & ButtonHTMLAttributes<HTMLButtonElement>;

const Button: FC<Props> = (props) => {
    const {children, variant = 'solid', color = 'primary', textColor = 'black', full = false, ...rest} = props;

    const typeButton = (): string => {
        const bg = color === 'primary' ? 'primary-500' : color;

        let css = '';

        switch (variant) {
            case 'solid':
                css += `bg-${bg} text-${textColor} border-0`;
            case 'text':
                    css += `border-0 bg-transparent text-${textColor}`;
            default:
                css += `bg-transparent text-black border-2`;
        }

        return css;
    }

    const classFull = (): string => {
        if (full === "") return '';
        if (typeof full === 'string') {
            if (full === 'true') {
                return ' w-full';
            }
            return ` ${full}`;
        }
        return full ? ' w-full' : '';
    };

    return (
        <>
            <button
                {...rest}
                className={`${typeButton()}${classFull()} px-6 py-2.5 rounded-xl font-bold text-md`.trim()}
            >
                {children}
            </button>
        </>
    )
};

export default Button;