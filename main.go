package main


import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
    "math"
)

func readVals(reader *bufio.Reader) []float64 {
    input, _ := reader.ReadString('\n')

    tokens := strings.Fields(input)
    result := make([]float64, len(tokens))

    for i := range tokens {
        fVal, err := strconv.ParseFloat(tokens[i], 64);
        if err != nil {
            fmt.Fprintf(os.Stderr, "Can not convert %s to float64.", tokens[i])
            os.Exit(1)
        }

        result[i] = fVal
    }

    return result
}

func fit(xVals []float64, yVals []float64) (float64, float64) {
    var a float64
    var b float64

    var sumY float64 = 0
    var sumX float64 = 0
    var dotXY float64 = 0
    var dotXX float64 = 0

    for i := range xVals {
        xVal := xVals[i]
        yVal := yVals[i]

        sumX += xVal
        sumY += yVal

        dotXY += xVal * yVal
        dotXX += math.Pow(xVal, 2)
    }

    n := float64(len(xVals))
    a = (n * dotXY - sumX * sumY) /
        (n * dotXX - math.Pow(sumX, 2))
    b = (sumY - a * sumX) / n

    return a, b
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Enter x values:")
    xVals := readVals(reader)

    fmt.Println("Enter y values:")
    yVals := readVals(reader)

    if len(xVals) != len(yVals) {
        fmt.Fprintf(os.Stderr, "Input vector lengths don't match.")
        os.Exit(2)
    }

    a, b := fit(xVals, yVals)

    fmt.Fprintf(os.Stdout, "Result line: f(x) = %fx + %f\n", a, b)
}
