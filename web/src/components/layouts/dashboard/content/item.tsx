'use client';

import React, {FC, PropsWithChildren} from 'react';

type Props = PropsWithChildren<{
    title: string;
    text: string;
}>;

const InfoItem: FC<Props> = (props) => {
    const {title, text} = props;
    return (
        <>
            <div className="flex flex-row lg:flex-col gap-2">
                <p className="font-bold text-base text-text-secondary">
                    {title}
                </p>
                <span className="block text-base text-black font-medium">{text}</span>
            </div>
        </>
    )
}

export default InfoItem;