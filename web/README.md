# Rent vs Buy Calculator - Web Version

A comprehensive financial calculator to compare buying vs renting, or selling vs keeping property. Built with Svelte, TypeScript, and Tailwind CSS.

## Features

- **Two Scenarios**:
  - BUY vs RENT: Compare purchasing vs renting
  - SELL vs KEEP: Analyze selling vs keeping existing property

- **Comprehensive Analysis**:
  - Loan amortization details
  - Expenditure comparisons
  - Sale proceeds analysis with capital gains tax
  - Net worth projections with investment returns
  - Inflation-adjusted costs

- **Local Storage**: Automatically saves your inputs for next visit

## Getting Started

### Prerequisites

- Node.js (v18 or higher recommended)
- npm or pnpm

### Installation

```bash
# Install dependencies
npm install
```

### Development

```bash
# Start development server
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) in your browser.

### Build for Production

```bash
# Build for production
npm run build
```

The built files will be in the `dist/` directory, ready to be deployed to any static hosting service.

## Deployment to blot.im

After building:

1. Run `npm run build`
2. Copy all files from the `dist/` directory to your blot.im folder
3. blot.im will automatically serve the files

The application is a single-page application (SPA) that runs entirely in the browser with no backend required.

## Input Format

- **Currency**: Enter as numbers or use suffixes (5K = 5,000, 2.5M = 2,500,000)
- **Percentages**: Enter as numbers (e.g., 6.5 for 6.5%)
- **Duration**: Use format like "30y" for 30 years, "6m" for 6 months
- **Rates Arrays**: Comma-separated for different years (e.g., "5,3,2" = 5% year 1, 3% year 2, 2% year 3+)

## Technology Stack

- **Svelte** - Reactive UI framework
- **TypeScript** - Type-safe code
- **Tailwind CSS** - Utility-first styling
- **Vite** - Fast build tool

## License

MIT
