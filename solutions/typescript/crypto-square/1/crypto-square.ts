export class Crypto {
  constructor(private plainText: string) {
      this.plainText = plainText.replace(/\W/gi, '').toLowerCase();
  }

  get ciphertext(): string{
      const size = Math.sqrt(this.plainText.length);
      const c = Math.ceil(size);
      const r = Math.round(size);

      const result = [];

      for (let i = 0; i < c; i++) {
          const row = [];
          for (let j = 0; j < r; j++) {
              row.push(this.plainText[i + j * c] || ' ');
          }
          result.push(row.join(''));
      }
      return result.join(' ');
  }
}
