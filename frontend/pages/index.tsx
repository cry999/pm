import Link from 'next/link';
import React, { useContext } from 'react';

import { Button, Typography } from '@material-ui/core';

import AuthContext from '../src/contexts/auth';

export default function Home() {
  const auth = useContext(AuthContext);
  const authenticated = !!auth.auth.accessToken;

  return (
    <div>
      {!authenticated ? (
        <div>
          Please{' '}
          <Link href="/signin" passHref>
            <Button>Sign In</Button>
          </Link>
        </div>
      ) : (
        <div>
          <header>
            <div>
              <Typography variant="h1">Top Page</Typography>
            </div>
            <div>
              <Link href="/tasks" passHref>
                <Button>Tasks</Button>
              </Link>
            </div>
          </header>
        </div>
      )}
    </div>
  );
}
