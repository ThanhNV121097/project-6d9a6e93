import { PersistedEditableGreeting } from "../components/PersistedEditableGreeting";
import { persistedEditableGreeting } from "../lib/mock/persisted-editable-greeting";

export default function Page() {
  return (
    <main>
      <PersistedEditableGreeting initialGreeting={persistedEditableGreeting} />
    </main>
  );
}
