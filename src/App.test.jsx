import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import App from './App'

async function addTodo(text) {
    const input = screen.getByPlaceholderText(/新しいTodoを入力/i);
    const button = screen.getByRole('button', { name: '追加'});
    await userEvent.type(input, text);
    await userEvent.click(button);
}

test('Todoを追加できる', async () => {
    render(<App />);

    const testTodo = 'テストを実施する';

    expect(screen.queryByText(testTodo)).not.toBeInTheDocument();

    await addTodo(testTodo);

    expect(screen.getByText(testTodo)).toBeInTheDocument();
});

test('Todoを削除できる', async () => {
    render(<App />);

    const testTodo = 'ラジオを聴く';
    await addTodo(testTodo);

    const todo =screen.getByText(testTodo);
    expect(todo).toBeInTheDocument();

    const deleteButton = todo.closest('li').querySelector('button[title="削除"]');
    await userEvent.click(deleteButton);

    expect(todo).not.toBeInTheDocument();
});

test('チェックボックスで完了状態を切り替えられる。完了状態ではテキストは打ち消し線のスタイルになる。', async () => {
    render (<App />);
    const testTodo = 'ダンスを踊る';
    await addTodo(testTodo);

    const todo = screen.getByText(testTodo);
    const checkbox = todo.closest('li').querySelector('input[type="checkbox"]');

    expect(checkbox.checked).toBe(false);
    expect(todo).not.toHaveClass('line-through');

    await userEvent.click(checkbox);
    expect(checkbox.checked).toBe(true);
    expect(todo).toHaveClass('line-through');

    await userEvent.click(checkbox);
    expect(checkbox.checked).toBe(false);
    expect(todo).not.toHaveClass('line-through');
});

test ('Todoを追加するとidが一意に割り振られる', async () => {
    render(<App />);
    await addTodo('A');
    await addTodo('B');

    const ids = screen.getAllByLabelText('情報')
        .map(icon => document.getElementById(icon.getAttribute('aria-describedby')))
        .map(span => parseInt(span.textContent.replace('id: ', ''), 10));

    expect(ids.length).toBeGreaterThanOrEqual(2);
    expect(ids.length).toBe(new Set(ids).size);
    expect(
        ids.slice(1).every((id, i) => id === ids[i] + 1)
    ).toBe(true);
});