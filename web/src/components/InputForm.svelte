<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import type { CalculatorInputs, ScenarioType } from '../types';
  import { parseAmount, parseDuration, parseAppreciationRates } from '../lib/formatter';

  export let formInputs: any;

  const dispatch = createEventDispatcher();

  function convertFormInputsToCalculatorInputs(): CalculatorInputs {
    return {
      scenario: formInputs.scenario as ScenarioType,
      inflationRate: parseFloat(formInputs.inflationRate) || 0,
      investmentReturnRate: parseFloat(formInputs.investmentReturnRate) || 0,
      include30Year: formInputs.include30Year,
      purchasePrice: parseAmount(formInputs.purchasePrice.toString()),
      currentMarketValue: formInputs.currentMarketValue ? parseAmount(formInputs.currentMarketValue.toString()) : undefined,
      loanAmount: parseAmount(formInputs.loanAmount.toString()),
      loanRate: parseFloat(formInputs.loanRate) || 0,
      loanTerm: parseDuration(formInputs.loanTerm),
      remainingLoanTerm: formInputs.remainingLoanTerm ? parseDuration(formInputs.remainingLoanTerm) : undefined,
      annualInsurance: parseAmount(formInputs.annualInsurance.toString()),
      annualTaxes: parseAmount(formInputs.annualTaxes.toString()),
      monthlyExpenses: parseAmount(formInputs.monthlyExpenses.toString()),
      appreciationRate: parseAppreciationRates(formInputs.appreciationRate),
      rentDeposit: parseAmount(formInputs.rentDeposit.toString()),
      monthlyRent: parseAmount(formInputs.monthlyRent.toString()),
      annualRentCosts: parseAmount(formInputs.annualRentCosts.toString()),
      otherAnnualCosts: parseAmount(formInputs.otherAnnualCosts.toString()),
      includeSelling: formInputs.includeSelling,
      includeRentingSell: formInputs.includeRentingSell,
      agentCommission: parseFloat(formInputs.agentCommission) || 0,
      stagingCosts: parseAmount(formInputs.stagingCosts.toString()),
      taxFreeLimits: parseAppreciationRates(formInputs.taxFreeLimits),
      capitalGainsTax: parseFloat(formInputs.capitalGainsTax) || 0,
    };
  }

  function handleSubmit() {
    try {
      const inputs = convertFormInputsToCalculatorInputs();
      dispatch('calculate', inputs);
    } catch (error) {
      alert(`Invalid input: ${error instanceof Error ? error.message : 'Unknown error'}`);
    }
  }
</script>

<form on:submit|preventDefault={handleSubmit} class="space-y-8">
  <!-- Scenario Selection -->
  <section class="bg-monokai-bg-light p-6 rounded-lg">
    <h2 class="group-title">Scenario Selection</h2>
    <div class="flex gap-4">
      <label class="flex items-center gap-2 cursor-pointer">
        <input
          type="radio"
          name="scenario"
          value="buy_vs_rent"
          bind:group={formInputs.scenario}
          class="w-4 h-4 text-monokai-pink"
        />
        <span>BUY vs RENT</span>
      </label>
      <label class="flex items-center gap-2 cursor-pointer">
        <input
          type="radio"
          name="scenario"
          value="sell_vs_keep"
          bind:group={formInputs.scenario}
          class="w-4 h-4 text-monokai-pink"
        />
        <span>SELL vs KEEP</span>
      </label>
    </div>
  </section>

  <!-- Economic Assumptions -->
  <section class="bg-monokai-bg-light p-6 rounded-lg">
    <h2 class="group-title">Economic Assumptions</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <div>
        <label class="label">
          Inflation Rate (%)
          <input
            type="text"
            bind:value={formInputs.inflationRate}
            class="input-field w-full mt-1"
            placeholder="3"
          />
        </label>
        <p class="help-text">Annual inflation for all recurring costs</p>
      </div>
      <div>
        <label class="label">
          Investment Return Rate (%)
          <input
            type="text"
            bind:value={formInputs.investmentReturnRate}
            class="input-field w-full mt-1"
            placeholder="10"
          />
        </label>
        <p class="help-text">Expected return on investments</p>
      </div>
      <div class="md:col-span-2">
        <label class="flex items-center gap-2 cursor-pointer">
          <input
            type="checkbox"
            bind:checked={formInputs.include30Year}
            class="w-4 h-4 text-monokai-pink"
          />
          <span>Include 30-Year Projections</span>
        </label>
        <p class="help-text">Show 15y, 20y, 30y periods (default: 10y max)</p>
      </div>
    </div>
  </section>

  <!-- Buying/Asset Section -->
  <section class="bg-monokai-bg-light p-6 rounded-lg">
    <h2 class="group-title">
      {formInputs.scenario === 'sell_vs_keep' ? 'Asset' : 'Buying'}
    </h2>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <div>
        <label class="label">
          {formInputs.scenario === 'sell_vs_keep' ? 'Original Purchase Price' : 'Asset Purchase Price'}
          <input
            type="text"
            bind:value={formInputs.purchasePrice}
            class="input-field w-full mt-1"
            placeholder="500K or 500000"
          />
        </label>
        <p class="help-text">
          {formInputs.scenario === 'sell_vs_keep' ? 'What you originally paid (for capital gains)' : 'Initial purchase price. Use K/M suffix (e.g., 500K)'}
        </p>
      </div>

      {#if formInputs.scenario === 'sell_vs_keep'}
        <div>
          <label class="label">
            Current Market Value
            <input
              type="text"
              bind:value={formInputs.currentMarketValue}
              class="input-field w-full mt-1"
              placeholder="2.2M"
            />
          </label>
          <p class="help-text">What the asset is worth today</p>
        </div>
      {/if}

      <div>
        <label class="label">
          {formInputs.scenario === 'sell_vs_keep' ? 'Original Loan Amount' : 'Loan Amount'}
          <input
            type="text"
            bind:value={formInputs.loanAmount}
            class="input-field w-full mt-1"
            placeholder="400K"
          />
        </label>
        <p class="help-text">Total mortgage/loan amount</p>
      </div>

      <div>
        <label class="label">
          Loan Rate (%)
          <input
            type="text"
            bind:value={formInputs.loanRate}
            class="input-field w-full mt-1"
            placeholder="6.5"
          />
        </label>
        <p class="help-text">Annual interest rate</p>
      </div>

      <div>
        <label class="label">
          Loan Term
          <input
            type="text"
            bind:value={formInputs.loanTerm}
            placeholder="30y or 360m"
            class="input-field w-full mt-1"
          />
        </label>
        <p class="help-text">Duration (e.g., 30y, 15y)</p>
      </div>

      {#if formInputs.scenario === 'sell_vs_keep'}
        <div>
          <label class="label">
            Remaining Loan Term
            <input
              type="text"
              bind:value={formInputs.remainingLoanTerm}
              placeholder="25y"
              class="input-field w-full mt-1"
            />
          </label>
          <p class="help-text">Time left on loan</p>
        </div>
      {/if}

      <div>
        <label class="label">
          Annual Tax & Insurance
          <input
            type="text"
            bind:value={formInputs.annualInsurance}
            class="input-field w-full mt-1"
            placeholder="3K"
          />
        </label>
        <p class="help-text">Yearly insurance cost</p>
      </div>

      <div>
        <label class="label">
          Other Annual Costs
          <input
            type="text"
            bind:value={formInputs.annualTaxes}
            class="input-field w-full mt-1"
            placeholder="5K"
          />
        </label>
        <p class="help-text">Taxes, HOA, maintenance, etc.</p>
      </div>

      <div>
        <label class="label">
          Monthly Expenses
          <input
            type="text"
            bind:value={formInputs.monthlyExpenses}
            class="input-field w-full mt-1"
            placeholder="500 or -4K"
          />
        </label>
        <p class="help-text">Utilities, etc. (can be negative if earning income)</p>
      </div>

      <div class="md:col-span-2">
        <label class="label">
          Appreciation Rate (%)
          <input
            type="text"
            bind:value={formInputs.appreciationRate}
            placeholder="3 or 5,3,2"
            class="input-field w-full mt-1"
          />
        </label>
        <p class="help-text">
          Annual rate (comma-separated for different years, e.g., '10,5,3' = 10% yr1, 5% yr2, 3% yr3+)
        </p>
      </div>
    </div>
  </section>

  <!-- Renting Section -->
  <section class="bg-monokai-bg-light p-6 rounded-lg">
    <h2 class="group-title">
      {formInputs.scenario === 'sell_vs_keep' ? 'Investing (if selling)' : 'Renting'}
    </h2>

    {#if formInputs.scenario === 'sell_vs_keep'}
      <div class="mb-4">
        <label class="flex items-center gap-2 cursor-pointer">
          <input
            type="checkbox"
            bind:checked={formInputs.includeRentingSell}
            class="w-4 h-4 text-monokai-pink"
          />
          <span>Include Renting Analysis</span>
        </label>
        <p class="help-text">Toggle if selling means you'll need to rent</p>
      </div>
    {/if}

    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <div>
        <label class="label">
          Rental Deposit
          <input
            type="text"
            bind:value={formInputs.rentDeposit}
            class="input-field w-full mt-1"
            placeholder="5K"
          />
        </label>
        <p class="help-text">Initial rental deposit</p>
      </div>

      <div>
        <label class="label">
          Monthly Rent
          <input
            type="text"
            bind:value={formInputs.monthlyRent}
            class="input-field w-full mt-1"
            placeholder="3K"
          />
        </label>
        <p class="help-text">Base monthly rent amount</p>
      </div>

      <div>
        <label class="label">
          Annual Rent Costs
          <input
            type="text"
            bind:value={formInputs.annualRentCosts}
            class="input-field w-full mt-1"
            placeholder="1K"
          />
        </label>
        <p class="help-text">Yearly rental-related costs</p>
      </div>

      {#if formInputs.scenario === 'buy_vs_rent'}
        <div>
          <label class="label">
            Other Annual Costs
            <input
              type="text"
              bind:value={formInputs.otherAnnualCosts}
              class="input-field w-full mt-1"
              placeholder="500"
            />
          </label>
          <p class="help-text">Additional yearly costs for renting</p>
        </div>
      {/if}
    </div>
  </section>

  <!-- Selling Section -->
  <section class="bg-monokai-bg-light p-6 rounded-lg">
    <h2 class="group-title">Selling</h2>

    {#if formInputs.scenario === 'buy_vs_rent'}
      <div class="mb-4">
        <label class="flex items-center gap-2 cursor-pointer">
          <input
            type="checkbox"
            bind:checked={formInputs.includeSelling}
            class="w-4 h-4 text-monokai-pink"
          />
          <span>Include Selling Analysis</span>
        </label>
        <p class="help-text">Toggle to enable/disable selling analysis</p>
      </div>
    {/if}

    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <div>
        <label class="label">
          Agent Commission (%)
          <input
            type="text"
            bind:value={formInputs.agentCommission}
            class="input-field w-full mt-1"
            placeholder="6"
          />
        </label>
        <p class="help-text">Percentage of sale price</p>
      </div>

      <div>
        <label class="label">
          Staging/Selling Costs
          <input
            type="text"
            bind:value={formInputs.stagingCosts}
            class="input-field w-full mt-1"
            placeholder="10K"
          />
        </label>
        <p class="help-text">Fixed costs to prepare and sell</p>
      </div>

      <div>
        <label class="label">
          Tax-Free Gains Limit
          <input
            type="text"
            bind:value={formInputs.taxFreeLimits}
            placeholder="250K or 500K,0"
            class="input-field w-full mt-1"
          />
        </label>
        <p class="help-text">Capital gains exempt from tax (comma-separated for different years)</p>
      </div>

      <div>
        <label class="label">
          Capital Gains Tax Rate (%)
          <input
            type="text"
            bind:value={formInputs.capitalGainsTax}
            class="input-field w-full mt-1"
            placeholder="20"
          />
        </label>
        <p class="help-text">Long-term capital gains tax rate</p>
      </div>
    </div>
  </section>

  <!-- Submit Button -->
  <div class="flex justify-center">
    <button type="submit" class="btn-primary text-lg px-8 py-3">
      Calculate Results
    </button>
  </div>
</form>
