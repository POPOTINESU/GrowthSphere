module.exports = {
  extends: ["@commitlint/config-conventional"],
  rules: {
    "type-enum": [
      2,
      "always",
      ["feat", "fix", "docs", "style", "perf", "chore", "refactor", "test"],
    ],
    "subject-case": [2, "always", "lower-case"],
  },
};
