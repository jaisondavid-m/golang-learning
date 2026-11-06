# 04 - Functions in Go

## Overview

Learned how to create reusable functions, pass parameters, return values, handle multiple return values, and return errors in Go.

## Topics Covered

- Function declaration and function calls
- Functions with parameters
- Functions with return values
- Multiple return values
- Error handling using `error` and `fmt.Errorf()`
- Named return values
- Exported and unexported functions (uppercase and lowercase naming)

## Practical Project: Go ATM

Built a command-line ATM application using reusable functions.

### Features

- Display ATM menu
- Check account balance
- Deposit money
- Withdraw money
- Validate deposit and withdrawal amounts
- Handle insufficient balance
- Display welcome and exit messages

### Functions Implemented

| Function | Purpose |
|---|---|
| `display_menu()` | Displays ATM menu |
| `get_balance()` | Displays current balance |
| `deposit_amount()` | Deposits money and returns updated balance |
| `withdraw_amount()` | Withdraws money and returns updated balance |
| `start_greet()` | Displays welcome message |
| `end_greet()` | Displays exit message |

## Key Learnings

- Functions help organize code into reusable blocks.
- Parameters allow functions to receive data.
- Return values allow functions to send results back.
- Go supports multiple return values.
- Errors are commonly returned as values and checked using `if err != nil`.
- Lowercase function names are unexported; uppercase names are exported.