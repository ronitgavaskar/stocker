// Mirrors the backend's JSON shape (camelCase). The Section type is a strict
// discriminated union: `data`'s type is tied to `kind`, and `data` is null when
// a section failed. The compiler is the safety net — you cannot read data fields
// without first narrowing both kind and null.

export type Status = "ok" | "flagged" | "failed" | "unknown";

export interface Overview {
  ticker: string;
  name: string;
  summary: string;
  sector: string;
  industry: string;
  marketCap: number;
}

export interface Financial {
  revenueTTM: number;
  grossProfitTTM: number;
  profitMargin: number; // fraction, e.g. 0.247 = 24.7%
  returnOnEquityTTM: number; // fraction, e.g. 0.247 = 24.7%
  eps: number;
  ebitda: number;
}

export interface AnalystRatings {
  strongBuy: number;
  buy: number;
  hold: number;
  sell: number;
  strongSell: number;
}

export interface Ownership {
  percentInsiders: number; // fraction, e.g. 0.247 = 24.7%
  percentInstitutions: number; // fraction, e.g. 0.247 = 24.7%
  analystTargetPrice: number;
  ratings: AnalystRatings;
}

export interface Froth {
  beta: number;
  pegRatio: number;
  trailingPE: number;
  forwardPE: number;
  week52High: number;
  week52Low: number;
}

// Discriminated union on `kind`. `data` is null on failure (never {}).
export type Section =
  | { kind: "overview"; status: Status; note: string; data: Overview | null }
  | { kind: "financial"; status: Status; note: string; data: Financial | null }
  | { kind: "ownership"; status: Status; note: string; data: Ownership | null }
  | { kind: "froth"; status: Status; note: string; data: Froth | null };

export type SectionKind = Section["kind"];

export interface Report {
  ticker: string;
  sections: Section[];
  generatedAt: string;
}
