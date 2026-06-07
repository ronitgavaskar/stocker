import type { Financial } from "../../types";
import { formatUSDCompact, formatPct, formatNum } from "../../format";

export function FinancialCard({ data }: { data: Financial }) {
  return (
    <dl className="kv">
      <div>
        <dt>Revenue (TTM)</dt>
        <dd>{formatUSDCompact(data.revenueTTM)}</dd>
      </div>
      <div>
        <dt>Gross profit (TTM)</dt>
        <dd>{formatUSDCompact(data.grossProfitTTM)}</dd>
      </div>
      <div>
        <dt>Profit margin</dt>
        <dd>{formatPct(data.profitMargin)}</dd>
      </div>
      <div>
        <dt>Return on equity</dt>
        <dd>{formatPct(data.returnOnEquityTTM)}</dd>
      </div>
      <div>
        <dt>EPS</dt>
        <dd>{formatNum(data.eps)}</dd>
      </div>
      <div>
        <dt>EBITDA</dt>
        <dd>{formatUSDCompact(data.ebitda)}</dd>
      </div>
    </dl>
  );
}
