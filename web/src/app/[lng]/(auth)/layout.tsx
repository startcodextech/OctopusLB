import type { Metadata } from 'next';
import { dir } from 'i18next';
import { fallbackLng, languages } from '@app/i18n/settings';
import { useTranslation } from '@app/i18n';

import '../globals.css';

export async function generateStaticParams() {
  return languages.map((lng) => ({ lng }));
}

export async function generateMetadata({
  params: { lng },
}: {
  params: {
    lng: string;
  };
}): Promise<Metadata> {
  if (languages.indexOf(lng) < 0) lng = fallbackLng;
  // eslint-disable-next-line react-hooks/rules-of-hooks
  const { t } = await useTranslation(lng, 'login');
  return {
    title: t('title'),
    description: '',
  };
}

const RootLayout = ({
  children,
  params: { lng },
}: Readonly<{ children: React.ReactNode; params: { lng: string } }>) => {
  return (
    <>
      <html lang={lng} dir={dir(lng)}>
        <body className='m-0 h-screen overflow-x-hidden bg-[url(/images/bg.jpg)] !bg-cover bg-center !bg-no-repeat p-0'>
          {children}
        </body>
      </html>
    </>
  );
};

export default RootLayout;
