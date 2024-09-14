import Image from 'next/image';
import { LoginForm } from '@modules/auth/components';
import { version } from '../../../../../package.json';
import { useTranslation } from '@app/i18n';

export default async function Login({
  params: { lng },
}: {
  params: { lng: string };
}) {
  const { t } = await useTranslation(lng, 'login');

  return (
    <>
      <div className='container mx-auto w-full max-w-lg pt-8'>
        <div className='mx-4 rounded-3xl bg-[rgba(255,255,255,.7)] p-8 backdrop-blur-lg sm:mx-0'>
          <div className='flex flex-col items-center justify-center'>
            <Image
              src='/images/icon.png'
              alt='Logo'
              className='w-24 pb-4'
              width={96}
              height={96}
            />
            <h2 className='pb-3 text-center text-4xl font-extrabold'>
              OctopusLB
            </h2>
          </div>
          <h3 className='mb-4 text-center'>{t('subtitle')}</h3>

          <LoginForm lng={lng} />
        </div>
      </div>

      <div className='-px-3 flex flex-wrap py-5'>
        <div className='mx-auto w-full max-w-full text-center sm:w-3/4'>
          <p className='py-1 text-sm font-medium text-white'>
            OctopusLB &nbsp;
            <a
              href='https://github.com/startcodextech/OctopusLB'
              target='_blank'
            >
              v{version}
            </a>{' '}
            by &nbsp;
            <a href='https://startcodex.com' target='_blank'>
              Start Codex
            </a>{' '}
            © {new Date().getFullYear()}.
          </p>
        </div>
      </div>
    </>
  );
}
