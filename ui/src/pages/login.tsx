import React, {FC} from 'react';
import {graphql, HeadFC, PageProps} from "gatsby";
import {Trans, useI18next} from "@herob191/gatsby-plugin-react-i18next";
import {LoginForm} from "@modules/auth";
import {getLocales} from "@modules/core/i18n";
import {StaticImage} from "gatsby-plugin-image";
import bgImage from "@images/bg.jpg";
import {version} from '../../package.json';



const LoginPage: FC<PageProps> = () => {
    const {t} = useI18next();
    return (
        <>
            <div className="container max-w-lg w-full mx-auto mt-8">
                <div className="p-8 bg-[rgba(255,255,255,.7)] backdrop-blur-lg rounded-3xl">
                    <div className="flex items-center flex-col justify-center">
                        <StaticImage placeholder="none" loading="lazy" src="../images/icon.png" alt={""} className="w-28 mb-4"/>
                        <h2 className="mb-3 text-4xl font-extrabold text-center">
                            <Trans i18nKey="app_name"/>
                        </h2>
                    </div>
                    <h3 className="text-center mb-4">
                        <Trans i18nKey="title"/>
                    </h3>

                    <LoginForm/>
                </div>
            </div>


            <div className="flex flex-wrap -mx-3 my-5">
                <div className="w-full max-w-full sm:w-3/4 mx-auto text-center">
                    <p className="text-sm text-white py-1 font-medium">
                        OctopusLB <a
                        href="https://github.com/startcodextech/OctopusLB"
                        target="_blank">v{version}</a> by <a
                        href="https://startcodex.com" target="_blank">Start
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
            <body className="bg-gray-200 !bg-cover !bg-no-repeat bg-center h-screen overflow-x-hidden" style={{background: `url(${bgImage})`}}/>
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