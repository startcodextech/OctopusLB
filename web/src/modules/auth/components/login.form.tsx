import React from 'react';
import {Input, Label} from "@components/commons";

const LoginForm = () => {
    return (
        <>
            <form>
                <div className="mb-4">
                    <Label htmlFor="email">
                        Username *
                    </Label>
                    <Input type="email" id="email" placeholder="admin"
                           fullWidth={true}/>
                </div>
                <div className="mb-4">
                    <Label htmlFor="password">Password *</Label>
                    <Input type="password" id="password" placeholder="123456"
                           fullWidth={true}/>
                </div>
                <div className="mt-12 mb-6 flex justify-center">
                    <button
                        className="w-full px-6 py-5 text-sm font-bold leading-none text-white transition duration-300 md:w-96 rounded-2xl hover:bg-primary-600 focus:ring-4 focus:ring-primary-100 bg-primary">
                        Login
                    </button>
                </div>
            </form>
        </>
    )
};

export default LoginForm;