import { useTranslation } from '@app/i18n';
import {
  ContentHeader,
  Action,
  ContentInfo,
  InfoItem,
} from '@components/layouts/dashboard';

const RootLayout = async ({
  children,
  params: { lng },
}: Readonly<{ children: React.ReactNode; params: { lng: string } }>) => {
  const { t } = await useTranslation(lng, 'dhcp');

  return (
    <>
      <ContentHeader
        imageUrl='/images/dhcp.svg'
        name={t('name')}
        status='running'
      >
        <Action>
          <i className='icon-play' />
        </Action>
        <Action>
          <i className='icon-pause' />
        </Action>
        <Action>
          <i className='icon-refresh' />
        </Action>
      </ContentHeader>

      <ContentInfo>
        <InfoItem title='Interface' text='eth0' />
        <InfoItem title='Network' text='192.168.200.0/24' />
        <InfoItem title='Gateway' text='192.168.200.1/24' />
        <InfoItem title='Time lease' text='14d:0h:0m:0s' />
        <InfoItem title='Range' text='192.168.200.2 - 192.168.200.254' />
        <InfoItem title='DNS' text='192.168.200.1,8.8.8.8' />
      </ContentInfo>

      {children}
    </>
  );
};

export default RootLayout;
