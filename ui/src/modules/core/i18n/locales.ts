type Data = {
    locales: Locales
}

type Locales = {
    edges: LocalesEdge[]
    site: {
        siteMetadata: {
            title: string
        }
    }
}

type LocalesEdge = {
    node: LocalesEdgeNode
}

type LocalesEdgeNode = {
    ns: string
    data: string
    language: string
}

const getLocales = (page: string, data: object): any => {
    if (!data || !data.hasOwnProperty('locales')) {
        return {};
    }
    const dataLanguage: string = (data as Data).locales.edges.find(x => x.node.ns === page )?.node.data || '{}';
    const commonLanguage: string = (data as Data).locales.edges.find(x => x.node.ns === 'common' )?.node.data || '{}';

    const lngPage = JSON.parse(dataLanguage);
    const lngCommon = JSON.parse(commonLanguage);
    return {...lngPage, ...lngCommon};
}

export default getLocales;