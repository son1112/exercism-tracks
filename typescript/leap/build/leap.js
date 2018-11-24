"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
function isDivisibleBy(year, div) {
    return (year % div === 0);
}
function isLeapYear(year) {
    return isDivisibleBy(year, 4)
        && (!isDivisibleBy(year, 100) || isDivisibleBy(year, 400));
}
exports.default = isLeapYear;
//# sourceMappingURL=leap.js.map