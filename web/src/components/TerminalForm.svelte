<script lang="ts">
  import { createEventDispatcher, onMount } from 'svelte';
  import type { CalculatorInputs, ScenarioType } from '../types';
  import { parseAmount, parseDuration, parseAppreciationRates } from '../lib/formatter';

  export let formInputs: any;

  const dispatch = createEventDispatcher();

  interface FormField {
    key: string;
    label: string;
    help: string;
    placeholder: string;
    visible: () => boolean;
    isHeader?: boolean;
    headerText?: string;
    toggleValues?: string[];  // For fields that can toggle between two values
  }

  let fields: FormField[] = [];
  let currentFieldIndex = 0;
  let inputRefs: HTMLInputElement[] = [];

  $: {
    // Rebuild fields list when scenario changes
    fields = [
      { key: 'header_scenario', label: '', help: '', placeholder: '', visible: () => true, isHeader: true, headerText: 'SCENARIO SELECTION' },
      { key: 'scenario', label: 'Scenario', help: 'buy_vs_rent or sell_vs_keep', placeholder: 'buy_vs_rent', visible: () => true, toggleValues: ['buy_vs_rent', 'sell_vs_keep'] },

      { key: 'header_economic', label: '', help: '', placeholder: '', visible: () => true, isHeader: true, headerText: 'ECONOMIC ASSUMPTIONS' },
      { key: 'inflationRate', label: 'Inflation Rate (%)', help: 'Annual inflation for all recurring costs', placeholder: '3', visible: () => true },
      { key: 'investmentReturnRate', label: 'Investment Return (%)', help: 'Expected return on investments', placeholder: '10', visible: () => true },
      { key: 'include30Year', label: '30-Year Projections', help: 'yes/no - Show extended periods', placeholder: 'no', visible: () => true, toggleValues: ['yes', 'no'] },

      { key: 'header_asset', label: '', help: '', placeholder: '', visible: () => true, isHeader: true, headerText: formInputs.scenario === 'sell_vs_keep' ? 'ASSET DETAILS' : 'BUYING' },
      { key: 'purchasePrice', label: formInputs.scenario === 'sell_vs_keep' ? 'Original Purchase Price' : 'Purchase Price', help: 'Initial/original purchase price (use K/M)', placeholder: '500K', visible: () => true },
      { key: 'currentMarketValue', label: 'Current Market Value', help: 'What the asset is worth today (use K/M)', placeholder: '2.2M', visible: () => formInputs.scenario === 'sell_vs_keep' },
      { key: 'loanAmount', label: formInputs.scenario === 'sell_vs_keep' ? 'Original Loan Amount' : 'Loan Amount', help: 'Total mortgage/loan amount (use K/M)', placeholder: '400K', visible: () => true },
      { key: 'loanRate', label: 'Loan Rate (%)', help: 'Annual interest rate', placeholder: '6.5', visible: () => true },
      { key: 'loanTerm', label: 'Loan Term', help: 'Duration (e.g., 30y, 15y)', placeholder: '30y', visible: () => true },
      { key: 'remainingLoanTerm', label: 'Remaining Loan Term', help: 'Time left on loan (e.g., 25y)', placeholder: '25y', visible: () => formInputs.scenario === 'sell_vs_keep' },
      { key: 'annualInsurance', label: 'Annual Tax & Insurance', help: 'Yearly insurance + tax (use K/M)', placeholder: '3K', visible: () => true },
      { key: 'annualTaxes', label: 'Other Annual Costs', help: 'HOA, maintenance, etc (use K/M)', placeholder: '5K', visible: () => true },
      { key: 'monthlyExpenses', label: 'Monthly Expenses', help: 'Utilities, etc (can be negative, use K)', placeholder: '500', visible: () => true },
      { key: 'appreciationRate', label: 'Appreciation Rate (%)', help: 'Annual rate, comma-separated (e.g., 10,5,3)', placeholder: '3', visible: () => true },

      { key: 'header_renting', label: '', help: '', placeholder: '', visible: () => true, isHeader: true, headerText: formInputs.scenario === 'sell_vs_keep' ? 'INVESTING (IF SELLING)' : 'RENTING' },
      { key: 'includeRentingSell', label: 'Include Renting (Sell)', help: 'yes/no - If selling means renting', placeholder: 'no', visible: () => formInputs.scenario === 'sell_vs_keep', toggleValues: ['yes', 'no'] },
      { key: 'rentDeposit', label: 'Rental Deposit', help: 'Initial rental deposit (use K/M)', placeholder: '5K', visible: () => true },
      { key: 'monthlyRent', label: 'Monthly Rent', help: 'Base monthly rent (use K/M)', placeholder: '3K', visible: () => true },
      { key: 'annualRentCosts', label: 'Annual Rent Costs', help: 'Yearly rental-related costs (use K)', placeholder: '1K', visible: () => true },
      { key: 'otherAnnualCosts', label: 'Other Annual Costs (Rent)', help: 'Additional yearly costs for renting', placeholder: '500', visible: () => formInputs.scenario === 'buy_vs_rent' },

      { key: 'header_selling', label: '', help: '', placeholder: '', visible: () => true, isHeader: true, headerText: 'SELLING' },
      { key: 'includeSelling', label: 'Include Selling', help: 'yes/no - Enable selling analysis', placeholder: 'no', visible: () => formInputs.scenario === 'buy_vs_rent', toggleValues: ['yes', 'no'] },
      { key: 'agentCommission', label: 'Agent Commission (%)', help: 'Percentage of sale price', placeholder: '6', visible: () => true },
      { key: 'stagingCosts', label: 'Staging/Selling Costs', help: 'Fixed costs to sell (use K/M)', placeholder: '10K', visible: () => true },
      { key: 'taxFreeLimits', label: 'Tax-Free Gains Limit', help: 'Capital gains exempt (comma-separated, use K/M)', placeholder: '250K', visible: () => true },
      { key: 'capitalGainsTax', label: 'Capital Gains Tax (%)', help: 'Long-term capital gains rate', placeholder: '20', visible: () => true },
    ];
  }

  $: visibleFields = fields.filter(f => f.visible());

  function handleKeyDown(event: KeyboardEvent) {
    const currentField = visibleFields[currentFieldIndex];

    if (event.key === 'ArrowDown') {
      event.preventDefault();
      moveToNextField();
    } else if (event.key === 'ArrowUp') {
      event.preventDefault();
      moveToPreviousField();
    } else if (event.key === 'ArrowLeft' || event.key === 'ArrowRight') {
      // Toggle value for fields with toggleValues
      if (currentField && currentField.toggleValues && currentField.toggleValues.length === 2) {
        event.preventDefault();
        toggleFieldValue(currentField);
      }
    } else if (event.key === 'Enter') {
      event.preventDefault();
      if (event.ctrlKey) {
        handleSubmit();
      } else {
        // Regular Enter just moves to next field
        moveToNextField();
      }
    } else if (event.key === 'Tab') {
      event.preventDefault();
      if (event.shiftKey) {
        moveToPreviousField();
      } else {
        moveToNextField();
      }
    }
  }

  function toggleFieldValue(field: FormField) {
    if (!field.toggleValues || field.toggleValues.length !== 2) return;

    const currentValue = formInputs[field.key];
    const [value1, value2] = field.toggleValues;

    // Toggle to the other value
    formInputs[field.key] = currentValue === value1 ? value2 : value1;
  }

  function moveToNextField() {
    let nextIndex = currentFieldIndex + 1;
    // Skip headers
    while (nextIndex < visibleFields.length && visibleFields[nextIndex].isHeader) {
      nextIndex++;
    }
    if (nextIndex < visibleFields.length) {
      currentFieldIndex = nextIndex;
      focusCurrentField();
    }
  }

  function moveToPreviousField() {
    let prevIndex = currentFieldIndex - 1;
    // Skip headers
    while (prevIndex >= 0 && visibleFields[prevIndex].isHeader) {
      prevIndex--;
    }
    if (prevIndex >= 0) {
      currentFieldIndex = prevIndex;
      focusCurrentField();
    }
  }

  function focusCurrentField() {
    setTimeout(() => {
      const input = inputRefs[currentFieldIndex];
      if (input) {
        input.focus();
        input.select();
      }
    }, 0);
  }

  function handleFieldFocus(index: number) {
    currentFieldIndex = index;
  }

  function convertFormInputsToCalculatorInputs(): CalculatorInputs {
    return {
      scenario: formInputs.scenario as ScenarioType,
      inflationRate: parseFloat(formInputs.inflationRate) || 0,
      investmentReturnRate: parseFloat(formInputs.investmentReturnRate) || 0,
      include30Year: formInputs.include30Year === 'yes' || formInputs.include30Year === true,
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
      includeSelling: formInputs.includeSelling === 'yes' || formInputs.includeSelling === true,
      includeRentingSell: formInputs.includeRentingSell === 'yes' || formInputs.includeRentingSell === true,
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

  onMount(() => {
    // Find first non-header field to focus
    let firstFieldIndex = 0;
    while (firstFieldIndex < visibleFields.length && visibleFields[firstFieldIndex].isHeader) {
      firstFieldIndex++;
    }
    currentFieldIndex = firstFieldIndex;
    focusCurrentField();
  });
</script>

<div class="terminal-container font-mono">
  <div class="mb-6 border-l-2 border-monokai-purple pl-3 py-2">
    <div class="mb-1 text-monokai-text text-sm">
      <span class="text-monokai-purple">⌨</span> Navigation: <span class="text-monokai-cyan">↑↓</span> arrows to move | <span class="text-monokai-cyan">Ctrl+Enter</span> to calculate
    </div>
    <div class="text-monokai-text-muted text-xs">
      Field <span class="text-monokai-pink">{currentFieldIndex + 1}</span>/<span class="text-monokai-cyan">{visibleFields.length}</span>
    </div>
  </div>

  <form on:submit|preventDefault={handleSubmit} class="space-y-1">
    {#each visibleFields as field, index}
      {#if field.isHeader}
        <div class="section-header">
          <span class="text-monokai-orange">{field.headerText}</span>
        </div>
      {:else}
        <div class="terminal-field" class:focused={index === currentFieldIndex}>
          <div class="flex items-center gap-4">
            <div class="field-label w-72 flex-shrink-0">
              <span class="text-monokai-pink">{index === currentFieldIndex ? '>' : ' '}</span>
              <span class="ml-2" class:text-monokai-pink={index === currentFieldIndex} class:text-monokai-text={index !== currentFieldIndex}>{field.label}:</span>
            </div>
            <div class="field-input flex-1 min-w-0">
              <input
                type="text"
                bind:value={formInputs[field.key]}
                bind:this={inputRefs[index]}
                on:keydown={handleKeyDown}
                on:focus={() => handleFieldFocus(index)}
                placeholder={field.placeholder}
                class="terminal-input w-full"
              />
            </div>
            <div class="field-help w-96 flex-shrink-0 text-monokai-text-muted text-xs">
              {field.help}
            </div>
          </div>
        </div>
      {/if}
    {/each}

    <div class="mt-8 pt-4 border-t border-monokai-border">
      <button type="submit" class="terminal-button">
        <span class="text-monokai-green">$</span> ./calculate --run
      </button>
    </div>
  </form>
</div>

<style>
  .terminal-container {
    background: #000;
    padding: 2rem;
    border: 2px solid #2d2d2d;
    border-radius: 0.5rem;
  }

  .terminal-field {
    padding: 0.25rem 1rem;
    transition: all 0.15s;
    border-left: 3px solid transparent;
  }

  .terminal-field.focused {
    background-color: #0a0a0a;
    border-left-color: #FF6188;
  }

  .terminal-input {
    background: transparent;
    border: none;
    outline: none;
    color: #FCFCFA;
    font-family: 'Fira Code', monospace;
    font-size: 0.95rem;
    padding: 0.125rem 0;
  }

  .terminal-input::placeholder {
    color: #5c5c5c;
    font-style: italic;
  }

  .terminal-input:focus {
    border-bottom: 1px solid #FF6188;
  }

  .terminal-button {
    background: transparent;
    border: 2px solid #78DCE8;
    color: #78DCE8;
    padding: 0.75rem 2rem;
    border-radius: 0.375rem;
    font-family: 'Fira Code', monospace;
    cursor: pointer;
    transition: all 0.2s;
  }

  .terminal-button:hover {
    background: #78DCE8;
    color: #000;
  }

  .section-header {
    margin-top: 1rem;
    margin-bottom: 0.25rem;
    padding: 0.25rem 1rem;
    font-weight: bold;
    font-size: 0.875rem;
    letter-spacing: 0.05em;
    border-bottom: 1px solid #2d2d2d;
  }
</style>
