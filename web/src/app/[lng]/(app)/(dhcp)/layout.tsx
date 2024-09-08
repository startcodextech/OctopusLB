import Image from "next/image";
import {useTranslation} from "@app/i18n";
import {Chip} from "@components/commons";

const RootLayout = async ({children, params: {lng}}: Readonly<{ children: React.ReactNode; params: {lng: string}}>) => {
    const { t } = await useTranslation(lng, "dhcp")

    return (
        <>
            <div className="flex items-center">
                <div className="flex items-center gap-2 w-full text-text-primary">
                    <Image src="/images/dhcp.svg" alt="" width={44} height={44}/>
                    <span className="text-2xl text-text-primary font-bold">{t("name")}</span>
                    <Chip type="error" showDot={true}>Starting</Chip>
                </div>

                <div className="flex items-center flex-row gap-2">
                    <button
                        className="w-11 h-11 flex items-center justify-center p-2.5 rounded-[0.5rem] border-2 bg-white border-neutral-background hover:bg-neutral-background">
                        <i className="icon-play text-2xl"/>
                    </button>
                    <button
                        className="w-11 h-11 flex items-center justify-center p-2.5 rounded-[0.5rem] border-2 bg-white border-neutral-background hover:bg-neutral-background">
                        <i className="icon-pause text-2xl"/>
                    </button>
                    <button
                        className="w-11 h-11 flex items-center justify-center p-2.5 rounded-[0.5rem] border-2 bg-white border-neutral-background hover:bg-neutral-background">
                        <i className="icon-refresh text-2xl"/>
                    </button>
                </div>

            </div>
            {children}
        </>
    )
}

export default RootLayout;