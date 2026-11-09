# Day 6 – Structs and Methods

## Overview

Learned how to use structs to group related data and methods to define behavior associated with custom types in Go.

## Topics Covered

- Struct declaration and initialization
- Accessing and updating struct fields
- Struct zero values
- Creating methods with receivers
- Value receivers vs pointer receivers
- Modifying original struct values using pointer receivers

## Practical Examples

### 1. Student

- Created a `Student` struct with ID, Name, Email, and Department.
- Implemented a `Display()` method.
- Implemented an `UpdateName()` method using a pointer receiver.

### 2. Bank Account

Built an interactive Bank Account application.

Features:
- Display account information
- Deposit money
- Withdraw money
- Validate deposit and withdrawal amounts
- Prevent withdrawals exceeding the available balance
- Interactive menu using loops and switch statements

## Key Learnings

- Structs group related fields into a custom data type.
- Methods associate behavior with a type.
- Value receivers work with copies of structs.
- Pointer receivers can modify the original struct.
- Go automatically handles pointer method calls on addressable variables.
- Uninitialized struct fields receive their respective zero values.