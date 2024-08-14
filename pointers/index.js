// Primitive values are passed by value: When you pass a to the function int1, you're passing a copy of a. Any changes made to a inside the function don't affect the original a outside the function.

const int1 = (a) => {
  return a + 10;
};

let a = 1;

int1(a);

console.log("a:", a);
