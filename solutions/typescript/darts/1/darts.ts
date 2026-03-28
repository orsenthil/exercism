export function score(x: number, y: number): number {
    // Outer circle has radius 10
    // Middle circle has radius 5
    // Inner circle has radius 1
    var distance: number = Math.sqrt(x * x + y * y);
    if (distance <= 1) {
        return 10;
    } else if (distance <= 5) {
        return 5;
    } else if (distance <= 10) {
        return 1;
    }
    return 0;
}
