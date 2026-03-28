export function decodedValue(colors : string[]) {
    return colorCode(colors[0]) * 10 + colorCode(colors[1]);
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
