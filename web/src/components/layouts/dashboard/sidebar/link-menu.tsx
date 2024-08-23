'use client';
import React, {FC, PropsWithChildren, ReactNode, useState} from "react";

type Props = PropsWithChildren<{
    children: (handleClick: () => void, open: boolean) => ReactNode;
    active?: boolean;
}>;

const LinkItemMenu: FC<Props> = (props) => {
    const {children, active = false} = props;

    const [open, setOpen] = useState(active);

    const handleClick = () => {
        setOpen(!open);
    }

    return (
        <>
            <li>{children(handleClick, open)}</li>
        </>
    )
}

export default LinkItemMenu;