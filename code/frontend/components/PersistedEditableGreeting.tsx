"use client";

import { FormEvent, useRef, useState } from "react";
import { saveGreeting, type Greeting } from "../lib/persisted-editable-greeting";
import styles from "./PersistedEditableGreeting.module.css";

type PersistedEditableGreetingProps = {
  initialGreeting: Greeting;
};

export function PersistedEditableGreeting({
  initialGreeting,
}: PersistedEditableGreetingProps) {
  const [greeting, setGreeting] = useState(initialGreeting.text);
  const [inputValue, setInputValue] = useState(initialGreeting.text);
  const [message, setMessage] = useState("");
  const inputRef = useRef<HTMLInputElement>(null);

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const nextGreeting = inputValue.trim();

    if (!nextGreeting) {
      setMessage("Enter a greeting.");
      inputRef.current?.focus();
      return;
    }

    setGreeting(nextGreeting);
    setInputValue(nextGreeting);
    setMessage("Saved.");
  }

  return (
    <section className={styles.section} aria-labelledby="greeting-heading">
      <h1 id="greeting-heading" className={styles.heading}>
        {greeting}
      </h1>
      <form className={styles.form} onSubmit={handleSubmit} noValidate>
        <label className={styles.label} htmlFor="greeting-input">
          Greeting
        </label>
        <input
          ref={inputRef}
          id="greeting-input"
          name="greeting"
          type="text"
          value={inputValue}
          autoComplete="off"
          required
          onChange={(event) => setInputValue(event.target.value)}
          className={styles.input}
        />
        <button type="submit" className={styles.button}>
          Save
        </button>
      </form>
      <p className={styles.message} aria-live="polite">
        {message}
      </p>
    </section>
  );
}
