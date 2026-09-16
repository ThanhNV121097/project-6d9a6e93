INSERT INTO greetings (id, text, updated_at) VALUES (true, 'Hello, World!', now())
ON CONFLICT (id) DO UPDATE SET text = greetings.text;
