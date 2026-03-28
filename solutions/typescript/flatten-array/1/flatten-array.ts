export function flatten(input: any[]): number[] {
  // recursive function to flatten the array
  const flattenArray = (arr: number[] | number[]): number[] => {
    return arr.reduce((acc: number[], val) => {
      if (Array.isArray(val)) {
        return acc.concat(flattenArray(val));
      } else {
        if (typeof val === 'number') {
          return acc.concat(val);
        } else {
          return acc;
        }
      }
    }, []);
  };


  return flattenArray(input);
  
}
