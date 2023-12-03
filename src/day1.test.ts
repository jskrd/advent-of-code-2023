import { expect, test, describe } from "bun:test";
import { findFirstAndLastDigit, solvePart1, solvePart2 } from "./day1";

describe("solvePart1", () => {
  const input = "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet";

  test("example", () => {
    expect(solvePart1(input)).toBe(142);
  });
});

describe("solvePart2", () => {
  const input =
    "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen";

  test("example", () => {
    expect(solvePart2(input)).toBe(281);
  });
});

describe("findFirstAndLastDigit", () => {
  test.each([
    {
      line: "two1nine",
      expected: { first: 2, last: 9 },
    },
    {
      line: "eightwothree",
      expected: { first: 8, last: 3 },
    },
    {
      line: "abcone2threexyz",
      expected: { first: 1, last: 3 },
    },
    {
      line: "xtwone3four",
      expected: { first: 2, last: 4 },
    },
    {
      line: "4nineeightseven2",
      expected: { first: 4, last: 2 },
    },
    {
      line: "zoneight234",
      expected: { first: 1, last: 4 },
    },
    {
      line: "7pqrstsixteen",
      expected: { first: 7, last: 6 },
    },
    {
      line: "eighthree",
      expected: { first: 8, last: 3 },
    },
    {
      line: "sevenine",
      expected: { first: 7, last: 9 },
    },
  ])("%o", ({ line, expected }) => {
    expect(findFirstAndLastDigit(line)).toEqual(expected);
  });
});
