function isPrime(n: number): boolean {
  // A prime number is a number that is divisible only by itself and one
  // Return true if n is prime, false otherwise
  // This function is not optimized for speed
  for (let i = 2; i < n; i++) {
    if (n % i == 0) return false;
  }
  return true;
}
