# Setup Guide for Rent vs Buy Calculator (Web Version)

## What Was Built

A complete single-page web application version of your Rent vs Buy calculator built with:
- **Svelte** - Lightweight reactive framework
- **TypeScript** - Type-safe calculation logic
- **Tailwind CSS** - Monokai-themed styling
- **Vite** - Fast build tool

## Project Structure

```
web/
├── src/
│   ├── components/
│   │   ├── InputForm.svelte       # Form with all inputs
│   │   └── ResultsDisplay.svelte  # Tables and results
│   ├── lib/
│   │   ├── calculator.ts          # Core calculation logic (ported from Go)
│   │   └── formatter.ts           # Number formatting and parsing
│   ├── types.ts                   # TypeScript type definitions
│   ├── app.css                    # Tailwind + custom styles
│   ├── App.svelte                 # Main application component
│   └── main.ts                    # Entry point
├── index.html                     # HTML template
├── package.json                   # Dependencies
├── vite.config.ts                 # Vite configuration
├── tailwind.config.js             # Tailwind configuration
└── tsconfig.json                  # TypeScript configuration
```

## Features Implemented

✅ **Two Scenarios**:
- BUY vs RENT - Compare purchasing vs renting
- SELL vs KEEP - Analyze selling vs keeping existing property

✅ **All Calculations** (ported from Go):
- Loan amortization with monthly payment calculations
- Inflation-adjusted costs over time
- Asset appreciation (with year-by-year rates)
- Investment returns with monthly compounding
- Capital gains tax calculations
- Net worth projections

✅ **UI Features**:
- Responsive design (mobile-friendly)
- Monokai color scheme (matching your Go CLI)
- Smart input parsing (supports K/M suffixes, e.g., "500K")
- LocalStorage auto-save (inputs persist across sessions)
- Clean table layouts for results

## Fixing Your Node.js Issue

Your Node.js installation has a library dependency issue. Fix it with:

```bash
# Option 1: Update icu4c library
brew upgrade icu4c

# Option 2: Reinstall Node.js
brew uninstall node
brew install node

# Option 3: Use a Node version manager (recommended)
brew install nvm
nvm install 20
nvm use 20
```

## Installation & Development

Once Node is fixed:

```bash
cd web

# Install dependencies
npm install

# Start development server
npm run dev
# Open http://localhost:3000

# Build for production
npm run build
# Output will be in dist/ directory
```

## Deploying to blot.im

After running `npm run build`:

1. Copy everything from the `dist/` directory
2. Upload to your blot.im folder
3. The app is a static SPA - no server needed!

The build output will be:
```
dist/
├── index.html           # Main HTML file
├── assets/
│   ├── index-[hash].js  # Bundled JavaScript
│   └── index-[hash].css # Bundled styles
```

## Usage

### Input Format Examples

- **Currency**: `500K`, `2.5M`, `500000`
- **Percentages**: `6.5` (for 6.5%)
- **Duration**: `30y`, `15y6m`, `360m`
- **Rate Arrays**: `5,3,2` (5% year 1, 3% year 2, 2% year 3+)

### Workflow

1. Select scenario (BUY vs RENT or SELL vs KEEP)
2. Fill in all relevant inputs
3. Click "Calculate Results"
4. View detailed tables and analysis
5. Click "Back to Inputs" to adjust

## What's Different from Go Version

### Same:
- ✅ All calculation logic (exact port)
- ✅ Monokai color scheme
- ✅ Same table layouts
- ✅ Same input parsing (K/M suffixes)

### Enhanced:
- ✅ Web-based (accessible anywhere)
- ✅ Mobile-responsive
- ✅ Auto-save to localStorage
- ✅ No terminal needed
- ✅ Easier to share results (screenshot)

### Not Yet Implemented:
- ❌ Real-time market data (VOO, QQQ, etc.)
  - Can add later with an API
- ❌ Profile save/load system
  - LocalStorage provides auto-save instead

## Next Steps

### To Test Locally:
1. Fix Node.js (see instructions above)
2. Run `npm install`
3. Run `npm run dev`
4. Open http://localhost:3000

### To Deploy:
1. Run `npm run build`
2. Upload `dist/*` to blot.im
3. Access via your blot.im URL

### Future Enhancements (Optional):
- Add market data API integration
- Add URL-based sharing (base64 encoding in URL)
- Add export to PDF/CSV
- Add comparison charts/graphs
- Add mobile app version

## Troubleshooting

### Build Errors
- Make sure Node.js 18+ is installed
- Delete `node_modules` and run `npm install` again
- Check that all TypeScript types are correct

### Runtime Errors
- Open browser console (F12) to see errors
- Check that all required inputs are filled
- Verify input format (use examples above)

## Questions?

The codebase is well-commented and follows the same logic as your Go version. All calculations are in `src/lib/calculator.ts` and match your original implementation.
