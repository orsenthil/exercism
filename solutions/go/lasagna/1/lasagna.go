package lasagna

// TODO: define the 'OvenTime' constant

var OvenTime = 40;

// TODO: define the 'RemainingOvenTime()' function

func RemainingOvenTime(minutes int) int {

    return OvenTime - minutes;
    
}
// TODO: define the 'PreparationTime()' function

func PreparationTime(layers int) int {
    return layers * 2;
    
}
// TODO: define the 'ElapsedTime()' function

func ElapsedTime(layers int, minutes int) int {

    return minutes + PreparationTime(layers);

}