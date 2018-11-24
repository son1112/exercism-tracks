// An Armstrong number is a number that is the sum of its own digits each raised to the power of the number of digits.

// For example:

// 9 is an Armstrong number, because 9 = 9^1 = 9
// 10 is not an Armstrong number, because 10 != 1^2 + 0^2 = 2
// 153 is an Armstrong number, because: 153 = 1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153
// 154 is not an Armstrong number, because: 154 != 1^3 + 5^3 + 4^3 = 1 + 125 + 64 = 190



export default class ArmstrongNumbers {
    isArmstrongNumber( num: number ) {

        let powers = [];

        chars.forEach(function(char) {
            let powered = Math.pow(num, length);
            powers.push(parseInt(powered));
        });

        return powers.reduce((a,b) => {
            return a + b
        }, 0);
    }

    private stringNum( num: number ) {
        return num.toString;
    }
    private chars( str: string ) {
        return str.split('');
    }
    private length( arr: array ) {
        return arr.length;
    }
}
