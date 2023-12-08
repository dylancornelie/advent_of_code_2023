const fs = require("fs")

const file = fs.readFileSync("./input.txt", "utf-8")
const inputs = file.split("\n")

// const inputs = [
//         "two1nine",
//         "eightwothree",
//         "abcone2threexyz",
//         "xtwone3four",
//         "4nineeightseven2",
//         "zoneight234",
//         "7pqrstsixteen",
// ]

const regex = [
    new RegExp("one", "g"),      // -> Index 0
    new RegExp("two", "g"),      // -> Index 1
    new RegExp("three", "g"),    // -> Index 2
    new RegExp("four", "g"),     // -> Index 3
    new RegExp("five", "g"),     // -> Index 4
    new RegExp("six", "g"),      // -> Index 5
    new RegExp("seven", "g"),    // -> Index 6
    new RegExp("eight", "g"),    // -> Index 7
    new RegExp("nine", "g")      // -> Index 8
]

const solve = (input) => {
    const matchIndexes = regex.map((regex) => [...input.matchAll(regex)].map(match => match.index))
    let firstNumber = null
    let secondNumber = null
    const chars = input.split("")
    chars.forEach((_, index) => {

        let writtenNumberIndex = null
        matchIndexes.forEach((regexResult, regexIndex) => {
            if (regexResult.includes(index)) {
                writtenNumberIndex = regexIndex
            }
        })

        if (writtenNumberIndex !== null) {
            if (firstNumber == null) {
                firstNumber = writtenNumberIndex + 1
            } else {
                secondNumber = writtenNumberIndex + 1
            }
        } else {
            const number = parseInt(chars[index])
            if (!isNaN(number)) {
                if (firstNumber == null) {
                    firstNumber = number
                } else {
                    secondNumber = number
                }
            }
        }
    })
    secondNumber == null ? secondNumber = firstNumber : {}
    return parseInt(`${firstNumber}${secondNumber}`)
}

const solution = inputs.reduce((acc, input) => acc + solve(input), 0)

console.log(solution);
