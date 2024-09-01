'use client';
import React, {FC, useEffect, useRef, Fragment} from 'react';
import {usePathname} from "next/navigation";
import Link from 'next/link';
import LinkGroup from "@components/layouts/dashboard/sidebar/link-group";
import Image from "next/image";

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
                className={`absolute left-0 top-0 z-50 flex h-screen w-72 flex-col overflow-y-hidden bg-sidebar duration-300 ease-linear lg:static lg:translate-x-0 ${open ? 'translate-x-0' : '-translate-x-full'}`}
            >
                <div className="flex items-center justify-between gap-2 px-6 py-5 lg:py-6">
                    <Link href="/" className="flex items-center font-black text-[24px]">
                        <Image src="/images/logo.png" className="mr-3" alt="Octopus LB" width={36} height={36}/>
                        Octopus LB
                    </Link>

                    <button
                        ref={trigger}
                        onClick={() => toggle(!open)}
                        aria-controls="sidebar"
                        aria-expanded={open}
                        className="block"
                    >
                        <svg className="fill-current" width="24" height="24" viewBox="0 0 24 24" fill="none">
                            <path
                                d="M23.0625 11.0625H0.9375C0.419719 11.0625 0 11.4822 0 12C0 12.5178 0.419719 12.9375 0.9375 12.9375H23.0625C23.5803 12.9375 24 12.5178 24 12C24 11.4822 23.5803 11.0625 23.0625 11.0625Z"
                                fill=""/>
                            <path
                                d="M23.0625 3.5625H0.9375C0.419719 3.5625 0 3.98222 0 4.5C0 5.01778 0.419719 5.4375 0.9375 5.4375H23.0625C23.5803 5.4375 24 5.01778 24 4.5C24 3.98222 23.5803 3.5625 23.0625 3.5625Z"
                                fill=""/>
                            <path
                                d="M23.0625 18.5625H0.9375C0.419719 18.5625 0 18.9822 0 19.5C0 20.0178 0.419719 20.4375 0.9375 20.4375H23.0625C23.5803 20.4375 24 20.0178 24 19.5C24 18.9822 23.5803 18.5625 23.0625 18.5625Z"
                                fill=""/>
                        </svg>
                    </button>
                </div>

                <div className="no-scrollbar flex flex-col overflow-y-auto duration-200 ease-linear">
                    <div className="mt-5 py-4 lg:mt-9">
                        <div>
                            <ul className="mb-6 flex flex-col gap-1.5">
                                {
                                    [
                                        {text: 'Inicio', icon: 'icon-dashboard', href: '/'},
                                        {text: 'Load Balancer', icon: 'icon-balancer', href: '/balancer'},
                                        {text: 'DHCP', icon: 'icon-dhcp', href: '/dhcp'},
                                        {text: 'Firewall', icon: 'icon-firewall', href: '/firewall'},
                                        {text: 'DNS', icon: 'icon-dns', href: '/dns'},
                                        {text: 'BGP', icon: 'icon-bgp', href: '/bgp'},
                                    ].map((item, index) => {

                                        const isActive = item.href === '/' ? pathname === item.href : pathname.includes(item.href);

                                        return (
                                            <li key={index}
                                                className={`group font-medium px-4 lg:px-6 ${isActive ? 'text-black' : 'text-sidebar-text'} hover:text-black cursor-pointer duration-100 ease-in-out`}>
                                                <Link
                                                    href={item.href}
                                                    className={`group relative flex items-center gap-2.5 rounded-sm py-2 px-4`}
                                                >
                                                    <i className={`${item.icon} group-hover:text-black text-2xl ${isActive ? 'text-black' : 'text-sidebar-icon'}`}/>
                                                    {item.text}
                                                </Link>
                                            </li>
                                        )
                                    })
                                }
                                <li>
                                    <Link
                                        href="/balancer"
                                        className={`group relative flex items-center gap-2.5 rounded-sm py-2 px-4 font-medium duration-300 ease-in-out hover:bg-graydark ${pathname.includes("dhcp") && 'bg-graydark'}`}
                                    >
                                        <i className="icon-balancer"/>
                                        Load Balancer
                                    </Link>
                                </li>
                                <li>
                                    <Link
                                        href="/dhcp"
                                        className={`group relative flex items-center gap-2.5 rounded-sm py-2 px-4 font-medium duration-300 ease-in-out hover:bg-graydark ${pathname.includes("dhcp") && 'bg-graydark'}`}
                                    >
                                        <i className="icon-dhcp"/>
                                        DHCP
                                    </Link>
                                </li>
                                <li>
                                    <Link
                                        href="/firewall"
                                        className={`group relative flex items-center gap-2.5 rounded-sm py-2 px-4 font-medium duration-300 ease-in-out hover:bg-graydark ${pathname.includes("dhcp") && 'bg-graydark'}`}
                                    >
                                        <i className="icon-firewall"/>
                                        Firewall
                                    </Link>
                                </li>
                                <li>
                                    <Link
                                        href="/dns"
                                        className={`group relative flex items-center gap-2.5 rounded-sm py-2 px-4 font-medium duration-300 ease-in-out hover:bg-sidebar ${pathname.includes("dhcp") && 'bg-graydark'}`}
                                    >
                                        <i className="icon-dns"/>
                                        DNS
                                    </Link>
                                </li>
                                <li>
                                    <Link
                                        href="/bgp"
                                        className={`group relative flex items-center gap-2.5 rounded-sm py-2 px-4 font-medium duration-300 ease-in-out hover:bg-graydark ${pathname.includes("dhcp") && 'bg-graydark'}`}
                                    >
                                        <i className="icon-bgp"/>
                                        BGP
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