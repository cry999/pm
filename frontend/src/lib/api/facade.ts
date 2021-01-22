import { useContext } from 'react';

import { BaseAPI, Configuration } from '../../../generated/api';
import AuthContext from '../../contexts/auth';
import { Constructor } from '../type';

export function facade<T extends BaseAPI>(ctor: Constructor<T>): T {
  const {
    auth: { accessToken },
  } = useContext(AuthContext);
  const conf = new Configuration({ accessToken });
  const api = new ctor(conf);
  return api;
}
