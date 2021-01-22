import { formatDistance } from 'date-fns';

export function datediff(next: Date, prev: Date): string {
  if (next < prev) [next, prev] = [prev, next];
  return formatDistance(prev, next);
}
