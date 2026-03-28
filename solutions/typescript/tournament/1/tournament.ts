const DEFAULT_SCORES = {mp: 0, w: 0, d: 0, l: 0, p: 0}
type Score = typeof DEFAULT_SCORES;
const formatNumber = (number: number) => number.toString().padStart(2, ' ');

export class Tournament {
  // eslint-disable-next-line no-unused-vars
  public tally(input: string): string {
    const lines = ['Team                           | MP |  W |  D |  L |  P'];
    const inputs = input.split('\n').filter(Boolean).map(line => line.split(';'));
    const map: Record<string, Score> = {};
    for (const [first, second, result] of inputs) {
      map[first] ??= { ...DEFAULT_SCORES};
      map[first].mp += 1;
      map[second] ??= { ...DEFAULT_SCORES};
      map[second].mp += 1;

      if (result === 'draw') {
        map[first].d += 1;
        map[first].p += 1;  
        map[second].d += 1;
        map[second].p += 1;
      } else {
        const [winner, loser] = result === 'win' ? [first, second] : [second, first]; 
        map[winner].w += 1;
        map[winner].p += 3;
        map[loser].l += 1;
      }
    }

    const results = Object.entries(map).sort(([n1, s1], [n2, s2]) => 
    (s1.p != s2.p ? s2.p - s1.p : n1 < n2 ? -1 : 1));
    for (const [team, score] of results) {
      lines.push(
        [
          team.padEnd(30, ' '),
          formatNumber(score.mp),
          formatNumber(score.w),
          formatNumber(score.d),
          formatNumber(score.l),
          formatNumber(score.p),
        ].join(' | ')
      )
    }

    return lines.join('\n');
  }
}
