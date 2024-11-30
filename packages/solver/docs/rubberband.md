# Rubberband

Rubberband is a syntax for structured sentences that enables low-code data input while maintaining precise control over input definitions, options, and associated values. It allows developers to define user-facing interfaces for data entry that map to specific intents.

## Core Concepts

### Input Placeholders

Input placeholders are denoted by curly braces containing an index number, starting from 0:

```txt
{0}, {1}, {2}, etc.
```

## Usage Patterns

### Raw Substitution

For independent inputs that don't share data relationships, use simple indexed placeholders:

```txt
Deposit {0} {1} into {2}.
```

### Pointers (Data Dependencies)

When subsequent inputs depend on previous selections, use the pointer syntax with =>:

```txt
Deposit {0} {1} into {1=>2}.
```

In this example, `{1=>2}` indicates that the options for input 2 are determined by the value selected in input 1. This helps filter irrelevant options and present only contextually appropriate choices.

### Delimiters

For cases requiring additional processing information or compound values, use the metadata break syntax:

```txt
Transfer {0} {1|:}
```

The `|` is a `metadata break` while the following value is the `delimiter` that will be used to separate the provided value. In this example, we would be passing an address to `1` with an appended delimiter so that we can define the `token standard`.
