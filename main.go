package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func find(search string, arr []string) int {
	for i, h := range arr {
		if h == search {
			return i
		}
	}
	return -1
}

type dataframe struct {
	headers []string
	data    []interface{}
}

func (df dataframe) display() {
	fmt.Println(df.headers)
	for _, row := range df.data {
		fmt.Println(row)
	}
}

func (df *dataframe) append(newRow []interface{}) {
	df.data = append(df.data, newRow)
}

func (df *dataframe) pull(name string) series {
	searchIndex := find(name, df.headers)
	if searchIndex == -1 {
		panic("Could not find search")
	}
	s := series{}
	s.name = name
	for _, row := range df.data {
		s.values = append(s.values, row.([]interface{})[searchIndex])
	}
	return s
}

type series struct {
	name   string
	values []interface{}
}

func (s series) mean() float64 {
	var sum float64
	var length int = len(s.values)
	for _, val := range s.values {
		switch t := val.(type) {
		case float64:
			sum += t
		case int:
			sum += float64(t)
		default:
			panic("Strings and booleans are currently not supported for this function. Check your series values.")
		}
	}
	return sum / float64(length)
}

func (s series) sum() float64 {
	var sum float64
	for _, val := range s.values {
		switch t := val.(type) {
		case float64:
			sum += t
		case int:
			sum += float64(t)
		default:
			panic("Strings and booleans are currently not supported for this function. Check your series values.")
		}
	}
	return sum
}

func (s series) max() float64 {
	if len(s.values) == 0 {
		panic("Your series has no values in it.")
	}
	var max float64
	for i, val := range s.values {
		switch t := val.(type) {
		case float64:
			if i == 0 {
				max = t
			} else if t > max {
				max = t
			}
		case int:
			if i == 0 {
				max = float64(t)
			} else if float64(t) > max {
				max = float64(t)
			}
		default:
			panic("Strings and booleans are currently not supported for this function. Check your series values.")
		}
	}
	return max
}

func (s series) min() float64 {
	if len(s.values) == 0 {
		panic("Your series has no values in it.")
	}
	var min float64
	for i, val := range s.values {
		switch t := val.(type) {
		case float64:
			if i == 0 {
				min = t
			} else if t < min {
				min = t
			}
		case int:
			if i == 0 {
				min = float64(t)
			} else if float64(t) < min {
				min = float64(t)
			}
		default:
			panic("Strings and booleans are currently not supported for this function. Check your series values.")
		}
	}
	return min
}

func main() {
	df := dataframe{}
	df.headers = []string{"Col1", "col2", "col3"}
	row1 := []interface{}{12, 15, 17}
	df.data = append(df.data, row1)
	row2 := []interface{}{"hi", 100, "weird"}
	df.data = append(df.data, row2)
	df.append(row2)
	fmt.Println("Displaying df:")
	df.display()

	s := df.pull("col2")
	fmt.Println("Series:", s)

	avg := df.pull("col2").mean()
	fmt.Println("Mean:", avg)

	fmt.Println("Sum:", s.sum())
	fmt.Println("Max:", s.max())
	fmt.Println("Min:", s.min())

}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

// Convert takes lines read from csv file (which are rows) and
// converts them into columns as maps
func Convert(lines [][]string) map[string][]interface{} {
	headers := lines[0] // headers

	columns := make(map[string][]interface{})
	for i, line := range lines {
		if i == 0 {
			continue
		}
		for j, value := range line {
			floatVal, err := strconv.ParseFloat(value, 64)
			if err != nil {
				columns[headers[j]] = append(columns[headers[j]], value)
			}
			columns[headers[j]] = append(columns[headers[j]], floatVal)
		}
	}
	return columns
}
