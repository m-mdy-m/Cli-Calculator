const { createInterface } = require("readline/promises");
function doTheMath(f, o, s) {
  switch (o) {
    case "*":
      output = parseFloat(f) * parseFloat(s);
      break;

    case "/":
      output = parseFloat(f) / parseFloat(s);
      break;
    case "+":
      output = parseFloat(f) + parseFloat(s);
      break;
    case "-":
      output = parseFloat(f) - parseFloat(s);
      break;
    default:
      output = "Invalid operation";
  }
  return output;
}
(async () => {
  const readline = createInterface({
    input: process.stdin,
    output: process.stdout,
  });
  const first_number = await readline.question("Enter first number:");
  const operation = await readline.question("Enter operation:(+, -, *, /):");
  const second_number = await readline.question("Enter second number:");
  const output = doTheMath(first_number, operation, second_number);
  console.log("Result:", output);
  readline.close();
})();
