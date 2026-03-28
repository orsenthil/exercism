export function transform(inputValue: { [key: string]: string[] }): { [key: string]: number } {
    let result: { [key: string]: number } = {}
    for (let key in inputValue) {
        inputValue[key].forEach((value) => {
            result[value.toLowerCase()] = parseInt(key)
        })
    }

    return result;
}
