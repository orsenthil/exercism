package thefarm

import (
    "errors"
    "fmt"
) 

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.

type SillyNephewError struct {
    count int
}

func (e *SillyNephewError) Error() string {
    return fmt.Sprintf("silly nephew, there cannot be %d cows", e.count)
}


// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
    fodder, err := weightFodder.FodderAmount()

    if cows == 0 {
        	return 0.0, errors.New("division by zero")
	}

	if cows < 0 {
        return 0.0, &SillyNephewError{count: cows}
    }

    if fodder < 0 {
        if err != nil && err != ErrScaleMalfunction {
            return 0.0, err
        }
        return 0.0, errors.New("negative fodder")
    }

    if err == nil {
        return float64(fodder) / float64(cows), nil
    } 

    if err == ErrScaleMalfunction {
        return float64(2 * fodder)/ float64(cows), nil
    }

    if err != nil {
        return 0.0, err
    }

    return 0.0, nil
}
