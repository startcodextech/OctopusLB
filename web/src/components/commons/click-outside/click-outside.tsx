'use client';
import React, {FC, PropsWithChildren, RefObject, useEffect} from "react";

type Props = PropsWithChildren<{
    exceptionRef?: RefObject<HTMLDivElement>;
    onClick: () => void;
    className?: string;
}>;

const ClickOutside: FC<Props> = (props) => {
    const {exceptionRef, className, onClick, children} = props;

    const wrapperRef = React.useRef<HTMLDivElement>(null);

    useEffect(() => {
        const handleClickListener = (e: MouseEvent) => {
            let clickedInside : null | boolean = false;
            if (exceptionRef) {
                clickedInside =
                    (wrapperRef.current &&
                        wrapperRef.current.contains(e.target as Node)) ||
                    (exceptionRef.current && exceptionRef.current === e.target) ||
                    (exceptionRef.current &&
                        exceptionRef.current.contains(e.target as Node));
            } else {
                clickedInside =
                    wrapperRef.current &&
                    wrapperRef.current.contains(e.target as Node);
            }

            if (!clickedInside) {
                onClick();
            }
        }

        document.addEventListener('mousedown', handleClickListener);

        return () => {
            document.removeEventListener('mousedown', handleClickListener);
        }
    }, [exceptionRef, onClick]);


    return (
        <>
            <div ref={wrapperRef} className={`${className || ''}`}>
                {children}
            </div>
        </>
    )
};

export default ClickOutside;