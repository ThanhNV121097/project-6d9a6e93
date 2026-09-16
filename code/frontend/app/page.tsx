import { PersistedEditableGreeting } from "../components/PersistedEditableGreeting";
import { getGreeting } from "../lib/persisted-editable-greeting";

export default async function Page() {
  const persistedEditableGreeting = await getGreeting();

  return (
    <main>
      <PersistedEditableGreeting initialGreeting={persistedEditableGreeting} />
    </main>
  );
}
