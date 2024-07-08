import React from 'react';
import {Input, Label} from "@components";
import {Trans, useI18next} from "gatsby-plugin-react-i18next";

const LoginForm = () => {
    const {t} = useI18next();
    return (
        <>
            <form>
                <div className="mb-4">
                    <Label htmlFor="email">
                        <Trans i18nKey="username"/> *
                    </Label>
                    <Input type="email" id="email" placeholder={t('user_placeholder')}
                           fullWidth={true}/>
                </div>
                <div className="mb-4">
                    <Label htmlFor="password"><Trans i18nKey="password"/> *</Label>
                    <Input type="password" id="password" placeholder={t('password_placeholder')}
                           fullWidth={true}/>
                </div>
                <div className="my-12 flex justify-center">
                    <button
                        className="w-full px-6 py-5 mb-5 text-sm font-bold leading-none text-white transition duration-300 md:w-96 rounded-2xl hover:bg-purple-blue-600 focus:ring-4 focus:ring-purple-blue-100 bg-purple-blue-500">
                        <Trans i18nKey="title"/>
                    </button>
                </div>
            </form>
        </>
    )
};

export default LoginForm;