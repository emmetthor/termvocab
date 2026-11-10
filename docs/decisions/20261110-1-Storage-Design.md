# ADR 20261110-1: Storage Design

## Context

TermVocab requires persistent vocabulary storage.

The application should work without an internet connection.

## Decision

Use SQLite with Go's database/sql package.

## Reasons

- No database server required
- Local-first
- Easy to distribute
- Supports transactions

## Consequences

- Simple installation
- No built-in cross-device synchronization
- Database migrations are required when schema changes
