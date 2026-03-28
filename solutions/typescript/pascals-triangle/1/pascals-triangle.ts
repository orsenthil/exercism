export class Triangle {
    rows: number[][] = []

    constructor(private readonly numrows: number) {
        for (let i = 0; i < numrows; i++) {
            this.rows.push(this.generateRow(i))
        }
    }

    get lastRow(): number[] {
        return this.rows[this.numrows - 1]
    }

    generateRow(rownum: number): number[] {
        let row: number[] = []
        for (let i = 0; i <= rownum; i++) {
            row.push(this.binomial(rownum, i))
        }
        return row
    }

    binomial(n: number, k: number): number {
        return this.factorial(n) / (this.factorial(k) * this.factorial(n - k))
    }

    factorial(n: number): number {
        if (n === 0) {
            return 1
        }
        return n * this.factorial(n - 1)
    }
}
