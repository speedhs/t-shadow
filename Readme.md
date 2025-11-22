# t-shadow

t-shadow (v1)
Ghost-table based, low-downtime schema/type migration tool for PostgreSQL

t-shadow is a lightweight PostgreSQL schema migration utility designed to safely change column types on large tables with minimal blocking.
It uses a shadow table, logical decoding, and a batched backfill process to migrate data without rewriting the entire source table under a lock.

This is v1, which focuses on correctness and simplicity.
Future versions will implement non-blocking cut-over logic and safe lock acquisition.

### Key Features (v1)

Zero-lock backfill using batched reads

Real-time change propagation via PostgreSQL logical decoding

Shadow table creation mirroring the original table

Consistent state transfer using WAL→shadow replay

Atomic cut-over using ALTER TABLE RENAME

Minimal downtime window (single ACCESS EXCLUSIVE lock during final swap)

### What t-shadow solves

PostgreSQL does not allow changing column types efficiently on large tables without a full rewrite.
t-shadow provides a safe workflow:

Create a shadow table with the desired schema

Copy existing rows in batches

Stream ongoing writes from WAL to the shadow table

Lock table briefly

Replay final changes

Swap tables atomically

This avoids long locks and downtime.