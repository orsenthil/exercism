package piecingittogether

import (
	"errors"
	"math"
)

type PuzzleDetails struct {
	Pieces      int
	Border      int
	Inside      int
	Rows        int
	Columns     int
	AspectRatio float64
	Format      string
}

func JigsawData(details PuzzleDetails) (PuzzleDetails, error) {
	rows, cols, ok := resolveRowsAndColumns(details)
	if !ok {
		return PuzzleDetails{}, errors.New("Insufficient data")
	}
	details.Rows = rows
	details.Columns = cols
	result := fillFromRowsAndColumns(details)

	if details.Pieces != 0 && details.Pieces != result.Pieces {
		return PuzzleDetails{}, errors.New("Contradictory data")
	}
	if details.Border != 0 && details.Border != result.Border {
		return PuzzleDetails{}, errors.New("Contradictory data")
	}
	if details.Inside != 0 && details.Inside != result.Inside {
		return PuzzleDetails{}, errors.New("Contradictory data")
	}
	if details.AspectRatio != 0 && math.Abs(details.AspectRatio - result.AspectRatio) > 1e-9 {
		return PuzzleDetails{}, errors.New("Contradictory data")
	}
	if details.Format != "" && details.Format != result.Format {
		return PuzzleDetails{}, errors.New("Contradictory data")
	}
	return result, nil
}

func fillFromRowsAndColumns(d PuzzleDetails) PuzzleDetails {
	pieces := d.Rows * d.Columns
	inside := (d.Rows - 2) * (d.Columns - 2)
	border := pieces - inside
	rows := d.Rows
	columns := d.Columns
	aspectRatio := float64(columns) / float64(rows)
	format := "square"
	if aspectRatio < 1 {
		format = "portrait"
	} else if aspectRatio > 1 {
		format = "landscape"
	}
	return PuzzleDetails{
		Pieces:      pieces,
		Border:      border,
		Inside:      inside,
		Rows:        rows,
		Columns:     columns,
		AspectRatio: aspectRatio,
		Format:      format,
	}
}

func resolveRowsAndColumns(d PuzzleDetails) (rows, cols int, ok bool) {
	// Case 1: Both rows and columns are provided
	if d.Rows != 0 && d.Columns != 0 {
		return d.Rows, d.Columns, true
	}
	// Case 2: Pieces and aspect ratio are provided
	if d.Pieces != 0 && d.AspectRatio != 0 {
		rows = int(math.Round(math.Sqrt(float64(d.Pieces) / d.AspectRatio)))
		cols = int(float64(d.Pieces) / float64(rows))
		return rows, cols, true
	}
	// Case 3: Rows and Aspect Ratio are provided
	if d.Rows != 0 && d.AspectRatio != 0 {
		cols = int(float64(d.Rows) * d.AspectRatio)
		return d.Rows, cols, true
	}
	// Case 4: Columns and Aspect Ratio are provided
	if d.Columns != 0 && d.AspectRatio != 0 {
		rows = int(float64(d.Columns) / d.AspectRatio)
		return rows, d.Columns, true
	}
	// Case 5: Rows and Format Square
	if d.Rows != 0 && d.Format == "square" {
		return d.Rows, d.Rows, true
	}
	// Case 6: Columns and Format Square
	if d.Columns != 0 && d.Format == "square" {
		return d.Columns, d.Columns, true
	}
	// Case 7: Inside and Aspect Ratio are provided
	if d.Inside != 0 && d.AspectRatio != 0 {
		rows = int(math.Round(math.Sqrt(float64(d.Inside) / d.AspectRatio)))
		cols = int(math.Round(math.Sqrt(float64(d.Inside) * d.AspectRatio)))
		return rows + 2, cols + 2, true
	}
	// Case 8: Pieces and Border are provided
	if d.Pieces != 0 && d.Border != 0 {
		inside := d.Pieces - d.Border
		rowsPlusCols := (inside - 4 - d.Pieces) / -2
		discriminant := float64(rowsPlusCols*rowsPlusCols - 4*d.Pieces)
		x1 := (float64(rowsPlusCols) + math.Sqrt(discriminant)) / 2
		x2 := (float64(rowsPlusCols) - math.Sqrt(discriminant)) / 2
		var rows, cols int
		if d.Format == "square" {
			rows = int(x1)
			cols = int(x1)
		} else if d.Format == "portrait" {
			if x1 > x2 {
				rows = int(x1)
				cols = int(x2)
			} else {
				rows = int(x2)
				cols = int(x1)
			}
		} else if d.Format == "landscape" {
			if x1 > x2 {
				rows = int(x2)
				cols = int(x1)
			} else {
				rows = int(x1)
				cols = int(x2)
			}
		}
		return rows, cols, true
	}
	return 0, 0, false

}