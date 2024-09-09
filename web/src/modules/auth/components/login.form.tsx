'use client';
import React, {FC, useContext} from 'react';
import {Button, Input, Label} from "@components/commons";
import {useTranslation} from "@app/i18n/client";

type Props = {
    lng: string;
};

const LoginForm: FC<Props> = ({lng}) => {

    const {t} = useTranslation(lng, 'login');

    return (
        <>
            <form>
                <div className="mb-4">
                    <Label htmlFor="email">
                        {t('username')} *
                    </Label>
                    <Input type="email" id="email" placeholder="admin"
                           fullWidth={true}/>
                </div>
                <div className="mb-4">
                    <Label htmlFor="password">{t('password')} *</Label>
                    <Input type="password" id="password" placeholder="123456"
                           fullWidth={true}/>
                </div>
                <div className="mt-12 mb-6 flex justify-center flex-col">
                    <Button full="true">{t('login')}</Button>
                </div>
            </form>
        </>
    )
};

export default LoginForm;