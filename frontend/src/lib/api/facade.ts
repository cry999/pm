import { useContext } from 'react';

import AuthContext from '../../contexts/auth';
import { Constructor } from '../type';
import { BaseAPI, Configuration } from './generated';

export function useFacade<T extends BaseAPI>(ctor: Constructor<T>): T {
  const {
    auth: { accessToken },
  } = useContext(AuthContext);
  const conf = new Configuration({ accessToken });
  const api = new ctor(conf);
  return api;
}
