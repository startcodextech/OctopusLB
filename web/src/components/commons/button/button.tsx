'use client';
import React, {FC, PropsWithChildren, ButtonHTMLAttributes} from 'react';

type Props = PropsWithChildren<{
    variant?: 'solid' | 'outline' | 'text';
    color?: 'primary' | string;
    textColor?: string;
    width?: string | number;
}> & ButtonHTMLAttributes<HTMLButtonElement>;

const Button: FC<Props> = (props) => {
    const {children, variant = 'solid', color = 'primary', textColor = 'white', width, ...rest} = props;

    const typeButton = (): string => {
        const bg = color === 'primary' ? 'primary-500' : color;

        let css = '';

        switch (variant) {
            case 'solid':
                css += `bg-${bg} text-${textColor} border-0`;
                break;
            case 'text':
                css += `border-0 bg-transparent text-${textColor}`;
                break;
            default:
                css += `bg-transparent text-black border-2`;
                break;
        }

        return css;
    }

    const classWidth = (): string => {
        if (width === undefined) return '';
        if (width === '') return '';
        return width ? ` w-${width}` : '';
    };

    return (
        <>
            <button
                {...rest}
                className={`${typeButton()}${classWidth()} flex items-center justify-center px-6 py-2.5 rounded-xl font-bold text-md max-h-11`.trim()}
            >
                {children}
            </button>
        </>
    )
};

export default Button;