import { FaTrash, FaInfoCircle } from "react-icons/fa";

function TodoItem({ todo, onToggle, onDelete }) {
    return (
        <li
            className="flex items-center gap-3 py-3 px-4 border-b last:border-none transition hover:bg-gray-50"
        >
            <input
                type="checkbox"
                checked={todo.done}
                onChange={onToggle}
                className="w-5 h-5 accent-blue-600"
            />
            <span
                className={`
                    flex-1 text-lg
                    ${todo.done ? "line-through text-gray-400" : "text-gray-900"}
                `}
            >
                {todo.text}
            </span>
            <button
                onClick={onDelete}
                className="ml-2 p-1 text-gray-400 hover:text-red-400 rounded-full transition"
                title="削除"
            >
                <FaTrash className="w-5 h-5" />
            </button>
            <div className="relative group">
                <FaInfoCircle
                    className="ml-2 w-5 h-5 text-gray-400 hover:text-blue-500 cursor-pointer"
                    aria-label="情報"
                    aria-describedby={`todo-info-id-${todo.id}`}
                />
                <span
                    className="absolute right-1/2 top-0 -translate-y-1/2 mr-2 scale-0 group-hover:scale-100 transition bg-gray-900 text-white text-xs px-2 py-1 rounded shadow-lg z-10 pointer-events-none whitespace-nowrap"
                    id={`todo-info-id-${todo.id}`}
                >
                    id: {todo.id}
                </span>
            </div>
        </li>
    );
}

export default TodoItem;