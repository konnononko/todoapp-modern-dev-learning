import './App.css';
import { useState, useRef } from "react";
import TodoList from './components/TodoList';
import TodoInput from './components/TodoInput';
import type { Todo } from '@/types'

const dummyTodos: Todo[] = [
  { id: 1, text: "コーヒー豆を買う", done: false },
  { id: 2, text: "日記を書く", done: false },
  { id: 3, text: "散歩する", done: true },
];

const App: React.FC<{}> = () => {
  const [todos, setTodos] = useState(
    dummyTodos
  );
  const nextId = useRef(dummyTodos.length + 1);

  const handleAddTodo = (text: string) => {
    setTodos([
      ...todos,
      { id: nextId.current++, text: text, done: false }
    ]);
  };

  const handleToggleTodo = (id: number) => {
    setTodos(todos =>
      todos.map(todo =>
        todo.id === id ? { ...todo, done: !todo.done } : todo
      )
    )
  }

  const handleDeleteTodo = (id: number) => {
    setTodos(todos => todos.filter(todo => todo.id !== id));
  };

  return (
    <>
      <div
        className="min-h-[calc(100vh-4rem)] bg-gray-100 flex flex-col items-center py-12"
      >
        <h1
          className="text-3xl font-bold mb-8 text-gray-800 drop-shadow"
        >
          Todoリスト
        </h1>
        <TodoInput onAdd={handleAddTodo} />
        <TodoList
          todos={todos}
          onToggle={handleToggleTodo}
          onDelete={handleDeleteTodo}
        />
      </div>
    </>
  );
}

export default App;
