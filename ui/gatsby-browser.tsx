import './src/styles/global.css';

export const onInitialClientRender = () => {
    if (!isLoggedIn()) {
        if (window.location.pathname !== "/login/") {
            window.location.href = "/login/"
        }
    }
}

const isLoggedIn = (): boolean => {
    return !!window.localStorage.getItem("token")

}