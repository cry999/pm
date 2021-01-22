import { ApiError } from 'next/dist/next-server/server/api-utils';

import { TaskApi } from '../api';

export type Constructor<T> = { new (...args: any[]): T };

export type ReturnTypeAsync<
  T extends (...args: any[]) => any
> = ReturnType<T> extends Promise<infer P> ? P : never;

export type KeyType<T, K extends keyof T> = T[K];

export type ArrayElementType<T> = T extends (infer P)[] ? P : never;

export type ResultsItemType<
  T extends (...args: any[]) => any
> = ReturnTypeAsync<T> extends infer P
  ? 'results' extends keyof P
    ? ArrayElementType<KeyType<P, 'results'>>
    : never
  : never;
