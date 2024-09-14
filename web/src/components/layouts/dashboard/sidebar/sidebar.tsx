'use client';
import React, { FC, useEffect, useRef, useContext, Fragment } from 'react';
import { usePathname } from 'next/navigation';
import Link from 'next/link';
import Image from 'next/image';
import { Context } from '@components/layouts/dashboard/dashboard';
import { useTranslation } from '@app/i18n/client';

type Props = {
  open: boolean;
  toggle: (open: boolean) => void;
};

const Active: FC = () => (
  <>
    <div className='absolute left-4 top-0 h-full w-12 rounded-lg bg-frame-dark dark:bg-frame' />
  </>
);

const Sidebar: FC<Props> = (props) => {
  const { open, toggle } = props;

  const pathname = usePathname();

  const context = useContext(Context);
  const { lng } = context;

  const trigger = useRef<HTMLButtonElement>(null);
  const sidebar = useRef<HTMLElement>(null);

  const { t } = useTranslation(lng, 'menu');

  const [expanded, setExpanded] = React.useState<boolean>(false);

  useEffect(() => {
    const clickHandler = ({ target }: MouseEvent) => {
      if (!sidebar.current || !trigger.current) return;
      if (!(target instanceof Node)) return;
      if (
        !open ||
        sidebar.current.contains(target) ||
        trigger.current.contains(target)
      )
        return;
      toggle(false);
    };
    document.addEventListener('click', clickHandler);
    return () => document.removeEventListener('click', clickHandler);
  }, [open, toggle]);

  useEffect(() => {
    const keyHandler = ({ key }: KeyboardEvent) => {
      if (!open || key !== 'Escape') return;
      toggle(false);
    };

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
      <div
        className={`absolute left-0 top-0 z-50 flex h-screen w-72 flex-col overflow-y-hidden duration-300 ease-linear lg:static lg:translate-x-0 lg:p-2 ${open ? 'translate-x-0' : '-translate-x-full'}`}
      >
        <aside className='h-screen w-full bg-frame lg:rounded-3xl dark:bg-frame-dark'>
          <div className='flex justify-center py-6 text-xl font-bold text-black dark:text-white'>
            <Link href='/' className='flex flex-row items-center gap-3'>
              <Image src='/images/logo.png' alt='logo' width={36} height={36} />
              Octopus LB
            </Link>
          </div>

          <div className='flex flex-col items-center py-11 text-primary-text dark:text-primary-text-dark'>
            <Image
              src='/images/profile.png'
              alt='Profile'
              className='rounded-[100%]'
              width={100}
              height={100}
            />
            <button className='mt-2 w-full flex flex-col items-center gap-0'>
              <div className='flex flex-row items-center gap-2 text-lg font-medium'>
                Julio Caicedo
                <svg
                  className='hidden fill-current sm:block'
                  width='16'
                  height='16'
                  viewBox='0 0 16 16'
                  fill='none'
                >
                  <g clipPath='url(#clip0_34_896)'>
                    <path
                      d='M15.8174 4.13714C15.5736 3.89279 15.1779 3.89242 14.9335 4.13623L8.44194 10.6143C8.19822 10.858 7.80175 10.858 7.55759 10.6139L1.0665 4.13623C0.822157 3.89239 0.426469 3.89279 0.182594 4.13714C-0.0612184 4.38148 -0.0608122 4.7772 0.1835 5.02101L6.67416 11.4982C7.03981 11.8638 7.51997 12.0465 8.00019 12.0465C8.48013 12.0465 8.96016 11.8639 9.32534 11.4986L15.8165 5.02101C16.0608 4.7772 16.0612 4.38148 15.8174 4.13714Z'
                      fill='fill-current'
                    />
                  </g>
                  <defs>
                    <clipPath id='clip0_34_896'>
                      <rect width='16' height='16' fill='white' />
                    </clipPath>
                  </defs>
                </svg>
              </div>
              <span className='text-grey-5 dark:text-grey-dark-5'>Admin</span>
            </button>
          </div>

          <div className='flex flex-col overflow-y-auto duration-200 ease-linear'>
            <div className='flex flex-col'>
              {[
                { text: t('dashboard'), icon: 'icon-dashboard', href: '/' },
                {
                  text: t('balancer'),
                  icon: 'icon-balancer',
                  href: '/balancer',
                },
                { text: t('dhcp'), icon: 'icon-dhcp', href: '/dhcp' },
                {
                  text: t('firewall'),
                  icon: 'icon-firewall',
                  href: '/firewall',
                },
                { text: t('dns'), icon: 'icon-dns', href: '/dns' },
                { text: t('bgp'), icon: 'icon-bgp', href: '/bgp' },
              ].map((item, index) => {
                const isActive =
                  item.href === '/'
                    ? pathname === `/${lng}`
                    : pathname.includes(item.href);
                return (
                  <Fragment key={index}>
                    <Link
                      href={`/${lng}${item.href}`}
                      className='relative px-6 text-base text-grey-6 hover:text-primary-3 dark:text-grey-dark-6 hover:dark:text-primary-dark-2'
                    >
                      {isActive && <Active />}
                      <div
                        className={`${isActive ? 'bg-primary-2 font-medium hover:text-grey-6 dark:bg-primary-dark-2 hover:dark:text-grey-dark-6' : ''} relative z-10 flex w-full items-center gap-4 rounded-lg py-2 pl-4`.trim()}
                      >
                        <i className={`${item.icon}`} />
                        {item.text}
                      </div>
                    </Link>
                  </Fragment>
                );
              })}
            </div>
          </div>
        </aside>
      </div>
    </>
  );
};

export default Sidebar;
