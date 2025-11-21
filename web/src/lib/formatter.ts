export function formatCurrency(amount: number, fullNumbers = false): string {
  const sign = amount < 0 ? '-' : '';
  const abs = Math.abs(amount);

  if (fullNumbers) {
    const formatted = abs.toFixed(1);
    const [intPart, decPart] = formatted.split('.');
    const withCommas = intPart.replace(/\B(?=(\d{3})+(?!\d))/g, ',');
    return `${sign}$${withCommas}.${decPart}`;
  }

  // Compact format
  if (abs >= 1000000) {
    return `${sign}${(abs / 1000000).toFixed(1)}M`;
  } else if (abs >= 1000) {
    return `${sign}${(abs / 1000).toFixed(1)}K`;
  } else {
    return `${sign}${abs.toFixed(1)}`;
  }
}

export function parseAmount(input: string): number {
  input = input.toLowerCase().trim();

  if (input === '') {
    return 0;
  }

  // Remove % sign if present
  input = input.replace(/%$/, '').trim();

  let multiplier = 1;
  let numStr = input;

  if (input.endsWith('k')) {
    multiplier = 1000;
    numStr = input.slice(0, -1);
  } else if (input.endsWith('m')) {
    multiplier = 1000000;
    numStr = input.slice(0, -1);
  } else if (input.endsWith('b')) {
    multiplier = 1000000000;
    numStr = input.slice(0, -1);
  }

  const value = parseFloat(numStr.trim());
  if (isNaN(value)) {
    return 0;
  }

  return value * multiplier;
}

export function parseDuration(duration: string): number {
  duration = duration.toLowerCase().trim();
  let years = 0;
  let months = 0;

  const yIndex = duration.indexOf('y');
  if (yIndex !== -1) {
    const yearStr = duration.substring(0, yIndex);
    years = parseInt(yearStr, 10);
    if (isNaN(years)) {
      throw new Error('Invalid year format');
    }
    duration = duration.substring(yIndex + 1);
  }

  const mIndex = duration.indexOf('m');
  if (mIndex !== -1) {
    const monthStr = duration.substring(0, mIndex);
    months = parseInt(monthStr, 10);
    if (isNaN(months)) {
      throw new Error('Invalid month format');
    }
  }

  const totalMonths = years * 12 + months;
  if (totalMonths <= 0) {
    throw new Error('Duration must be greater than 0');
  }

  return totalMonths;
}

export function parseAppreciationRates(input: string): number[] {
  input = input.trim();
  if (input === '') {
    return [0];
  }

  const parts = input.split(',');
  const rates: number[] = [];

  for (const part of parts) {
    const rate = parseAmount(part);
    rates.push(rate);
  }

  if (rates.length === 0) {
    return [0];
  }

  return rates;
}

export function formatPercent(value: number): string {
  return `${value.toFixed(2)}%`;
}
