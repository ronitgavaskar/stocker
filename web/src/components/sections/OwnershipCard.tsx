import type { Ownership } from "../../types";
import { formatPct, formatPrice } from "../../format";

export function OwnershipCard({ data }: { data: Ownership }) {
  const r = data.ratings;
  const rows: Array<[string, number]> = [
    ["Strong buy", r.strongBuy],
    ["Buy", r.buy],
    ["Hold", r.hold],
    ["Sell", r.sell],
    ["Strong sell", r.strongSell],
  ];

  return (
    <div>
      <dl className="kv">
        <div>
          <dt>Insiders</dt>
          <dd>{formatPct(data.percentInsiders, 2)}</dd>
        </div>
        <div>
          <dt>Institutions</dt>
          <dd>{formatPct(data.percentInstitutions)}</dd>
        </div>
        <div>
          <dt>Analyst target</dt>
          <dd>{formatPrice(data.analystTargetPrice)}</dd>
        </div>
      </dl>
      <div className="ratings">
        <span className="ratings__label">Analyst ratings</span>
        <ul className="ratings__list">
          {rows.map(([label, n]) => (
            <li key={label}>
              <span>{label}</span>
              <strong>{n}</strong>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}
