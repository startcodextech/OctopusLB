import type {GatsbyConfig} from "gatsby";
import {languages, defaultLanguage} from "./languages";

const siteUrl = `https://www.yourdomain.tld`

const config: GatsbyConfig = {
    siteMetadata: {
        title: `ManagerLB`,
        siteUrl,
    },
    // More easily incorporate content into your pages through automatic TypeScript type generation and better GraphQL IntelliSense.
    // If you use VSCode you can also use the GraphQL plugin
    // Learn more at: https://gatsby.dev/graphql-typegen
    graphqlTypegen: true,
    plugins: [
        "gatsby-plugin-postcss",
        {
            resolve: `gatsby-source-filesystem`,
            options: {
                path: `${__dirname}/locales`,
                name: `locale`
            }
        },
        {
            resolve: "gatsby-plugin-react-i18next",
            options: {
                localeJsonSourceName: `locale`,
                languages,
                defaultLanguage,
                siteUrl,
                i18nextOptions: {
                    fallbackLng: defaultLanguage,
                    supportedLngs: languages,
                    defaultNS: "common",
                    interpolation: {
                        escapeValue: false
                    },
                }
            }
        }
    ]
};

export default config;
