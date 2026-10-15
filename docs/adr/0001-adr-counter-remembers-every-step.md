# ADR-0001: The counter remembers every step

---

- **Status**: Accepted
- **Author(s)**: Mikael Grön
- **Created**: 2026-10-15
- **Deciders**: Mikael Grön
- **Supersedes**: N/A
- **Superseded by**: N/A

---

## 1. Context

A plain tally answers "how many" but not "what happened". A consumer
that needs the increments with their order has no way to get them, and
adding history later to call sites that never recorded it is guesswork.

## 2. Decision

The counter keeps a step log. Every call to Inc MUST append exactly one
entry to the log, and each entry MUST carry its ordinal position and a
timestamp of the increment it records. Log MUST return every entry in
order, oldest first.

## 3. Alternatives Considered

- Growing only the integer count. It is the state of affairs today and
  answers fewer questions each year.

## 4. Consequences

The Counter struct gains a slice it did not have, and Log is a read-only
view: callers get a copy, never the underlying slice.
