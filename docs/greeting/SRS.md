# SRS — Greeting

Module: `greeting`
Design: [View the approved design](http://localhost:8080/design/6d9a6e93-f5de-4153-bcf6-61e9579ad3f2)
Design system: `design/design-system.md`

> One file per module, at `docs/greeting/SRS.md`. It covers only the functions
> that belong to this module. Never write `docs/SRS.md`.

## 1. Purpose

The greeting module lets any visitor view and update the single greeting for "Hello World Acceptance 9". It proves the full path from PostgreSQL storage through the Go API to the Next.js page. Without this module, the product is only static text and cannot satisfy the pipeline acceptance goal.

## 2. Actors

| Actor | Who they are | What they may do in this module |
|---|---|---|
| Visitor | Anyone opening the page; no sign-in exists | View the stored greeting, edit the greeting text, and save the new greeting |

## 3. Scope

**In scope** — the functions specified below, by their plan titles:

- Persisted editable greeting

**Out of scope** — name what a reader would reasonably expect here and say where it lives instead. This section prevents the same argument twice.

- Sign-in and roles — deliberately not built; the approved scope has no authentication or permissions beyond public visitor access.
- Navigation or extra sections — deliberately not built; approved structure is one centered greeting section only.
- External services — deliberately not built; stakeholder specified no external services.
- Multiple greetings or history — deliberately not built; approved scope has one stored greeting only.

## 4. Functional requirements

### 4.1 Persisted editable greeting

**Requirement GREETING-001 — Show stored greeting**

*As a* visitor, *I want to* see the stored greeting as the page heading, *so that* the page reflects the value persisted for the product.

Behaviour:

1. When a visitor opens the greeting page, the page requests the current greeting from the backend service.
2. The page shows the current greeting as the only `h1` heading.
3. The initial stored greeting for a new environment is `Hello, World!`.
4. The text input starts with the same current greeting value shown in the heading.

**Acceptance criteria** — each is proved by at least one test case in `docs/greeting/test-cases/persisted-editable-greeting.md`, through the story plan that cites it (`SC-1 [GREETING-001 AC-1]`). Given/When/Then, no compound conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | A new environment has no previously saved visitor change | Visitor opens the page | The heading text is `Hello, World!` |
| AC-2 | The stored greeting is `Hello, World!` | Visitor opens the page | The text input value is `Hello, World!` |
| AC-3 | The stored greeting is `Pipeline accepted` | Visitor opens the page | The heading text is `Pipeline accepted` |
| AC-4 | The stored greeting is `Pipeline accepted` | Visitor opens the page | The text input value is `Pipeline accepted` |

**Requirement GREETING-002 — Save greeting change**

*As a* visitor, *I want to* enter new greeting text and save it, *so that* the heading and stored greeting change to my text.

Behaviour:

1. Visitor edits the `Greeting` text field.
2. Visitor submits the form by clicking `Save` or using native form submission from the input.
3. If trimmed text is non-empty, the page saves that value through the backend service.
4. After a successful save, the heading changes to the saved value without requiring a page reload.
5. After a successful save, the input value changes to the saved value.
6. After a successful save, the live form message shows `Saved.`.
7. After a page reload, the heading and input still show the saved value.

**Acceptance criteria** — each is proved by at least one test case in `docs/greeting/test-cases/persisted-editable-greeting.md`, through the story plan that cites it (`SC-1 [GREETING-002 AC-1]`).

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Page is open and input value is `Hello, World!` | Visitor changes input to `Pipeline accepted` and clicks `Save` | Heading text becomes `Pipeline accepted` |
| AC-2 | Page is open and input value is `Pipeline accepted` | Visitor clicks `Save` | Form message text becomes `Saved.` |
| AC-3 | Visitor has saved `Pipeline accepted` | Visitor reloads the page | Heading text is `Pipeline accepted` |
| AC-4 | Visitor has saved `Pipeline accepted` | Visitor reloads the page | Text input value is `Pipeline accepted` |
| AC-5 | Page is open and input value is `  Trim me  ` | Visitor clicks `Save` | Heading text becomes `Trim me` |

**Requirement GREETING-003 — Reject empty greeting**

*As a* visitor, *I want to* be told when the greeting is empty, *so that* I do not save a blank heading.

Behaviour:

1. Visitor submits the form with a value that is empty after trimming whitespace.
2. The page does not save the value.
3. The heading remains unchanged.
4. The input receives focus.
5. The live form message shows `Enter a greeting.`.

**Acceptance criteria** — each is proved by at least one test case in `docs/greeting/test-cases/persisted-editable-greeting.md`, through the story plan that cites it (`SC-1 [GREETING-003 AC-1]`).

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Heading text is `Hello, World!` and input value is empty | Visitor clicks `Save` | Heading text remains `Hello, World!` |
| AC-2 | Input value is empty | Visitor clicks `Save` | Form message text becomes `Enter a greeting.` |
| AC-3 | Input value is spaces only | Visitor clicks `Save` | Form message text becomes `Enter a greeting.` |
| AC-4 | Input value is empty | Visitor clicks `Save` | Keyboard focus is on the `Greeting` input |

**Requirement GREETING-004 — Match approved screen controls**

*As a* visitor, *I want to* use the approved greeting form controls, *so that* the page matches the accepted design and remains accessible.

Behaviour:

1. The page has one centered greeting section on a white background.
2. The greeting appears as a large black heading.
3. The form contains one text input labelled `Greeting`; the label is available to assistive technology and visually hidden.
4. The form contains one primary action button labelled `Save`.
5. The form contains one polite live message area under the controls.
6. The input and button show visible keyboard focus outlines.
7. At viewport widths above 520px, the input and button sit in one row.
8. At viewport widths of 520px and below, the input and button stack vertically and fill the section width.

**Acceptance criteria** — each is proved by at least one test case in `docs/greeting/test-cases/persisted-editable-greeting.md`, through the story plan that cites it (`SC-1 [GREETING-004 AC-1]`).

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Visitor opens the page | Page renders | Exactly one `h1` is present and it contains the current greeting |
| AC-2 | Visitor opens the page | Page renders | A text input has accessible label `Greeting` |
| AC-3 | Visitor opens the page | Page renders | A submit button has visible text `Save` |
| AC-4 | Visitor opens the page | Page renders | A polite live message region exists below the form controls |
| AC-5 | Keyboard user tabs through controls | Input or button receives focus | A visible black focus outline appears on the focused control |
| AC-6 | Viewport width is 1280px | Page renders | Input and Save button are aligned in one horizontal row |
| AC-7 | Viewport width is 390px | Page renders | Input and Save button are stacked vertically and each spans available section width |
| AC-8 | Visitor opens the page | Page renders | Page background is white, text is black, and Save button background is `#2563EB` |

**Failure, boundary and permission behaviour** — the part most often skipped and most often the source of bugs. Every case this function actually has needs a defined outcome; "should not happen" is not an outcome.

| Case | Condition | Expected behaviour |
|---|---|---|
| Invalid input | Greeting is empty or whitespace after trimming | Nothing is saved; heading remains unchanged; message shows `Enter a greeting.`; focus moves to the input |
| Boundary | Greeting contains leading or trailing whitespace | Saved greeting trims leading and trailing whitespace |
| Boundary | Greeting is long enough to exceed the heading width | Heading wraps within the centered section; no horizontal page scroll appears |
| Not found | Stored greeting row does not exist in a new environment | Visitor sees and can edit seeded greeting `Hello, World!` |
| Not permitted | Visitor is not signed in | Not applicable: no sign-in or restricted role exists; all visitors may view and save the greeting |
| Conflict | Two visitors save different non-empty greetings | Last completed save becomes the stored greeting shown on subsequent loads |
| Upstream failure | Backend API or database is unavailable | Approved design has no error screen; service error envelope and HTTP status are specified in the service contract, and no partial greeting change is shown as saved |

**Data touched** — the fields this function reads and writes, in product terms. The physical schema is TL's job in `docs/architecture/erd.md`; this is the list that document has to satisfy.

| Field | Type | Required | Rule |
|---|---|---|---|
| Greeting text | text | yes | Initial value is `Hello, World!`; saved value is trimmed; empty trimmed value is rejected; exact saved text is returned on later loads |

## 5. Screens

The design is the source of truth for appearance; this section maps functions onto it so nothing in the design is unaccounted for and nothing specified here is missing from the design.

List only the states the approved design actually shows. A screen the design draws once, with no variant for waiting, for no data, or for a failure, has exactly **one** state and its name is `default`. That is not a placeholder and not an invented state: it is what "this screen has one appearance" is called, and it is the correct and complete answer for a static screen.

| Screen | Section in the design | Functions it serves | States that must exist |
|---|---|---|---|
| Greeting page | `main > .greeting-section` in approved design | GREETING-001, GREETING-002, GREETING-003, GREETING-004 | default |

Default state elements shown by approved design and covered above: one large greeting heading, visually hidden `Greeting` label, text input, `Save` submit button, and live form message area.

## 6. Non-functional requirements

Only what is real for this module. Delete rows that do not apply rather than inventing a number nobody will check.

| Area | Requirement |
|---|---|
| Accessibility | Input has accessible label `Greeting`; Save uses native submit button; live message uses polite announcement; input and button have visible keyboard focus; text contrast is at least 4.5:1 and UI border contrast is at least 3:1 as recorded in `design/design-system.md` |
| Responsive | Works at 320px viewport width and up with no horizontal page scroll; controls stack at widths of 520px and below |
| Privacy | Stores only greeting text; no personal data, account data, analytics, or external service data is stored |

## 7. Dependencies and assumptions

- **Depends on:** PostgreSQL persistence, for storing the current greeting across reloads and restarts.
- **Depends on:** Go backend API, for reading and saving the greeting between the frontend and persistence layer.
- **Depends on:** Next.js frontend, for rendering the approved greeting page and handling visitor input.
- **Assumption:** One shared greeting exists for the whole app, not one greeting per visitor. If false, scope changes to include identity or session behavior.

| Open question | Proposed default | Who decides |
|---|---|---|
| — | No open questions; stakeholder approved the single shared greeting behavior and no sign-in | — |

## 8. Traceability

Every plan item in this module appears exactly once, and every requirement id traces to a test case. A gap in this table is a gap in the build.

| Plan item | Requirement ids | Test cases |
|---|---|---|
| Persisted editable greeting | GREETING-001, GREETING-002, GREETING-003, GREETING-004 | `test-cases/persisted-editable-greeting.md` |
