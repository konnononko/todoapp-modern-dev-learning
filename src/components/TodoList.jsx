import TodoItem from './TodoItem';

function TodoList({ todos, onToggle, onDelete }) {
    return(
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
}

export default TodoList;