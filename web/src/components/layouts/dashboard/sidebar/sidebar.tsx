'use client';
import React, {FC, useEffect, useRef, useContext} from 'react';
import {usePathname} from "next/navigation";
import Link from 'next/link';
import Image from "next/image";
import {Context} from "@components/layouts/dashboard/dashboard";
import {useTranslation} from "@app/i18n/client";

type Props = {
    open: boolean;
    toggle: (open: boolean) => void;
};

const Active: FC = () => (
    <>
        <div className="absolute left-0 h-full w-1 bg-primary-500 rounded-tr-xl rounded-br-xl"/>
    </>
)

const Sidebar: FC<Props> = (props) => {
    const {open, toggle} = props;

    const pathname = usePathname();

    const context = useContext(Context);
    const {lng} = context;

    const trigger = useRef<HTMLButtonElement>(null);
    const sidebar = useRef<HTMLElement>(null);

    const {t} = useTranslation(lng, 'menu');

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
                className={`absolute left-0 top-0 z-50 flex h-screen w-72 flex-col overflow-y-hidden bg-sidebar dark:bg-sidebar-dark duration-300 ease-linear lg:static lg:translate-x-0 ${open ? 'translate-x-0' : '-translate-x-full'}`}
            >
                <div className="flex items-center justify-between gap-2 px-6 py-5 lg:py-6">
                    <Link href="/" className="flex items-center font-black text-[24px] text-text-primary">
                        <Image src="/images/logo.png" className="mr-3" alt="Octopus LB" width={36} height={36}/>
                        Octopus LB
                    </Link>

                    <span className="bg-success-dark-1">hajhsk kbakjshdb</span>

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
                    <ul className="mb-6 flex flex-col gap-1.5">
                        {
                            [
                                {text: t('dashboard'), icon: 'icon-dashboard', href: '/'},
                                {text: t('balancer'), icon: 'icon-balancer', href: '/balancer'},
                                {text: t('dhcp'), icon: 'icon-dhcp', href: '/dhcp'},
                                {text: t('firewall'), icon: 'icon-firewall', href: '/firewall'},
                                {text: t('dns'), icon: 'icon-dns', href: '/dns'},
                                {text: t('bgp'), icon: 'icon-bgp', href: '/bgp'},
                            ].map((item, index) => {

                                const isActive = item.href === '/' ? pathname === item.href : pathname.includes(item.href);

                                return (
                                    <li key={index}
                                        className={`group relative font-medium px-4 lg:px-6 ${isActive ? 'text-text-primary !font-bold' : 'text-text-secondary'} hover:text-text-primary cursor-pointer ease-in-out`}>
                                        {isActive && <Active/>}
                                        <Link
                                            href={`/${lng}${item.href}`}
                                            className={`${isActive ? 'bg-primary-100 rounded-xl' : ''} group relative flex items-center gap-2.5 rounded-sm py-2 px-4`}
                                        >
                                            <i className={`${item.icon} text-2xl`}/>
                                            {item.text}
                                        </Link>
                                    </li>
                                )
                            })
                        }
                    </ul>
                </div>
            </aside>
        </>
    )
}

export default Sidebar;