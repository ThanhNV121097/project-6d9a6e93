# Design System — Hello World Acceptance 9

> Source of truth: the approved `index.html`.
> Every value below is extracted from it. Changing a value here without
> changing the approved design is a defect.

Last updated: 2026-09-16

## 1. Foundations

### 1.1 Color

Semantic tokens. Name by job, never by hue.

| Token | Value | Used for |
|---|---|---|
| `--color-bg` | `#FFFFFF` | Page background, input background, text on primary action |
| `--color-text` | `#000000` | Body text, heading, input text, input border, focus outline |
| `--color-primary-action` | `#2563EB` | Save button background and border |

#### Contrast audit

Every text-on-background pair actually used. Body text ≥ 4.5:1, large text (≥ 18.66px bold or ≥ 24px) ≥ 3:1, UI borders ≥ 3:1.

| Foreground | Background | Ratio | Passes |
|---|---|---|---|
| `--color-text` | `--color-bg` | `21:1` | AA |
| `--color-bg` | `--color-primary-action` | `5.17:1` | AA |
| `--color-text` | `--color-bg` as input border / focus outline | `21:1` | UI |
| `--color-primary-action` | `--color-bg` as button border | `4.06:1` | UI |

### 1.2 Spacing

Base unit: `1px`. Every margin, padding, and gap in the product uses one of these.

| Token | Value |
|---|---|
| `--space-0` | `0` |
| `--space-hairline-offset` | `-1px` |
| `--space-hairline` | `1px` |
| `--space-focus-offset` | `3px` |
| `--space-control-y` | `10px` |
| `--space-control-x` | `12px` |
| `--space-form-gap` | `12px` |
| `--space-button-x` | `20px` |
| `--space-page-mobile` | `20px` |
| `--space-page` | `24px` |
| `--space-heading-gap` | `32px` |

### 1.3 Typography

Font families (include the fallback stack and how the font is loaded):

- Body: `Arial, Helvetica, sans-serif`; system/local fonts, no external font load.
- Headings: `Arial, Helvetica, sans-serif`; inherits body family.

| Token | Size | Line height | Weight | Used for |
|---|---|---|---|---|
| `--text-message` | `16px` | `1.5` | `400` inherited | Live form message |
| `--text-control` | `18px` | `normal` inherited | Input and Save button |
| `--text-heading` | `clamp(44px, 10vw, 80px)` | `1.05` | `700` | Page h1 greeting |

Heading levels are used in order: one `h1` only.

Weight and letter-spacing are tokens too, not just columns in the table above:

| Token | Value | Used for |
|---|---|---|
| `--font-weight-body` | `400` inherited browser/body default | Running text, input, message |
| `--font-weight-action` | `700` | Save button label |
| `--font-weight-heading` | `700` | h1 greeting |
| `--tracking-heading-tight` | `-0.04em` | h1 greeting |
| `--tracking-normal` | `normal` | Controls and message |

### 1.4 Radius, border, shadow, motion

| Token | Value | Used for |
|---|---|---|
| `--radius-control` | `0` | Input and Save button |
| `--border-control` | `2px` | Input and Save button border |
| `--focus-outline-width` | `3px` | Input and Save button focus outline |
| `--focus-outline-offset` | `3px` | Input and Save button focus outline offset |

Motion: no animation or transition is used in approved design.

### 1.5 Layout and breakpoints

| Name | Max width | Container | Columns | Gutter |
|---|---|---|---|---|
| `compact` | `520px` | `100%` page padding `20px` | Form stacks to one column | `12px` form gap |
| `default` | above `520px` | Section max width `560px`; input max width `360px` | Form row: flexible input plus button | `12px` form gap |

## 2. Components

### 2.1 Greeting section

**Purpose** — Centers single greeting update task. Do not use for navigation or multi-section content.

**Anatomy** — `[heading] [form] [live message]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Centered | `--space-page`, `--space-page-mobile`, section max width `560px` | Only page section |

**Sizes**

| Size | Width | Padding | Text token |
|---|---|---|---|
| Default | `100%`, max `560px` | Page padding `24px` | `--text-heading` |
| Compact | `100%` | Page padding `20px` | `--text-heading` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Centered black text on white background | `--color-bg`, `--color-text` |

**Accessibility** — Section uses `aria-labelledby` pointing to h1. Main landmark wraps section.

### 2.2 Greeting heading

**Purpose** — Shows persisted greeting as primary page content.

**Anatomy** — `[greeting text]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Large greeting | `--text-heading`, `--font-weight-heading`, `--tracking-heading-tight` | Page greeting only |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Fluid | Content height | `0`; bottom margin `32px` | `--text-heading` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Large bold black centered text; wraps anywhere | `--color-text`, `--text-heading` |

**Accessibility** — Use one `h1`; preserve text content from stored greeting; allow wrapping for long greetings.

### 2.3 Greeting input

**Purpose** — Lets visitor edit greeting before saving.

**Anatomy** — `[visually hidden label] [text input]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Text greeting | `--color-bg`, `--color-text`, `--border-control`, `--radius-control` | Greeting form only |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | min `48px` | `10px 12px` | `--text-control` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | White field, black text, 2px black border | `--color-bg`, `--color-text`, `--border-control` |
| Focus visible | 3px black outline offset by 3px | `--color-text`, `--focus-outline-width`, `--focus-outline-offset` |

**Accessibility** — Text input has label text `Greeting`, hidden visually but available to assistive tech. Required input. Keyboard focus must show visible outline. Minimum hit target is at least `48px` high.

### 2.4 Save button

**Purpose** — Submits greeting change.

**Anatomy** — `[label]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Primary action | `--color-primary-action`, `--color-bg`, `--font-weight-action` | Saving greeting |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | min `48px` | `10px 20px` | `--text-control` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Blue fill and border, white bold text, pointer cursor | `--color-primary-action`, `--color-bg`, `--font-weight-action` |
| Focus visible | 3px black outline offset by 3px | `--color-text`, `--focus-outline-width`, `--focus-outline-offset` |

**Accessibility** — Native `button type="submit"`; keyboard activation through Enter/Space. Focus must show visible outline. Minimum hit target is at least `48px` high.

### 2.5 Form message

**Purpose** — Announces validation and save feedback under form.

**Anatomy** — `[message text]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Live text | `--text-message`, `--color-text` | Feedback below greeting form |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | min `24px` | `0`; top margin `12px` | `--text-message` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Empty | Reserves 24px height with no text | `--text-message` |
| Success | Shows `Saved.` in black text | `--color-text`, `--text-message` |
| Error | Shows `Enter a greeting.` in black text | `--color-text`, `--text-message` |

**Accessibility** — Message uses `aria-live="polite"` so updates are announced without moving focus except invalid empty input returns focus to input.

## 3. Content and formatting

- Voice and tone: direct, minimal system copy.
- Date, time, number, and currency formats: not displayed in approved design.
- Capitalization rule: heading preserves stored greeting exactly; button uses title case `Save`; validation and success messages use sentence case.
- Empty-state and error-message wording pattern: form message starts empty; error is imperative and specific: `Enter a greeting.`; success is concise: `Saved.`.

## 4. Known deviations

Places where the approved design does not follow its own rules or the anti-patterns in `references/ai-defaults.md`. Record, do not silently fix.

| Where | Deviation | Why it stands | Follow-up |
|---|---|---|---|
| Spacing scale | Uses exact one-off values `3px`, `10px`, `20px`, `24px`, `32px` instead of a smaller 4px/8px scale | Approved mockup is intentionally plain and minimal; values come from source HTML | Keep exact for this project unless stakeholder requests redesign |
| Interactive states | Button and input show default and focus-visible only; no hover, active, or disabled visual states are drawn | Approved design includes only default and focus styles | Do not invent extra states downstream unless design changes |

AI-default checks avoided by approved design: no purple/indigo default palette, no gradients, no maximum rounding, no heavy shadows, no generic multi-section layout, no emoji iconography, no filler copy, no removed focus state, no text over images, no hover-only affordance.

## 5. Change log

| Date | Change | Design PR |
|---|---|---|
| 2026-09-16 | Initial design system extracted from approved `index.html` | This PR |
