import React, {FC} from "react";
import {Link} from "gatsby";
import {StaticImage} from "gatsby-plugin-image";
import Search from "./search";
import ToggleSidebar from "./toggle-sidebar";
import DropdownUser from "./dropdown-user";

type Props = {
    sidebarOpen: boolean;
    setSidebarOpen: (open: boolean) => void;
};

const Navbar: FC<Props> = (props) => {
    const {sidebarOpen, setSidebarOpen} = props;
    return (
        <>
            <header
                className="sticky top-0 z-999 flex w-full bg-white drop-shadow-1">
                <div className="flex flex-grow items-center justify-between px-4 py-4 shadow-2 md:px-6 2xl:px-11">
                    <div className="flex items-center gap-2 sm:gap-4 lg:hidden">
                        <ToggleSidebar sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen}/>
                        <Link to="/" className="block flex-shrink-0 lg:hidden">
                            <StaticImage src={"../../../images/icon.png"} alt={""} className="w-9 mr-2" />
                        </Link>
                    </div>
                    <Search />
                    <div className="flex items-center gap-3 2xsm:gap-7">
                        <ul className="flex items-center gap-2 2xsm:gap-4">
                        </ul>
                        <DropdownUser/>
                    </div>
                </div>
            </header>
        </>
    )
};

export default Navbar;