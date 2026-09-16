# Test cases — Persisted editable greeting

Module: `greeting`
Function: Persisted editable greeting
Story: `docs/greeting/stories/persisted-editable-greeting.md`
Service contract: `docs/architecture/services.md`

Risk level: Medium. This is a small public single-page flow, but it writes shared persisted state through frontend, backend, and PostgreSQL. Coverage includes happy path, validation, persistence, responsive/accessibility design, and existing service contract failures.

## Page rendering and stored greeting

**Scenario**: New environment shows seeded heading
**Given**: Database has no previously saved visitor change and service has completed startup seeding.
**When**: Visitor opens the greeting page.
**Then**: Exactly one `h1` is visible and its text is `Hello, World!`.
Traces: SC-1 (GREETING-001 AC-1), SC-14 (GREETING-004 AC-1)
Check: render_url

**Scenario**: Seeded greeting fills input
**Given**: Stored greeting is `Hello, World!`.
**When**: Visitor opens the greeting page.
**Then**: The text input with accessible label `Greeting` has value `Hello, World!`.
Traces: SC-2 (GREETING-001 AC-2), SC-15 (GREETING-004 AC-2)
Check: render_url

**Scenario**: Existing saved greeting shows as heading
**Given**: Stored greeting is `Pipeline accepted`.
**When**: Visitor opens the greeting page.
**Then**: Exactly one `h1` is visible and its text is `Pipeline accepted`.
Traces: SC-3 (GREETING-001 AC-3), SC-14 (GREETING-004 AC-1)
Check: render_url

**Scenario**: Existing saved greeting fills input
**Given**: Stored greeting is `Pipeline accepted`.
**When**: Visitor opens the greeting page.
**Then**: The text input with accessible label `Greeting` has value `Pipeline accepted`.
Traces: SC-4 (GREETING-001 AC-4), SC-15 (GREETING-004 AC-2)
Check: render_url

## Save interactions

**Scenario**: Save updates heading without reload
**Given**: Page is open, stored greeting is `Hello, World!`, heading text is `Hello, World!`, and input value is `Hello, World!`.
**When**: Visitor changes the `Greeting` input to `Pipeline accepted` and clicks `Save`.
**Then**: Heading text becomes `Pipeline accepted` before any page reload.
Traces: SC-5 (GREETING-002 AC-1)
Check: interact_page

**Scenario**: Save shows success message
**Given**: Page is open and input value is `Pipeline accepted`.
**When**: Visitor clicks `Save`.
**Then**: The polite form message text becomes `Saved.`.
Traces: SC-6 (GREETING-002 AC-2)
Check: interact_page

**Scenario**: Saved heading persists after reload
**Given**: Visitor has saved `Pipeline accepted` successfully.
**When**: Visitor reloads the page.
**Then**: Heading text is `Pipeline accepted`.
Traces: SC-7 (GREETING-002 AC-3)
Check: interact_page

**Scenario**: Saved input value persists after reload
**Given**: Visitor has saved `Pipeline accepted` successfully.
**When**: Visitor reloads the page.
**Then**: The text input with accessible label `Greeting` has value `Pipeline accepted`.
Traces: SC-8 (GREETING-002 AC-4)
Check: interact_page

**Scenario**: Save trims leading and trailing whitespace
**Given**: Page is open and input value is `  Trim me  `.
**When**: Visitor clicks `Save`.
**Then**: Heading text becomes `Trim me`, and the input value becomes `Trim me`.
Traces: SC-9 (GREETING-002 AC-5)
Check: interact_page

**Scenario**: Native form submit saves greeting from input
**Given**: Page is open and input value is `Hello, World!`.
**When**: Visitor changes the `Greeting` input to `Submitted by Enter` and presses Enter in the input.
**Then**: Heading text becomes `Submitted by Enter` and the polite form message text becomes `Saved.`.
Traces: SC-5 (GREETING-002 AC-1), SC-6 (GREETING-002 AC-2)
Check: interact_page

**Scenario**: Last completed save wins for shared greeting
**Given**: Two visitors have the page open, Visitor A saves `First save`, and Visitor B saves `Second save` after Visitor A save completes.
**When**: A visitor opens or reloads the greeting page.
**Then**: Heading text is `Second save` and the input value is `Second save`.
Traces: SC-7 (GREETING-002 AC-3), SC-8 (GREETING-002 AC-4)
Check: interact_page

## Validation and recovery

**Scenario**: Empty greeting is not saved
**Given**: Stored greeting is `Hello, World!`, page is open, heading text is `Hello, World!`, and input value is empty.
**When**: Visitor clicks `Save`.
**Then**: Heading text remains `Hello, World!`.
Traces: SC-10 (GREETING-003 AC-1)
Check: interact_page

**Scenario**: Empty greeting shows validation message
**Given**: Page is open and input value is empty.
**When**: Visitor clicks `Save`.
**Then**: The polite form message text becomes `Enter a greeting.`.
Traces: SC-11 (GREETING-003 AC-2)
Check: interact_page

**Scenario**: Spaces-only greeting shows validation message
**Given**: Page is open and input value is spaces only.
**When**: Visitor clicks `Save`.
**Then**: The polite form message text becomes `Enter a greeting.`.
Traces: SC-12 (GREETING-003 AC-3)
Check: interact_page

**Scenario**: Empty greeting returns focus to input
**Given**: Page is open and input value is empty.
**When**: Visitor clicks `Save`.
**Then**: Keyboard focus is on the text input with accessible label `Greeting`.
Traces: SC-13 (GREETING-003 AC-4)
Check: interact_page

**Scenario**: Failed save does not show success
**Given**: Page is open, heading text is `Hello, World!`, and backend API is unreachable for browser saves.
**When**: Visitor changes the `Greeting` input to `Unsaved outage` and clicks `Save`.
**Then**: Heading text remains `Hello, World!` and the polite form message text is not `Saved.`.
Traces: SC-5 (GREETING-002 AC-1), SC-6 (GREETING-002 AC-2)
Check: interact_page

## Approved controls, layout, and visual style

**Scenario**: Save button is visible submit control
**Given**: Visitor opens the greeting page.
**When**: Page renders.
**Then**: A submit button is visible and its text is exactly `Save`.
Traces: SC-16 (GREETING-004 AC-3)
Check: render_url

**Scenario**: Polite live message region exists below controls
**Given**: Visitor opens the greeting page.
**When**: Page renders.
**Then**: A polite live message region exists below the input and `Save` button controls.
Traces: SC-17 (GREETING-004 AC-4)
Check: render_url

**Scenario**: Input focus outline is visible and black
**Given**: Visitor opens the greeting page.
**When**: Keyboard user tabs until the `Greeting` input receives focus.
**Then**: Focused input has a visible black focus outline.
Traces: SC-18 (GREETING-004 AC-5)
Check: interact_page

**Scenario**: Button focus outline is visible and black
**Given**: Visitor opens the greeting page.
**When**: Keyboard user tabs until the `Save` button receives focus.
**Then**: Focused button has a visible black focus outline.
Traces: SC-18 (GREETING-004 AC-5)
Check: interact_page

**Scenario**: Desktop controls align in one row
**Given**: Viewport width is 1280px.
**When**: Page renders.
**Then**: The `Greeting` input and `Save` button share one horizontal row.
Traces: SC-19 (GREETING-004 AC-6)
Check: measure_styles

**Scenario**: Mobile controls stack and fill section width
**Given**: Viewport width is 390px.
**When**: Page renders.
**Then**: The `Greeting` input is above the `Save` button, and each spans the available section width.
Traces: SC-20 (GREETING-004 AC-7)
Check: measure_styles

**Scenario**: Approved colors render
**Given**: Visitor opens the greeting page.
**When**: Page renders.
**Then**: Page background is `#FFFFFF`, page text is `#000000`, and `button[type="submit"]` background is `#2563EB`.
Traces: SC-21 (GREETING-004 AC-8)
Check: measure_styles

**Scenario**: Long greeting wraps without horizontal page scroll
**Given**: Stored greeting is `This is a deliberately long greeting that should wrap inside the centered section instead of creating horizontal page scroll across the viewport.`
**When**: Visitor opens the greeting page at 390px viewport width.
**Then**: Heading text is visible across multiple lines within the centered section and document horizontal scroll width does not exceed viewport width.
Traces: SC-20 (GREETING-004 AC-7)
Check: measure_styles

## Backend service contract

**Scenario**: GET greeting returns current greeting
**Given**: Stored greeting is `Hello, World!`.
**When**: Client requests `GET /v1/greeting`.
**Then**: Response status is `200` and body is exactly `{"text":"Hello, World!"}`.
Traces: contract (GET /v1/greeting)
Check: fetch_url

**Scenario**: GET greeting returns saved text exactly
**Given**: Stored greeting is `Pipeline accepted`.
**When**: Client requests `GET /v1/greeting`.
**Then**: Response status is `200` and body is exactly `{"text":"Pipeline accepted"}`.
Traces: contract (GET /v1/greeting)
Check: fetch_url

**Scenario**: GET greeting query failure returns internal error
**Given**: Greeting query fails while backend remains reachable.
**When**: Client requests `GET /v1/greeting`.
**Then**: Response status is `500` and body is exactly `{"error":{"code":"INTERNAL","message":"Internal server error."}}`.
Traces: contract (GET /v1/greeting)
Check: manual

**Scenario**: PUT greeting stores and returns trimmed text
**Given**: Backend is reachable and stored greeting is `Hello, World!`.
**When**: Client requests `PUT /v1/greeting` with JSON body `{"text":"  Pipeline accepted  "}`.
**Then**: Response status is `200` and body is exactly `{"text":"Pipeline accepted"}`; a later `GET /v1/greeting` returns `{"text":"Pipeline accepted"}`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects malformed JSON
**Given**: Backend is reachable.
**When**: Client requests `PUT /v1/greeting` with body `{`.
**Then**: Response status is `400` and body is exactly `{"error":{"code":"MALFORMED_REQUEST","message":"Request body is malformed."}}`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects wrong field type
**Given**: Backend is reachable.
**When**: Client requests `PUT /v1/greeting` with JSON body `{"text":123}`.
**Then**: Response status is `400` and body is exactly `{"error":{"code":"MALFORMED_REQUEST","message":"Request body is malformed."}}`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects unknown field
**Given**: Backend is reachable.
**When**: Client requests `PUT /v1/greeting` with JSON body `{"text":"Pipeline accepted","extra":"ignored?"}`.
**Then**: Response status is `400` and body is exactly `{"error":{"code":"MALFORMED_REQUEST","message":"Request body is malformed."}}`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects empty trimmed text
**Given**: Backend is reachable and stored greeting is `Hello, World!`.
**When**: Client requests `PUT /v1/greeting` with JSON body `{"text":"   "}`.
**Then**: Response status is `422` and body is exactly `{"error":{"code":"VALIDATION_FAILED","message":"Greeting must not be empty."}}`; a later `GET /v1/greeting` returns `{"text":"Hello, World!"}`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting query failure returns internal error
**Given**: Greeting update query fails while backend remains reachable.
**When**: Client requests `PUT /v1/greeting` with JSON body `{"text":"Pipeline accepted"}`.
**Then**: Response status is `500` and body is exactly `{"error":{"code":"INTERNAL","message":"Internal server error."}}`.
Traces: contract (PUT /v1/greeting)
Check: manual

**Scenario**: GET healthz returns ok after migrations and database check
**Given**: Backend startup has applied migrations and `SELECT 1` succeeds.
**When**: Client requests `GET /healthz`.
**Then**: Response status is `200` and body is exactly `{"status":"ok"}`.
Traces: contract (GET /healthz)
Check: fetch_url

**Scenario**: GET healthz returns unavailable when database dependency is unavailable
**Given**: Database dependency is unavailable during health check.
**When**: Client requests `GET /healthz`.
**Then**: Response status is `503` and body is exactly `{"error":{"code":"UNAVAILABLE","message":"Service unavailable."}}`.
Traces: contract (GET /healthz)
Check: manual

## Coverage check

Story criteria covered: SC-1 through SC-21.
Contract coverage included: `GET /v1/greeting` success and `500`; `PUT /v1/greeting` success, malformed JSON, wrong field type, unknown field, empty validation, and `500`; `GET /healthz` success and `503`.
Manual cases limited to database/query failure states that cannot be forced by page/browser/HTTP instruments without a database-refusal or query-failure control.
