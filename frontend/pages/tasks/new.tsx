import { NextPage } from 'next';
import React, { useState } from 'react';

import { Typography } from '@material-ui/core';

import AuthPage from '../../src/components/templates/AuthPage';

type P = {};

const NewTask: NextPage<P> = ({}) => {
  const [name, setName] = useState('');
  const [desc, setDesc] = useState('');

  return (
    <AuthPage title={<Typography variant="h3">New Task</Typography>}></AuthPage>
  );
};

export default NewTask;
