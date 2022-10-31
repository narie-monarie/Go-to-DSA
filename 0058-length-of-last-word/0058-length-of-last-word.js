/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {
    var x = s.trim().split(" ")
    return x[x.length - 1].length;
};