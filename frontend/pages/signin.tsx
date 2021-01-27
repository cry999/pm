import { useRouter } from 'next/dist/client/router';
import React, { useContext, useState } from 'react';

import { Button, TextField } from '@material-ui/core';

import AuthContext from '../src/contexts/auth';
import { OAuthApi } from '../src/lib/api';

const SignIn = () => {
  const router = useRouter();
  const auth = useContext(AuthContext);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  if (auth.auth.accessToken) {
    router.push('/');
  }

  const onSignIn = async () => {
    try {
      const api = new OAuthApi();
      const res = await api.signin({ credential: { email, password } });
      auth.authenticate(res.accessToken);
    } catch (e) {
      alert(e);
    }
  };

  return (
    <div className="container">
      <div>
        <div>
          <TextField
            label="Email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>
        <div>
          <TextField
            label="Password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <div>
          <Button color="primary" onClick={onSignIn}>
            SIGN IN
          </Button>
        </div>
      </div>
    </div>
  );
};

export default SignIn;
