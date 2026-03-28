package lasagna

var OvenTime = 40;

func RemainingOvenTime(minutes int) int {
    return OvenTime - minutes;
}

func PreparationTime(layers int) int {
    return layers * 2;
}

func ElapsedTime(layers int, minutes int) int {
    return minutes + PreparationTime(layers);
}