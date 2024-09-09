'use client';

import React, {FC, PropsWithChildren} from 'react';

type Props = PropsWithChildren<{

}>;

const HeaderInfo: FC<Props> = (props) => {
    const {children} = props;

    return (
        <>
            <div className="flex flex-row py-8 px-0 gap-4 items-center">
                <div className="w-full flex flex-wrap justify-between items-center content-center gap-y-4">
                    {children}
                </div>
                <button className="flex items-center justify-between max-w-11 max-h-11 text-3xl p-2.5 rounded-[0.5rem] hover:bg-neutral-background">
                    <i className="icon-setting" />
                </button>
            </div>
        </>
    )
};

export default HeaderInfo;