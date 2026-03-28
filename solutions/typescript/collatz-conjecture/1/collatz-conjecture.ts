export function steps(count: number): number {
    if (!Number.isInteger(count) || count <= 0) {
        throw new Error('Only positive integers are allowed')
    }
    let steps = 0

    while (count > 1) {
        count = count % 2 === 0 ? count / 2 : count * 3 + 1
        steps++
    }

    return steps
}
