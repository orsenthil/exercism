const callbacks: (() => void)[] = [];
const computed: [() => number, number | undefined, boolean | undefined][] = [];
export const createInput = (init: number) => {
  callbacks.length = 0;
  computed.length = 0;
  let curr = init;
  return [
    () => curr,
    (next: number) => {
      const eqCompute = computed.find(([, , val]) => val);
      const predRes = () => eqCompute?.[0]?.();
      const currEq = predRes();
      curr = next;
      const nextEq = predRes(); 
      if (currEq && nextEq && currEq === nextEq) {
        return eqCompute![1];
      } else {
        return callbacks.reduce((acc, currPredicate) => {
          currPredicate();
          return acc
        }, curr);
      }
    }
  ] as const;
};

export const createComputed = (val: () => number, init?: number, eq?: boolean) => {
  computed.push([val, init, eq]);
  return val;
}

export const createCallback = (val: () => void) => {
  callbacks.push(val);
  return () => {
    const valPos = callbacks.indexOf(val);
    if (valPos !== -1) {
      callbacks.splice(valPos, 1);
    }
  };
};
