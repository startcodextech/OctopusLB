'use client';
import React, {FC, ReactNode, useState} from "react";

type Props = {
    children: (handleClick: () => void, open: boolean) => ReactNode;
    active?: boolean;
};

const LinkGroup: FC<Props> = (props) => {
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

export default LinkGroup;