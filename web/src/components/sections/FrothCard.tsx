import type { Froth } from "../../types";
import { formatNum, formatPrice } from "../../format";

export function FrothCard({ data }: { data: Froth }) {
  return (
    <dl className="kv">
      <div>
        <dt>Beta</dt>
        <dd>{formatNum(data.beta)}</dd>
      </div>
      <div>
        <dt>PEG ratio</dt>
        <dd>{formatNum(data.pegRatio)}</dd>
      </div>
      <div>
        <dt>Trailing P/E</dt>
        <dd>{formatNum(data.trailingPE)}</dd>
      </div>
      <div>
        <dt>Forward P/E</dt>
        <dd>{formatNum(data.forwardPE)}</dd>
      </div>
      <div>
        <dt>52-week high</dt>
        <dd>{formatPrice(data.week52High)}</dd>
      </div>
      <div>
        <dt>52-week low</dt>
        <dd>{formatPrice(data.week52Low)}</dd>
      </div>
    </dl>
  );
}
