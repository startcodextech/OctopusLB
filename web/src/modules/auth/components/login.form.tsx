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
                    <br/>
                    <button
                        className="w-full bg-primary-500 text-white px-6 py-5 text-sm font-bold leading-none transition duration-300 md:w-full rounded-xl hover:bg-primary-600 focus:ring-4 focus:ring-primary-100">
                        {t('login')}
                    </button>
                </div>
            </form>
        </>
    )
};

export default LoginForm;