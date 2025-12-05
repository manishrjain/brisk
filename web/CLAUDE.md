# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
npm run dev      # Start development server (http://localhost:3000)
npm run build    # Build for production (output: dist/)
npm run preview  # Preview production build
npm run check    # Run svelte-check for TypeScript validation
```

## Architecture

This is a Svelte/TypeScript SPA for comparing buy vs rent or sell vs keep decisions for any asset (real estate, cars, trucks, etc.). It is a port of a Go CLI calculator.

### Core Flow

1. **App.svelte** - Root component managing view state (form vs results), keyboard shortcuts (Ctrl+S save, Ctrl+O load, Ctrl+Shift+S share, Escape back), localStorage persistence, and URL sharing
2. **TerminalForm.svelte** - Input form with terminal aesthetic
3. **ResultsDisplay.svelte** - Renders calculation result tables

### Calculation Engine

**src/lib/calculator.ts** contains all financial calculations:
- `calculate()` - Main entry point, returns `CalculationResults`
- `calculateMonthlyPayment()` - Loan amortization formula
- `populateMonthlyCosts()` - Generates 360-month arrays for buying/renting costs, loan balances, principal/interest tracking
- `calculateAssetValue()` - Computes appreciated value using year-by-year rates array
- `calculateSaleProceeds()` - Sale price, costs, capital gains tax, net proceeds
- `calculateRentingNetWorth()` - Investment growth from renting scenario savings

### Two Scenarios

- **buy_vs_rent**: Compares purchasing vs renting. Outputs amortization, expenditure, sale proceeds, and net worth comparison tables.
- **sell_vs_keep**: Analyzes selling now vs keeping property. Outputs keep expenses breakdown, sale proceeds, and sell vs keep comparison tables.

### Types (src/types.ts)

- `CalculatorInputs` - All form inputs including scenario-specific fields
- `CalculationResults` - Monthly cost arrays plus table data for display
- Row types: `AmortizationRow`, `ExpenditureRow`, `SaleProceedsRow`, `ComparisonRow`, `SellVsKeepRow`, `KeepExpensesRow`

### Supporting Modules

- **src/lib/formatter.ts** - Parse currency (500K, 2.5M), duration (30y, 6m), percentages, rate arrays
- **src/lib/storage.ts** - Named config save/load to localStorage
- **src/lib/share.ts** - URL encoding/decoding for sharing calculator state
- **src/lib/theme.ts** - Dark/light mode with Monokai color scheme

### Input Formats

- Currency: `500K`, `2.5M`, or plain numbers
- Duration: `30y` (years), `6m` (months), `15y6m` (combined)
- Rate arrays: `5,3,2` means 5% year 1, 3% year 2, 2% year 3+
