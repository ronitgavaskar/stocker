import type { Section, SectionKind } from "../types";
import { OverviewCard } from "./sections/OverviewCard";
import { FinancialCard } from "./sections/FinancialCard";
import { OwnershipCard } from "./sections/OwnershipCard";
import { FrothCard } from "./sections/FrothCard";

const TITLES: Record<SectionKind, string> = {
  overview: "Overview",
  financial: "Financial health",
  ownership: "Ownership & sentiment",
  froth: "Froth check",
};

export function SectionCard({ section }: { section: Section }) {
  const title = TITLES[section.kind];

  // THE worst-case gate, handled ONCE and structurally: a failed section (or any
  // section whose data is null) never reaches a per-kind card. Below this line
  // the compiler knows section.data is non-null.
  if (section.status === "failed" || section.data === null) {
    return (
      <article className="card card--failed">
        <h2 className="card__title">{title}</h2>
        <p className="card__note">
          Unavailable — {section.note || "data could not be loaded"}
        </p>
      </article>
    );
  }

  const flagged = section.status === "flagged";

  // Per-kind cards receive GUARANTEED non-null data — they cannot be written to
  // handle a missing-data case, because the type proves it exists.
  let body;
  switch (section.kind) {
    case "overview":
      body = <OverviewCard data={section.data} />;
      break;
    case "financial":
      body = <FinancialCard data={section.data} />;
      break;
    case "ownership":
      body = <OwnershipCard data={section.data} />;
      break;
    case "froth":
      body = <FrothCard data={section.data} />;
      break;
  }

  return (
    <article className={`card card--${flagged ? "flagged" : "ok"}`}>
      <h2 className="card__title">{title}</h2>
      {flagged && <p className="card__note">{section.note}</p>}
      {body}
    </article>
  );
}
