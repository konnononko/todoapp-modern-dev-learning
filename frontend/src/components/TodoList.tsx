import TodoItem from './TodoItem';
import type { Todo } from '@/types'

type TodoListProps = {
    todos: Todo[];
    onToggle: (id: number) => void;
    onDelete: (id: number) => void;
};

const TodoList: React.FC<TodoListProps> = ({ todos, onToggle, onDelete }) => (
        <ul
            className="bg-white rounded-2xl shadow-xl max-w-md w-full mx-auto divide-y mt-8 overflow-hidden"
        >
            {todos.map((todo) => (
                <TodoItem
                    key={todo.id}
                    todo={todo}
                    onToggle={() => onToggle(todo.id)}
                    onDelete={() => onDelete(todo.id)}
                />
            ))}
        </ul>
    );

export default TodoList;