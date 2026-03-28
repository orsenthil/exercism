export function encode(inputByte: number[]): number[] {
  let result: number[] = [];
  for (let i = 0; i < inputByte.length; i++) {
    let byte = inputByte[i];
    if (byte === 0) {
      result.push(0);
      continue;
    }
    let bytes: number[] = [];
    while (byte > 0) {
      bytes.unshift(byte & 0x7f);
      byte >>= 7;
    }
    for (let j = 0; j < bytes.length - 1; j++) {
      bytes[j] |= 0x80;
    }
    result = result.concat(bytes);
  }

  return result;
}

export function decode(inputByte: number[]): number[] {
  let result: number[] = [];
  let byte: number = 0;
  for (let i = 0; i < inputByte.length; i++) {
    byte = (byte << 7) | (inputByte[i] & 0x7f);
    if ((inputByte[i] & 0x80) === 0) {
      result.push(byte);
      byte = 0;
    }
  }
  return result;

}
