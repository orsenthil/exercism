export function toRna(dna: string) {
    var answer: string = ''
    for (let i = 0; i < dna.length; i++) {
        if (dna[i] != 'A' && dna[i] != 'C' && dna[i] != 'G' && dna[i] != 'T') {
            throw new Error("Invalid input DNA.");
        }
        if (dna[i] == 'A') {
            answer += 'U'
        }
        if (dna[i] == 'C') {
            answer += 'G'
        }
        if (dna[i] == 'G') {
            answer += 'C'
        }
        if (dna[i] == 'T') {
            answer += 'A'
        }
    }
    return answer
}
