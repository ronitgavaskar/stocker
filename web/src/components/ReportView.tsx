import type { Report } from "../types";
import { SectionCard } from "./SectionCard";

export function ReportView({ report }: { report: Report }) {
  // Pull the company name from the overview section for the header subtitle.
  const overview = report.sections.find((s) => s.kind === "overview");
  const companyName =
    overview && overview.kind === "overview" && overview.data
      ? overview.data.name
      : null;
  const date = new Date(report.generatedAt).toLocaleDateString();

  return (
    <section className="report-card">
      <header className="report-card__head">
        <h1>{report.ticker}</h1>
        <p className="report-card__sub">
          {companyName ? `${companyName} · ` : ""}research note · {date}
        </p>
      </header>
      <hr className="report-card__divider" />
      <div className="report-grid">
        {report.sections.map((section) => (
          <SectionCard key={section.kind} section={section} />
        ))}
      </div>
    </section>
  );
}
