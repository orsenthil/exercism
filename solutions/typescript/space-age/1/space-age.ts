export function age(planet: unknown, seconds: unknown): unknown {
    if (planet == "earth") {
        return Number((Number(seconds) / 31557600).toFixed(2));
    } else if (planet == "mercury") {
        return Number((Number(seconds) / 31557600 / 0.2408467).toFixed(2));
    } else if (planet == "venus") {
        return Number((Number(seconds) / 31557600 / 0.61519726).toFixed(2));
    } else if (planet == "mars") {
        return Number((Number(seconds) / 31557600 / 1.8808158).toFixed(2));
    } else if (planet == "jupiter") {
        return Number((Number(seconds) / 31557600 / 11.862615).toFixed(2));
    } else if (planet == "saturn") {
        return Number((Number(seconds) / 31557600 / 29.447498).toFixed(2));
    } else if (planet == "uranus") {
        return Number((Number(seconds) / 31557600 / 84.016846).toFixed(2));
    } else if (planet == "neptune") {
        return Number((Number(seconds) / 31557600 / 164.79132).toFixed(2));
    } else {
        throw new Error("Invalid planet");
    }
}
