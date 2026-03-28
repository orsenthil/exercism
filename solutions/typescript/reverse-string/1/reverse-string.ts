export function reverse(str: string) {
    // Reverse the string

    str = str.split('').reverse().join('');
    return str;
}

