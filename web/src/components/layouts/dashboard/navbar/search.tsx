import React, { FC, useContext } from 'react';
import { useTranslation } from '@app/i18n/client';
import { Context } from '@components/layouts/dashboard/dashboard';

type Props = {};

const Search: FC<Props> = () => {
  const { lng } = useContext(Context);
  const { t } = useTranslation(lng);
  return (
    <>
      <div className='hidden sm:block'>
        <div className='relative'>
          <button className='absolute left-3 top-1/2 -translate-y-1/2 p-2.5'>
            <svg
              width={16}
              height={16}
              viewBox='0 0 56.966 56.966'
              className='fill-current text-grey-5 focus:text-black dark:text-grey-dark-5 focus:dark:text-white'
              fill='none'
            >
              <path
                fill='fill-current'
                d='M55.146,51.887L41.588,37.786c3.486-4.144,5.396-9.358,5.396-14.786c0-12.682-10.318-23-23-23s-23,10.318-23,23
	s10.318,23,23,23c4.761,0,9.298-1.436,13.177-4.162l13.661,14.208c0.571,0.593,1.339,0.92,2.162,0.92
	c0.779,0,1.518-0.297,2.079-0.837C56.255,54.982,56.293,53.08,55.146,51.887z M23.984,6c9.374,0,17,7.626,17,17s-7.626,17-17,17
	s-17-7.626-17-17S14.61,6,23.984,6z'
              ></path>
            </svg>
          </button>
          <input
            type='search'
            placeholder={t('search')}
            className='border-1 box-border min-w-[295px] rounded-xl border-grey-2 bg-transparent pl-14 placeholder-grey-5 shadow-none !outline-none focus:border-grey-2 focus:ring-grey-2 dark:border-grey-dark-2 dark:placeholder-grey-dark-5 focus:dark:border-grey-dark-2 focus:dark:ring-grey-dark-2'
          />
        </div>
      </div>
    </>
  );
};

export default Search;
