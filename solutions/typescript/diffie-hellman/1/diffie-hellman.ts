export class DiffieHellman {
  p: number
  g: number

  constructor(p: number, g: number) {
    if (p < 2 || g < 2) throw new Error('p and g must be greater than 1')
    if (!this.isPrime(p) || !this.isPrime(g)) throw new Error('p and g must be prime')

    this.p = p
    this.g = g

  }

  public getPublicKey(privateKey: number): number{
    if (privateKey < 2 || privateKey >= this.p) throw new Error('privateKey is invalid')
    return Math.pow(this.g, privateKey) % this.p
  }

  public getSecret(theirPublicKey: number, myPrivateKey: number): number {
    if (myPrivateKey < 2 || myPrivateKey >= this.p) throw new Error('myPrivateKey is invalid')
    return Math.pow(theirPublicKey, myPrivateKey) % this.p
  }

  private isPrime(n: number): boolean {
    for (let i = 2; i <= Math.sqrt(n); i++) {
      if (n % i === 0) return false
    }
    return true
  }
}
