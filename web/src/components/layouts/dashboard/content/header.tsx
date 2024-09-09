'use client';

import React, {FC, useContext, PropsWithChildren} from 'react';
import Image from "next/image";
import {Chip} from "@components/commons";
import {Context} from "@components/layouts/dashboard/dashboard";
import {useTranslation} from "@app/i18n/client";

type Props = PropsWithChildren<{
    name: string;
    status?: 'starting' | 'running' | 'stopped' | 'error';
    imageUrl?: string;
}>;

const ContentHeader: FC<Props> = (props) => {
    const {name, status = 'stopped', children, imageUrl} = props;

    const {lng} = useContext(Context);
    const {t} = useTranslation(lng);

    const chipType = status === 'running' ? 'success' : status === 'error' || status === 'stopped'  ? 'error' : 'warning';

    return (
        <div className="flex items-center">
            <div className="flex items-center gap-2 w-full text-text-primary">
                {imageUrl && <Image src={imageUrl} alt="" width={44} height={44}/>}
                <span className="text-2xl text-text-primary font-bold">{name}</span>
                <Chip type={chipType} showDot={true}>{t(status)}</Chip>
            </div>

            <div className="flex items-center flex-row gap-2">
                {children}
            </div>

        </div>
    )
};

export default ContentHeader;