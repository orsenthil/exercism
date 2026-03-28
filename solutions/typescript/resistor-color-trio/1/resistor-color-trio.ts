export function decodedResistorValue(colors: string[]) {
    const value = colorCode(colors[0]) * 10 + colorCode(colors[1]);
    const exponent = colorCode(colors[2]);

    const answer = value * Math.pow(10, exponent);

    if (answer < 1000) {
        return `${answer} ohms`;
    }

    if (answer >= Math.pow(10, 3) && answer < Math.pow(10, 6)) {
        return `${answer / Math.pow(10, 3)} kiloohms`;
    }
    if (answer >= Math.pow(10, 6) && answer < Math.pow(10, 9)) {
        return `${answer / Math.pow(10, 6)} megaohms`;
    }
    if (answer >= Math.pow(10, 9)) {
        return `${answer / Math.pow(10, 9)} gigaohms`;
    }

}

export const colorCode = (color : string) => {
    return COLORS.indexOf(color);
}


export const COLORS = [
    'black',
    'brown',
    'red',
    'orange',
    'yellow',
    'green',
    'blue',
    'violet',
    'grey',
    'white',
]
