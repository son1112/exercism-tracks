function isDivisibleBy( year: number, div: number ) {
    return (year % div === 0)
}

function isLeapYear( year: number ) {
    return isDivisibleBy(year, 4)
        && (!isDivisibleBy(year, 100) || isDivisibleBy(year, 400));
}

export default isLeapYear
