'use client';
import React, {FC, useEffect, useRef, Fragment} from 'react';
import {usePathname} from "next/navigation";
import Link from 'next/link';

type Props = {
    open: boolean;
    toggle: (open: boolean) => void;
};

const Sidebar: FC<Props> = (props) => {
    const {open, toggle} = props;

    const pathname = usePathname();

    const trigger = useRef<HTMLButtonElement>(null);
    const sidebar = useRef<HTMLElement>(null);

    const [expanded, setExpanded] = React.useState<boolean>(false);

    useEffect(() => {
        const clickHandler = ({target}: MouseEvent) => {
            if (!sidebar.current || !trigger.current) return;
            if (!(target instanceof Node)) return;
            if (!open || sidebar.current.contains(target) || trigger.current.contains(target)) return;
            toggle(false);
        };
        document.addEventListener('click', clickHandler);
        return () => document.removeEventListener('click', clickHandler);
    }, [open, toggle]);

    useEffect(() => {
        const keyHandler = ({key}: KeyboardEvent) => {
            if (!open || key !== 'Escape') return;
            toggle(false);
        }

        document.addEventListener('keydown', keyHandler);
        return () => document.removeEventListener('keydown', keyHandler);
    }, [open, toggle]);

    useEffect(() => {
        if (expanded) {
            document.querySelector('body')?.classList.add('sidebar-expanded');
        } else {
            document.querySelector('body')?.classList.remove('sidebar-expanded');
        }
    }, [expanded]);

    return (
        <>
            <aside
                ref={sidebar}
                className={`absolute left-0 top-0 z-50 flex h-screen w-72 flex-col overflow-y-hidden bg-primary duration-300 ease-linear lg:static lg:translate-x-0 ${open ? 'translate-x-0' : '-translate-x-full'}`}
            >
                <div className="flex items-center justify-between gap-2 px-6 py-5 lg:py-6">
                    <button
                        ref={trigger}
                        onClick={() => toggle(!open)}
                        aria-controls="sidebar"
                        aria-expanded={open}
                        className="block lg:hidden"
                    >
                        <svg
                            className="fill-current"
                            width="20"
                            height="18"
                            viewBox="0 0 20 18"
                            fill="none"
                            xmlns="http://www.w3.org/2000/svg"
                        >
                            <path
                                d="M19 8.175H2.98748L9.36248 1.6875C9.69998 1.35 9.69998 0.825 9.36248 0.4875C9.02498 0.15 8.49998 0.15 8.16248 0.4875L0.399976 8.3625C0.0624756 8.7 0.0624756 9.225 0.399976 9.5625L8.16248 17.4375C8.31248 17.5875 8.53748 17.7 8.76248 17.7C8.98748 17.7 9.17498 17.625 9.36248 17.475C9.69998 17.1375 9.69998 16.6125 9.36248 16.275L3.02498 9.8625H19C19.45 9.8625 19.825 9.4875 19.825 9.0375C19.825 8.55 19.45 8.175 19 8.175Z"
                                fill=""
                            />
                        </svg>
                    </button>
                </div>

                <div className="no-scrollbar flex flex-col overflow-y-auto duration-200 ease-linear">
                    <div className="mt-5 py-4 px-4 lg:mt-9 lg:px-6">
                        <div>
                            <h3 className="mb-4 ml-4 text-sm font-semibold">
                                MENU
                            </h3>

                            <ul className="mb-6 flex flex-col gap-1.5">
                                <li>
                                    <Link
                                        href="/"
                                        className={`group relative flex items-center gap-2.5 rounded-sm py-2 px-4 font-medium duration-300 ease-in-out hover:bg-graydark ${pathname.includes("dhcp") && 'bg-graydark'}`}
                                    >
                                        <i className={`icon-home${pathname === "/" ? '-fill' : ''} text-2xl`}/>
                                        Inicio
                                    </Link>
                                </li>
                                <li>
                                    <Link
                                        href="/dhcp"
                                        className={`group relative flex items-center gap-2.5 rounded-sm py-2 px-4 font-medium duration-300 ease-in-out hover:bg-graydark ${pathname.includes("dhcp") && 'bg-graydark'}`}
                                    >
                                        <i className={`icon-rj45${pathname.includes("dhcp") ? '-fill' : ''} text-2xl`}/>
                                        Load Balancer
                                    </Link>
                                </li>
                                <li>
                                    <Link
                                        href="/dhcp"
                                        className={`group relative flex items-center gap-2.5 rounded-sm py-2 px-4 font-medium duration-300 ease-in-out hover:bg-graydark ${pathname.includes("dhcp") && 'bg-graydark'}`}
                                    >
                                        <i className={`icon-rj45${pathname.includes("dhcp") ? '-fill' : ''} text-2xl`}/>
                                        DHCP
                                    </Link>
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
            </aside>
        </>
    )
}

export default Sidebar;