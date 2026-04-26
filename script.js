function isPowerOfThree(n)  {
	for (i = 1; i < n ; i++) {
		result = i*i*i;
		if( result === n) {
			return true
		}
		if (result > n) {
			return false
		}
	}
	return false;
}

console.log(isPowerOfThree(27));
console.log(isPowerOfThree(0));
console.log(isPowerOfThree(-1));