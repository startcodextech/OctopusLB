import {resolve} from 'path';
import {CreateWebpackConfigArgs} from "gatsby";

export const onCreateWebpackConfig = ({actions}: CreateWebpackConfigArgs) => {
    actions.setWebpackConfig({
        resolve: {
            alias: {
                '@components': resolve(__dirname, 'src/components'),
                '@styles': resolve(__dirname, 'src/styles'),
                '@images': resolve(__dirname, 'src/images'),
                '@modules': resolve(__dirname, 'src/modules'),
                '@pages': resolve(__dirname, 'src/pages'),
                '@layouts': resolve(__dirname, 'src/layouts'),
            },
        }
    })
};