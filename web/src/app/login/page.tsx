import Image from 'next/image';
import {LoginForm} from "@modules/auth/components";
import {version} from "../../../package.json";

export default function Login() {
    return (
        <>
            <div className="container max-w-lg w-full mx-auto pt-8">
                <div className="p-8 bg-[rgba(255,255,255,.7)] backdrop-blur-lg rounded-3xl mx-4 sm:mx-0">
                    <div className="flex items-center flex-col justify-center">
                        <Image src="/images/icon.png" alt="Logo" className="w-24 pb-4" width={96} height={96}/>
                        <h2 className="pb-3 text-4xl font-extrabold text-center">
                            OctopusLB
                        </h2>
                    </div>
                    <h3 className="text-center mb-4">
                        Login to your account
                    </h3>

                    <LoginForm/>
                </div>
            </div>


            <div className="flex flex-wrap -px-3 py-5">
                <div className="w-full max-w-full sm:w-3/4 mx-auto text-center">
                    <p className="text-sm text-white py-1 font-medium">
                        OctopusLB &nbsp;
                        <a href="https://github.com/startcodextech/OctopusLB"
                           target="_blank">
                            v{version}
                        </a> by &nbsp;
                        <a href="https://startcodex.com" target="_blank">
                            Start Codex
                        </a> © {new Date().getFullYear()}.
                    </p>
                </div>
            </div>
        </>
    )
};