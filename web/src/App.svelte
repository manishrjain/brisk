<script lang="ts">
  import { onMount } from 'svelte';
  import type { CalculatorInputs, CalculationResults } from './types';
  import { calculate } from './lib/calculator';
  import InputForm from './components/InputForm.svelte';
  import ResultsDisplay from './components/ResultsDisplay.svelte';

  // String versions for form binding
  let formInputs = {
    scenario: 'buy_vs_rent',
    inflationRate: '3',
    investmentReturnRate: '10',
    include30Year: false,
    purchasePrice: '500000',
    currentMarketValue: '',
    loanAmount: '400000',
    loanRate: '6.5',
    loanTerm: '30y',
    remainingLoanTerm: '',
    annualInsurance: '3000',
    annualTaxes: '5000',
    monthlyExpenses: '500',
    appreciationRate: '3',
    rentDeposit: '5000',
    monthlyRent: '3000',
    annualRentCosts: '1000',
    otherAnnualCosts: '500',
    includeSelling: false,
    includeRentingSell: false,
    agentCommission: '6',
    stagingCosts: '10000',
    taxFreeLimits: '250000',
    capitalGainsTax: '20',
  };

  let results: CalculationResults | null = null;
  let calculatedInputs: CalculatorInputs | null = null;
  let showResults = false;

  onMount(() => {
    // Load saved inputs from localStorage
    const saved = localStorage.getItem('rentobuy_inputs');
    if (saved) {
      try {
        formInputs = { ...formInputs, ...JSON.parse(saved) };
      } catch (e) {
        console.error('Failed to load saved inputs:', e);
      }
    }
  });

  function handleCalculate(event: CustomEvent) {
    try {
      const inputs: CalculatorInputs = event.detail;
      calculatedInputs = inputs;
      results = calculate(inputs);
      showResults = true;
      // Save form inputs to localStorage
      localStorage.setItem('rentobuy_inputs', JSON.stringify(formInputs));
    } catch (error) {
      console.error('Calculation error:', error);
      alert('Error calculating results. Please check your inputs.');
    }
  }

  function handleReset() {
    showResults = false;
    results = null;
  }
</script>

<main class="min-h-screen bg-monokai-bg text-monokai-text p-4 md:p-8">
  <div class="max-w-7xl mx-auto">
    <header class="mb-8">
      <h1 class="text-4xl font-bold text-monokai-pink mb-2">
        Rent vs Buy Calculator
      </h1>
      <p class="text-monokai-text-muted">
        Compare buying vs renting, or selling vs keeping your property
      </p>
    </header>

    {#if !showResults}
      <InputForm bind:formInputs on:calculate={handleCalculate} />
    {:else}
      <div class="mb-6">
        <button class="btn-secondary" on:click={handleReset}>
          ← Back to Inputs
        </button>
      </div>
      {#if results && calculatedInputs}
        <ResultsDisplay inputs={calculatedInputs} {results} />
      {/if}
    {/if}
  </div>
</main>
