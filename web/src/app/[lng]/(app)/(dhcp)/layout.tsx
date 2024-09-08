import Image from "next/image";
import {useTranslation} from "@app/i18n";

const RootLayout = async ({children, params: {lng}}: Readonly<{ children: React.ReactNode; params: {lng: string}}>) => {
    const { t } = await useTranslation(lng, "dhcp")

    return (
        <>
            <div className="flex items-center">
                <div className="flex items-center gap-2 w-full">
                    <Image src="/images/dhcp.png" alt="" width={44} height={44}/>
                    <span className="text-2xl text-[#080D30] font-bold">{t("name")}</span>
                    <div className="flex items-center gap-2 py-0 px-4 bg-[#CEEFDF] text-[#458850] text-sm leading-6 rounded-lg">
                        <div className="rounded-full w-2.5 h-2.5 bg-[#34C759]"/>
                        Starting
                    </div>
                </div>

                aaaa

            </div>
            {children}
        </>
    )
}

export default RootLayout;