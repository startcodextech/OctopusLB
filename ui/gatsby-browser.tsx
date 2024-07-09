import './src/styles/global.css';

export const onInitialClientRender = () => {
    if (!isLoggedIn()) {
        if (!/^(\/login\/)$/.test(window.location.pathname)) {
            //window.location.href = "/login/";
        }
    }
}

const isLoggedIn = (): boolean => {
    return !!window.localStorage.getItem("token")

}