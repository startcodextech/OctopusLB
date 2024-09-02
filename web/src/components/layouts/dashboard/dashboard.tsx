'use client';
import React, {FC, useState, PropsWithChildren, createContext} from "react";
import {Navbar} from "@components/layouts/dashboard";
import {Sidebar} from "@components/layouts/dashboard/sidebar";

export const Context = createContext({
    lng: 'es',
})

type Props = PropsWithChildren<{
    lng: string;
}>;

const Dashboard:FC<Props> = (props) => {
    const {children, lng} = props;

    const [sidebarOpen, setSidebarOpen] = useState(false);

    return (
        <>
            <Context.Provider value={{lng}}>
                <div className="flex h-screen overflow-hidden">
                    <Sidebar open={sidebarOpen} toggle={setSidebarOpen}/>
                    <div className="relative flex flex-1 flex-col overflow-y-auto overflow-x-hidden">
                        <Navbar sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen}/>
                        <main>
                            <div className="mx-auto max-w-screen-2xl p-4 md:p-6 2xl:p-10">
                                {children}
                            </div>
                        </main>
                    </div>
                </div>
            </Context.Provider>

        </>
    )
};

export default Dashboard;