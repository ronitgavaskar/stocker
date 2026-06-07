import type { Overview } from "../../types";
import { formatUSDCompact } from "../../format";

// data is GUARANTEED non-null: SectionCard narrowed it once at the top.
export function OverviewCard({ data }: { data: Overview }) {
  return (
    <dl className="kv">
      <div>
        <dt>Name</dt>
        <dd>{data.name}</dd>
      </div>
      <div>
        <dt>Sector</dt>
        <dd>{data.sector}</dd>
      </div>
      <div>
        <dt>Industry</dt>
        <dd>{data.industry}</dd>
      </div>
      <div>
        <dt>Market cap</dt>
        <dd>{formatUSDCompact(data.marketCap)}</dd>
      </div>
      <div className="kv__full">
        <dt>Summary</dt>
        <dd>{data.summary}</dd>
      </div>
    </dl>
  );
}
