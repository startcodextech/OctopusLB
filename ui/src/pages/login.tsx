import React, {FC} from 'react';
import {graphql, HeadFC, PageProps} from "gatsby";
import {Trans, useI18next} from "gatsby-plugin-react-i18next";
import {LoginForm} from "@modules/auth";
import {getLocales} from "@modules/core/i18n";
import {version} from '../../package.json';
import {StaticImage} from "gatsby-plugin-image";


const LoginPage: FC<PageProps> = () => {
    const {t} = useI18next();
    return (
        <>
            <div className="container flex flex-col mx-auto bg-white rounded-lg pt-12 my-5">
                <div
                    className="flex justify-center w-full h-full my-auto xl:gap-14 lg:justify-normal md:gap-5">
                    <div className="flex items-center justify-center w-full">
                        <div className="flex items-center max-w-lg w-full">
                            <div className="flex flex-col w-full h-full bg-white rounded-3xl">
                                <div className="flex items-center flex-col justify-center">
                                    <StaticImage src="../images/icon.png" alt={""} className="w-28 mb-4"/>
                                    <h2 className="mb-3 text-4xl font-extrabold text-grey-dark-900 text-center">
                                        <Trans i18nKey="app_name"/>
                                    </h2>
                                </div>
                                <p className="text-center text-grey-dark-700 mb-4">
                                    <Trans i18nKey="title"/>
                                </p>
                                <div className="flex items-center justify-center my-5">
                                    <hr className="border-b border-grey-500 w-full"/>
                                </div>
                                <LoginForm/>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div className="flex flex-wrap -mx-3 my-5">
                <div className="w-full max-w-full sm:w-3/4 mx-auto text-center">
                    <p className="text-sm text-slate-500 py-1">
                        ManagerLB <a
                        href="https://github.com/startcodextech/managerlb"
                        className="text-slate-700 hover:text-slate-900" target="_blank">v{version}</a> by <a
                        href="https://startcodex.com" className="text-slate-700 hover:text-slate-900" target="_blank">Start
                        Codex</a> © {new Date().getFullYear()}.
                    </p>
                </div>
            </div>
        </>
    )
};

export default LoginPage;

export const Head: HeadFC = ({data}) => {
    const lang = getLocales('login', data);

    return (
        <>
            <body className="bg-gray-200"/>
            <title>
                {lang.title} | {lang.app_name}
            </title>
        </>
    )
}

export const query = graphql`
  query ($language: String!) {
    site {
        siteMetadata {
            title
        }
    }
    locales: allLocale(
      filter: { ns: { in: ["common", "login"] }, language: { eq: $language } }
    ) {
      edges {
        node {
          ns
          data
          language
        }
      }
    }
  }
`;