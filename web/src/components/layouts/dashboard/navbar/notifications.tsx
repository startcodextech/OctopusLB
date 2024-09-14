'use client';
import React, { FC, useState } from 'react';
import { ClickOutside } from '@components/commons';

const DropdownNotifications: FC = () => {
  const [dropdownOpen, setDropdownOpen] = useState(false);
  const [notifying, setNotifying] = useState(true);

  const handleDropdown = (value: boolean) => {
    setNotifying(false);
    setDropdownOpen(value);
  };

  return (
    <>
      <ClickOutside onClick={() => handleDropdown(false)} className='relative'>
        <li>
          <a
            href='#'
            onClick={() => handleDropdown(!dropdownOpen)}
            className='relative flex h-9 w-9 items-center justify-center rounded-full border-[0.5px] bg-gray-100 stroke-1 hover:text-primary'
          >
            <span
              className={`absolute -top-0.5 right-0 z-10 h-2 w-2 rounded-full bg-red-500 ${!notifying ? 'hidden' : 'inline'}`}
            >
              <span className='absolute -z-10 inline-flex h-full w-full animate-ping rounded-full bg-red-500 opacity-75'></span>
            </span>

            <svg fill='none' height='24' viewBox='0 0 24 24' width='24'>
              <g fill='rgb(0,0,0)'>
                <path d='m4.45549 13.88-.57156-.4856zm.87927-2.0415.74523.0844zm13.33044 0 .7453-.0844zm.8793 2.0415.5716-.4856zm-1.2241-5.08597-.7453.08443zm-12.64076 0-.74524-.08442zm12.49026 7.45597h-12.33976v1.5h12.33976zm-.5948-7.37154.3449 3.04444 1.4905-.1688-.3449-3.04449zm-11.49511 3.04444.34488-3.04444-1.49047-.16885-.34487 3.04449zm-1.05293 2.4428c.58466-.6882.95069-1.5402 1.05293-2.4428l-1.49046-.1688c-.06915.6104-.31633 1.1822-.7056 1.6403zm12.89294-2.4428c.1023.9026.4683 1.7546 1.0529 2.4428l1.1432-.9713c-.3893-.4581-.6365-1.0299-.7056-1.6403zm-12.08986 4.3271c-.8868 0-1.45088-1.1219-.80308-1.8843l-1.14313-.9713c-1.41902 1.6703-.30565 4.3556 1.94621 4.3556zm12.33976 1.5c2.2518 0 3.3652-2.6853 1.9462-4.3556l-1.1432.9713c.6478.7624.0838 1.8843-.803 1.8843zm.8957-9.04039c-.4152-3.66488-3.4377-6.45961-7.0656-6.45961v1.5c2.8302 0 5.2419 2.18698 5.5751 5.12846zm-12.64073.16885c.33321-2.94148 2.7449-5.12846 5.57513-5.12846v-1.5c-3.62788 0-6.65044 2.79473-7.0656 6.45961z'></path>
                <path d='m15.7023 19.2632c.1454-.3879-.0512-.8201-.4391-.9655s-.8201.0512-.9655.4391zm-6-.5264c-.14536-.3879-.57763-.5845-.9655-.4391s-.58446.5776-.4391.9655zm4.5954 0c-.3227.8611-1.2131 1.5132-2.2977 1.5132v1.5c1.6855 0 3.1516-1.0175 3.7023-2.4868zm-2.2977 1.5132c-1.0846 0-1.975-.6521-2.2977-1.5132l-1.4046.5264c.55065 1.4693 2.0168 2.4868 3.7023 2.4868z'></path>
              </g>
            </svg>
          </a>

          {dropdownOpen && (
            <>
              <div
                className={`absolute -right-[4rem] mt-2.5 flex h-[22.5rem] w-[18.75rem] flex-col rounded-sm border bg-white stroke-1 shadow sm:right-0 sm:w-[20rem]`}
              ></div>
            </>
          )}
        </li>
      </ClickOutside>
    </>
  );
};

export default DropdownNotifications;
