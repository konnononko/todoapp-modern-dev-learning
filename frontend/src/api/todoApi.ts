import type { Todo } from '@/types'

const API_URL = 'http://localhost:8080/todos'

export type TodoApi = {
    getTodos: () => Promise<Todo[]>;
    addTodo: (text: string) => Promise<Todo>;
    updateTodo: (id: number, patch: Partial<Omit<Todo, 'id'>>) => Promise<Todo>;
    deleteTodo: (id: number) => Promise<void>;
}

export const todoApi: TodoApi = {
    getTodos,
    addTodo,
    updateTodo,
    deleteTodo,
}

export async function getTodos(): Promise<Todo[]> {
    const res = await fetch(API_URL);
    if (!res.ok) throw new Error('Failed to fetch todos');
    return res.json();
}

export async function addTodo(text: string): Promise<Todo> {
    const res = await fetch(API_URL, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title: text, done: false }),
    });
    if (!res.ok) throw new Error('Failed to add todo');
    return res.json();
}

export async function updateTodo(id: number, patch: Partial<Omit<Todo, 'id'>>): Promise <Todo> {
    const res = await fetch(`${API_URL}/${id}`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(patch),
    });
    if (!res.ok) throw new Error('Failed to update todo');
    return res.json();
}

export async function deleteTodo(id: number): Promise<void> {
    const res = await fetch(`${API_URL}/${id}`, { method: 'DELETE' });
    if (!res.ok) throw new Error('Failed to delete todo');
}
