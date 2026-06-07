// Display formatters. Percentages take a FRACTION (0.247 -> "24.7%"), matching
// the backend convention.

const compactUSD = new Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
  notation: "compact",
  maximumFractionDigits: 2,
});

const priceUSD = new Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
  minimumFractionDigits: 2,
  maximumFractionDigits: 2,
});

export function formatUSDCompact(n: number): string {
  return compactUSD.format(n);
}

export function formatPrice(n: number): string {
  return priceUSD.format(n);
}

export function formatPct(fraction: number, digits = 1): string {
  return new Intl.NumberFormat("en-US", {
    style: "percent",
    maximumFractionDigits: digits,
  }).format(fraction);
}

export function formatNum(n: number, digits = 2): string {
  return new Intl.NumberFormat("en-US", { maximumFractionDigits: digits }).format(n);
}
