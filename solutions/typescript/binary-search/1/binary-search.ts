export function find(haystack: number[], needle: number): number | never {
    // Implement the binary search algorithm to find the index of the needle in the haystack
    //
    let left = 0;
    let right = haystack.length - 1;
    while (left <= right) {
        const mid = left + Math.floor((right - left) / 2);
        if (haystack[mid] === needle) {
            return mid;
        }
        if (haystack[mid] < needle) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }

    throw new Error('Value not in array');
}
