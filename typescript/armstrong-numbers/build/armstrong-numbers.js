"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
class ArmstrongNumbers {
    isArmstrongNumber(num) {
        // sum(split.each ^ length)
        // An Armstrong number is a number that is the sum of its own digits each raised to the power of the number of digits.
        // For example:
        // 9 is an Armstrong number, because 9 = 9^1 = 9
        // 10 is not an Armstrong number, because 10 != 1^2 + 0^2 = 2
        // 153 is an Armstrong number, because: 153 = 1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153
        // 154 is not an Armstrong number, because: 154 != 1^3 + 5^3 + 4^3 = 1 + 125 + 64 = 190
        // let count = 
        // num
        let length = num.toString.length;
        Math.pow(num, length);
    }
}
exports.ArmstrongNumbers = ArmstrongNumbers;
//# sourceMappingURL=armstrong-numbers.js.map