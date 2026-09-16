# Story — Persisted editable greeting

Module: `greeting`
Plan item: Persisted editable greeting

## User story

As a Visitor, I want to view and update the shared stored greeting, so that the page shows the current PostgreSQL-backed greeting after save and reload.

## In scope

- Render one centered greeting page with one `h1`, one `Greeting` text input, one `Save` submit button, and one polite live message area.
- Read current greeting from Go backend API backed by PostgreSQL.
- Seed a new environment with `Hello, World!` when no visitor change exists.
- Let any Visitor save one shared non-empty greeting.
- Trim leading and trailing whitespace before saving.
- Update heading, input, and form message after successful save without requiring reload.
- Persist saved greeting so reload shows saved value.
- Reject empty or whitespace-only input without saving and return focus to input.
- Match approved minimal visual design and responsive control layout.

## Out of scope

- Sign-in, roles, per-user greetings, or permissions; this app has public Visitor access only.
- Multiple greetings, greeting history, audit log, or conflict UI; last completed save wins.
- Navigation, extra sections, secondary pages, animation, hover/active/disabled visual variants, or external services.
- Custom API/database outage UI; service failure behavior stays in backend service contract, and UI must not show unsaved changes as saved.
- Changing design tokens, global styling, architecture decisions, or project scaffold.

## UI scope

- Screen: Greeting page, `main > .greeting-section`, default state from approved design.
- Elements: one large centered black `h1` with current greeting, visually hidden label `Greeting`, text input, blue `#2563EB` `Save` submit button, and polite live message below controls.
- Layout: white page, black text, no animation; above 520px input and button align in one row, at 520px and below they stack vertically and span available section width.
- Interaction states: default, focus-visible for input/button, success message `Saved.`, validation message `Enter a greeting.`.

## Acceptance criteria

- SC-1 [GREETING-001 AC-1]: Given new environment has no previously saved visitor change, when Visitor opens page, heading text is `Hello, World!`.
- SC-2 [GREETING-001 AC-2]: Given stored greeting is `Hello, World!`, when Visitor opens page, text input value is `Hello, World!`.
- SC-3 [GREETING-001 AC-3]: Given stored greeting is `Pipeline accepted`, when Visitor opens page, heading text is `Pipeline accepted`.
- SC-4 [GREETING-001 AC-4]: Given stored greeting is `Pipeline accepted`, when Visitor opens page, text input value is `Pipeline accepted`.
- SC-5 [GREETING-002 AC-1]: Given page is open and input value is `Hello, World!`, when Visitor changes input to `Pipeline accepted` and clicks `Save`, heading text becomes `Pipeline accepted` without reload.
- SC-6 [GREETING-002 AC-2]: Given page is open and input value is `Pipeline accepted`, when Visitor clicks `Save`, form message text becomes `Saved.`.
- SC-7 [GREETING-002 AC-3]: Given Visitor has saved `Pipeline accepted`, when Visitor reloads page, heading text is `Pipeline accepted`.
- SC-8 [GREETING-002 AC-4]: Given Visitor has saved `Pipeline accepted`, when Visitor reloads page, text input value is `Pipeline accepted`.
- SC-9 [GREETING-002 AC-5]: Given page is open and input value is `  Trim me  `, when Visitor clicks `Save`, heading text becomes `Trim me`.
- SC-10 [GREETING-003 AC-1]: Given heading text is `Hello, World!` and input value is empty, when Visitor clicks `Save`, heading text remains `Hello, World!`.
- SC-11 [GREETING-003 AC-2]: Given input value is empty, when Visitor clicks `Save`, form message text becomes `Enter a greeting.`.
- SC-12 [GREETING-003 AC-3]: Given input value is spaces only, when Visitor clicks `Save`, form message text becomes `Enter a greeting.`.
- SC-13 [GREETING-003 AC-4]: Given input value is empty, when Visitor clicks `Save`, keyboard focus is on the `Greeting` input.
- SC-14 [GREETING-004 AC-1]: Given Visitor opens page, when page renders, exactly one `h1` is present and it contains current greeting.
- SC-15 [GREETING-004 AC-2]: Given Visitor opens page, when page renders, a text input has accessible label `Greeting`.
- SC-16 [GREETING-004 AC-3]: Given Visitor opens page, when page renders, a submit button has visible text `Save`.
- SC-17 [GREETING-004 AC-4]: Given Visitor opens page, when page renders, a polite live message region exists below form controls.
- SC-18 [GREETING-004 AC-5]: Given keyboard user tabs through controls, when input or button receives focus, visible black focus outline appears on focused control.
- SC-19 [GREETING-004 AC-6]: Given viewport width is 1280px, when page renders, input and Save button are aligned in one horizontal row.
- SC-20 [GREETING-004 AC-7]: Given viewport width is 390px, when page renders, input and Save button are stacked vertically and each spans available section width.
- SC-21 [GREETING-004 AC-8]: Given Visitor opens page, when page renders, page background is white, text is black, and Save button background is `#2563EB`.

## Dependencies

- PostgreSQL storage exists and can persist one shared greeting row.
- Go backend API exposes greeting read and save operations under `/v1/...` contract.
- Backend startup applies migrations and seeds `Hello, World!` for new environments.
- Next.js App Router shell exists and mounts this story component from `app/page.tsx`.
- Frontend can reach backend via configured `API_ORIGIN` for server reads and `NEXT_PUBLIC_API_URL` for browser saves.
- No external accounts, secrets, sign-in, or stakeholder decisions required.
