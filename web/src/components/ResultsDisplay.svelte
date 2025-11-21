<script lang="ts">
  import type { CalculatorInputs, CalculationResults } from '../types';
  import { formatCurrency, formatPercent } from '../lib/formatter';

  export let inputs: CalculatorInputs;
  export let results: CalculationResults;

  function formatDuration(months: number): string {
    if (months % 12 === 0) {
      return `${months / 12}y`;
    }
    const years = Math.floor(months / 12);
    const remainingMonths = months % 12;
    if (years === 0) {
      return `${remainingMonths}m`;
    }
    return `${years}y ${remainingMonths}m`;
  }
</script>

<div class="space-y-8">
  <!-- Input Parameters Summary -->
  <section class="bg-monokai-bg-light p-6 rounded-lg">
    <h2 class="section-title">Input Parameters</h2>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 text-sm">
      <div>
        <span class="text-monokai-cyan">Scenario:</span>
        <span class="ml-2">{inputs.scenario === 'buy_vs_rent' ? 'BUY vs RENT' : 'SELL vs KEEP'}</span>
      </div>
      <div>
        <span class="text-monokai-cyan">Inflation Rate:</span>
        <span class="ml-2">{formatPercent(inputs.inflationRate)}</span>
      </div>
      <div>
        <span class="text-monokai-cyan">Investment Return:</span>
        <span class="ml-2">{formatPercent(inputs.investmentReturnRate)}</span>
      </div>
      <div>
        <span class="text-monokai-cyan">Purchase Price:</span>
        <span class="ml-2">{formatCurrency(inputs.purchasePrice, true)}</span>
      </div>
      {#if inputs.currentMarketValue}
        <div>
          <span class="text-monokai-cyan">Current Market Value:</span>
          <span class="ml-2">{formatCurrency(inputs.currentMarketValue, true)}</span>
        </div>
      {/if}
      <div>
        <span class="text-monokai-cyan">Loan Amount:</span>
        <span class="ml-2">{formatCurrency(inputs.loanAmount, true)}</span>
      </div>
      <div>
        <span class="text-monokai-cyan">Loan Rate:</span>
        <span class="ml-2">{formatPercent(inputs.loanRate)}</span>
      </div>
      <div>
        <span class="text-monokai-cyan">Loan Term:</span>
        <span class="ml-2">{formatDuration(inputs.loanTerm)}</span>
      </div>
    </div>
  </section>

  <!-- Amortization Table (BUY vs RENT only) -->
  {#if results.amortizationTable && inputs.loanAmount > 0}
    <section class="bg-monokai-bg-light p-6 rounded-lg">
      <h2 class="section-title">Loan Amortization Details</h2>
      <div class="table-container">
        <table class="data-table">
          <thead>
            <tr>
              <th>Period</th>
              <th class="text-right">Principal Paid</th>
              <th class="text-right">Interest Paid</th>
              <th class="text-right">Loan Balance</th>
            </tr>
          </thead>
          <tbody>
            {#each results.amortizationTable as row}
              <tr>
                <td class="font-mono">{row.period}</td>
                <td class="text-right font-mono">{formatCurrency(row.principalPaid)}</td>
                <td class="text-right font-mono">{formatCurrency(row.interestPaid)}</td>
                <td class="text-right font-mono">{formatCurrency(row.loanBalance)}</td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
      <p class="help-text mt-2">
        Note: Monthly payment is fixed. Each payment covers interest on remaining balance, with the rest going to principal.
      </p>
    </section>
  {/if}

  <!-- Expenditure Table (BUY vs RENT only) -->
  {#if results.expenditureTable}
    <section class="bg-monokai-bg-light p-6 rounded-lg">
      <h2 class="section-title">Total Expenditure Comparison</h2>
      <div class="table-container">
        <table class="data-table">
          <thead>
            <tr>
              <th>Period</th>
              <th class="text-right">Buying Expend.</th>
              <th class="text-right">Renting Expend.</th>
              <th class="text-right">Difference</th>
            </tr>
          </thead>
          <tbody>
            {#each results.expenditureTable as row}
              <tr>
                <td class="font-mono">{row.period}</td>
                <td class="text-right font-mono">{formatCurrency(row.buyingExpenditure)}</td>
                <td class="text-right font-mono">{formatCurrency(row.rentingExpenditure)}</td>
                <td class="text-right font-mono">{formatCurrency(row.difference)}</td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
      <p class="help-text mt-2">
        Note: All recurring costs (insurance, taxes, rent, HOA, etc.) are inflated annually at {formatPercent(inputs.inflationRate)} rate.
      </p>
    </section>
  {/if}

  <!-- Sale Proceeds -->
  <section class="bg-monokai-bg-light p-6 rounded-lg">
    <h2 class="section-title">Sale Proceeds Analysis</h2>
    <div class="table-container">
      <table class="data-table">
        <thead>
          <tr>
            <th>Period</th>
            <th class="text-right">Sale Price</th>
            <th class="text-right">Selling Cost</th>
            <th class="text-right">Loan Payoff</th>
            <th class="text-right">Cap Gains</th>
            <th class="text-right">Tax</th>
            <th class="text-right">Net Proceeds</th>
          </tr>
        </thead>
        <tbody>
          {#each results.saleProceedsTable as row}
            <tr>
              <td class="font-mono">{row.period}</td>
              <td class="text-right font-mono">{formatCurrency(row.salePrice)}</td>
              <td class="text-right font-mono">{formatCurrency(row.totalSellingCosts)}</td>
              <td class="text-right font-mono">{formatCurrency(row.loanPayoff)}</td>
              <td class="text-right font-mono">{formatCurrency(row.capitalGains)}</td>
              <td class="text-right font-mono">{formatCurrency(row.taxOnGains)}</td>
              <td class="text-right font-mono">{formatCurrency(row.netProceeds)}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
    <p class="help-text mt-2">
      Note: Appreciation rates are applied year-by-year (compounded). Sale price = compounded property value.
    </p>
  </section>

  <!-- Comparison Table (BUY vs RENT) -->
  {#if results.comparisonTable}
    <section class="bg-monokai-bg-light p-6 rounded-lg">
      <h2 class="section-title">Net Worth Projections: BUY vs RENT</h2>
      <div class="table-container">
        <table class="data-table">
          <thead>
            <tr>
              <th>Period</th>
              <th class="text-right">Asset Value</th>
              <th class="text-right">Buying NW</th>
              <th class="text-right">Cum Savings</th>
              <th class="text-right">Market Return</th>
              <th class="text-right">Renting NW</th>
              <th class="text-right">RENT - BUY</th>
            </tr>
          </thead>
          <tbody>
            {#each results.comparisonTable as row}
              <tr>
                <td class="font-mono">{row.period}</td>
                <td class="text-right font-mono">{formatCurrency(row.assetValue)}</td>
                <td class="text-right font-mono">{formatCurrency(row.buyingNetWorth)}</td>
                <td class="text-right font-mono">{formatCurrency(row.cumulativeSavings)}</td>
                <td class="text-right font-mono">{formatCurrency(row.marketReturn)}</td>
                <td class="text-right font-mono">{formatCurrency(row.rentingNetWorth)}</td>
                <td class="text-right font-mono" class:text-monokai-green={row.difference < 0} class:text-monokai-pink={row.difference > 0}>
                  {formatCurrency(row.difference)}
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
      <p class="help-text mt-2">
        Note: 'Cum Savings' = raw cost difference without investment growth.
        'Market Return' = investment growth at {formatPercent(inputs.investmentReturnRate)} annual rate.
        Positive RENT - BUY means renting wins, negative means buying wins.
      </p>
    </section>
  {/if}

  <!-- Sell vs Keep Comparison -->
  {#if results.sellVsKeepTable}
    <section class="bg-monokai-bg-light p-6 rounded-lg">
      <h2 class="section-title">Net Worth Projections: SELL vs KEEP</h2>
      <div class="table-container">
        <table class="data-table">
          <thead>
            <tr>
              <th>Period</th>
              {#if inputs.includeRentingSell}
                <th class="text-right">SELL Cum. Exp</th>
              {/if}
              <th class="text-right">SELL Net Worth</th>
              <th class="text-right">KEEP Net Position</th>
              <th class="text-right">KEEP Net Proceeds</th>
              <th class="text-right">KEEP - SELL</th>
            </tr>
          </thead>
          <tbody>
            {#each results.sellVsKeepTable as row}
              <tr>
                <td class="font-mono">{row.period}</td>
                {#if row.sellCumulativeExpenses !== undefined}
                  <td class="text-right font-mono">{formatCurrency(row.sellCumulativeExpenses)}</td>
                {/if}
                <td class="text-right font-mono">{formatCurrency(row.sellNetWorth)}</td>
                <td class="text-right font-mono">{formatCurrency(row.keepNetPosition)}</td>
                <td class="text-right font-mono">{formatCurrency(row.keepNetProceeds)}</td>
                <td class="text-right font-mono" class:text-monokai-green={row.difference > 0} class:text-monokai-pink={row.difference < 0}>
                  {formatCurrency(row.difference)}
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
      <p class="help-text mt-2">
        Note: 'SELL Net Worth' = Net proceeds from selling now invested at {formatPercent(inputs.investmentReturnRate)} return.
        'KEEP Net Position' = Investment value from income minus real costs.
        Positive KEEP - SELL means keeping wins, negative means selling wins.
      </p>
    </section>
  {/if}
</div>
