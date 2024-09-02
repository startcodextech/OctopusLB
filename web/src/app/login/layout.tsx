import type {Metadata} from "next";
import "@app/[lng]/globals.css";

export const metadata: Metadata = {
    title: "Create Next App",
    description: "Generated by create next app",
};

const RootLayout = ({children}: Readonly<{ children: React.ReactNode; }>) => {
    return (
        <>
            <html lang="es">
                <body className="bg-[url(/images/bg.jpg)] !bg-cover !bg-no-repeat bg-center h-screen overflow-x-hidden p-0 m-0">
                    {children}
                </body>
            </html>
        </>
    )
};

export default RootLayout;