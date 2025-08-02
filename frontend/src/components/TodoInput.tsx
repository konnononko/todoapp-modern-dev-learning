import { useState } from "react";

type TodoInputProps = {
    onAdd: (text: string) => void;
}

const TodoInput: React.FC<TodoInputProps> = ({ onAdd }) => {
    const [value, setValue] = useState("");

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const text = value.trim();
        if (text) {
            onAdd(text);
            setValue("");
        }
    };

    return (
        <form
            onSubmit={handleSubmit}
            className="flex gap-2 mt-4"
        >
            <input
                type="text"
                value={value}
                onChange={e => setValue(e.target.value)}
                placeholder="新しいTodoを入力"
                className="flex-1 bg-white text-gray-900 border border-gray-300 rounded-lg p-2 focus:outline-none focus:ring-2 focus:ring-blue-400 placeholder-gray-400 w-full"
            />
            <button
                type="submit"
                className="bg-blue-500 text-white rounded-lg px-4 py-2 hover:bg-blue-600 transition"
            >
                追加
            </button>
        </form>
    );
};

export default TodoInput