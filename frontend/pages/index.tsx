import Link from 'next/link';
import React from 'react';

import { Grid, List, ListItem, Typography } from '@material-ui/core';

import AuthPage from '../src/components/templates/AuthPage';

export default function Home() {
  return (
    <AuthPage title={<Typography variant="h3">Dashboard</Typography>}>
      <List>
        <ListItem>
          <Link href="/tasks">Tasks</Link>
        </ListItem>
        <ListItem>
          <Link href="/projects">Projects</Link>
        </ListItem>
      </List>
    </AuthPage>
  );
}
