import Link from 'next/link';
import { useRouter } from 'next/router';
import React, { useContext, useEffect, useState } from 'react';

import {
  Breadcrumbs,
  createStyles,
  Divider,
  Grid,
  Link as MaterialLink,
  makeStyles,
  Paper,
  Typography,
} from '@material-ui/core';
import CachedIcon from '@material-ui/icons/Cached';

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

  const paths = router.asPath.split('/').filter((p) => p !== '');
  const breadcrumbs = paths.reduce(
    (prev, path) => {
      const last = prev[prev.length - 1];
      return prev.concat({
        path: last.path + '/' + path,
        name: path,
      });
    },
    [{ name: 'Home', path: '' }]
  );

  useEffect(() => {
    if (!auth.accessToken) {
      router.push('/signin');
    } else if (loading) {
      setLoading(false);
    }
  }, [loading]);

  return (
    <div>
      <header className={classNames.header}>
        <Typography variant="h4">PM</Typography>
      </header>
      <main role="main">
        <Paper className={classNames.main}>
          <Grid container className={classNames.title}>
            <Grid item xs={12}>
              {title}
            </Grid>
            <Grid item xs={12}>
              <Breadcrumbs aria-label="breadcrumb">
                {breadcrumbs.map((b, i) => {
                  return i !== breadcrumbs.length - 1 ? (
                    <Link key={'prv' + i} href={b.path || '/'} passHref>
                      <MaterialLink href={b.path || '/'} color="inherit">
                        {b.name}
                      </MaterialLink>
                    </Link>
                  ) : (
                    <Typography key={'cur' + i} color="textPrimary">
                      {b.name}
                    </Typography>
                  );
                })}
              </Breadcrumbs>
            </Grid>
          </Grid>
          <Divider />
          <Grid className={classNames.content}>
            {loading ? <CachedIcon /> : children}
          </Grid>
        </Paper>
      </main>
    </div>
  );
};

export default AuthPage;
