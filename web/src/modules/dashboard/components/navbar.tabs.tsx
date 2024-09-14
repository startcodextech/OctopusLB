'use client';
import React, { FC, useEffect, useCallback, useContext } from 'react';
import { Context } from '@components/layouts/dashboard/dashboard';

const Tabs = () => {
  const { setTabNavbar } = useContext(Context);

  useEffect(() => {
    const tabs = <>hola mundo</>;

    setTabNavbar(tabs);
  }, [setTabNavbar]);

  return <></>
};

export default Tabs;
