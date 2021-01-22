export function abbrev(s: string, len: number): string {
  if (s.length <= len) return s;
  return s.substr(0, len - 3) + '...';
}
