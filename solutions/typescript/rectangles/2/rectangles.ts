export function count(data: string[]): number {
  let rows = data.length;
  let cols = 0;
  let rectangles = 0;
  if( rows ) {
    cols = data[0].length;
  }

  for( let y = 0; y < rows; y++) {
    for( let x = 0; x < cols; x++) {
      for( let y2 = y + 1; y2 < rows; y2++) {
        for( let x2 = x + 1; x2 < cols; x2++) {

          if( data[y][x] ===  "+" && data[y2][x2] === "+" && data[y][x2] === "+" && data[y2][x] === "+") {
            let ok = true;
            
            for(let i = x + 1; i < x2; i++) {
              if( data[y][i] !== "-" && data[y][i]!=="+") {
                ok = false;
              }
              if( data[y2][i] !== "-" && data[y2][i]!=="+") {
                ok = false;          
              }            
            }
            for(let i = y + 1; i < y2; i++) {
              if( data[i][x] !== "|" && data[i][x]!=="+") {
                ok = false;
              }
              if( data[i][x2] !== "|" && data[i][x2]!=="+") {
                ok = false;
              }              
            }
            if( ok ) {
              rectangles++;
            }
          }
        }
      }
    }
  }

  return rectangles;
}