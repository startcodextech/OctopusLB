import React, {FC, useState, PropsWithChildren} from "react";
import {Navbar} from "@layouts/dashboard/navbar";

const Dashboard:FC<PropsWithChildren> = (props) => {
    const {children} = props;

    const [sidebarOpen, setSidebarOpen] = useState(false);

    return (
        <>
            <div className="flex h-screen overflow-hidden">
                <div className="relative flex flex-1 flex-col overflow-y-auto overflow-x-hidden">
                    <Navbar sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen}/>

                    <main>
                        <div className="mx-auto max-w-screen-2xl p-4 md:p-6 2xl:p-10">
                            {children}
                        </div>
                    </main>
                </div>
            </div>
        </>
    )
};

export default Dashboard;