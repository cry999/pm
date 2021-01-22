import { Router, useRouter } from 'next/dist/client/router';
import React, { useContext, useEffect, useState } from 'react';

import {
  createStyles,
  Divider,
  Grid,
  IconButton,
  makeStyles,
  Paper,
  Typography,
} from '@material-ui/core';

import AuthContext from '../../../contexts/auth';

const useStyles = makeStyles((theme) =>
  createStyles({
    header: {
      background: theme.palette.primary.light,
      height: '10vh',
    },
    main: {},
    title: {
      paddingTop: theme.spacing(1),
      paddingBottom: theme.spacing(2),
      paddingLeft: theme.spacing(5),
      paddingRight: theme.spacing(5),
      height: '15vh',
      textAlign: 'center',
    },
    content: {
      paddingTop: theme.spacing(2),
      paddingBottom: theme.spacing(2),
      paddingLeft: theme.spacing(6),
      paddingRight: theme.spacing(6),
      height: '75vh',
      overflow: 'scroll',
    },
  })
);

type P = {
  title: React.ReactNode;
};

const AuthPage: React.FC<P> = ({ title, children }) => {
  const classNames = useStyles();
  const { auth } = useContext(AuthContext);
  const router = useRouter();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!auth.accessToken) {
      router.push('/signin');
    } else if (loading) {
      setLoading(false);
    }
  }, [loading]);
  return (
    <div>
      {loading ? (
        <Typography variant="h1">Loading...</Typography>
      ) : (
        <div>
          <header className={classNames.header}>
            <Typography variant="h4">PM</Typography>
          </header>
          <main role="main">
            <Paper className={classNames.main}>
              <Grid container className={classNames.title}>
                {title}
              </Grid>
              <Divider />
              <Grid container className={classNames.content}>
                {children}
              </Grid>
            </Paper>
          </main>
        </div>
      )}
    </div>
  );
};

export default AuthPage;
