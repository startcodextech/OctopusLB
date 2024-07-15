'use client';
import React, {FC, LabelHTMLAttributes} from "react";

const Label: FC<LabelHTMLAttributes<HTMLLabelElement>> = (props) => {
    return (
        <>
            <label {...props} className="mb-2 inline-block text-base font-medium text-start text-grey-900"/>
        </>
    )
};

export default Label;