import { file } from "bun";

export default async function run(): Promise<void> {
  const input = await file("inputs/day1.txt").text();

  console.log("Day 1: Trebuchet?!");
  console.log(`- Part 1: ${solvePart1(input)}`);
  console.log(`- Part 2: ${solvePart2(input)}`);
}

export function solvePart1(input: string): number {
  return input
    .split("\n")
    .map((line: string): string[] =>
      line
        .split("")
        .filter((character: string): boolean => /^\d+$/.test(character))
    )
    .map((numbers: string[]): number =>
      parseInt(`${numbers[0]}${numbers[numbers.length - 1]}`)
    )
    .reduce((sum: number, number: number): number => sum + number);
}

export function solvePart2(input: string): number {
  return input
    .split("\n")
    .map((line: string): number => {
      const { first, last } = findFirstAndLastDigit(line);
      return parseInt(`${first}${last}`);
    })
    .reduce((sum: number, number: number): number => sum + number);
}

export function findFirstAndLastDigit(line: string): {
  first: number;
  last: number;
} {
  const digits: { [key: string]: string } = {
    one: "1",
    two: "2",
    three: "3",
    four: "4",
    five: "5",
    six: "6",
    seven: "7",
    eight: "8",
    nine: "9",
  };

  const reversedLine = line.split("").reverse().join("");

  const searchPatterns = [...Object.keys(digits), ...Object.values(digits)];
  const reversedSearchPatterns = searchPatterns.map((pattern: string): string =>
    pattern.split("").reverse().join("")
  );

  const match = line.match(new RegExp(searchPatterns.join("|")));
  const reversedMatch = reversedLine.match(
    new RegExp(reversedSearchPatterns.join("|"))
  );

  if (!match) {
    throw new Error("No first digit found");
  }

  if (!reversedMatch) {
    throw new Error("No last digit found");
  }

  let first = match[0];
  if (!/^\d+$/.test(first)) {
    first = digits[first];
  }

  let last = reversedMatch[0].split("").reverse().join("");
  if (!/^\d+$/.test(last)) {
    last = digits[last];
  }

  return { first: parseInt(first), last: parseInt(last) };
}
