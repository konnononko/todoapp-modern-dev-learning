import './App.css';
import { useState, useRef } from "react";
import TodoList from './components/TodoList';
import TodoInput from './components/TodoInput';

const dummyTodos = [
  { id: 1, text: "コーヒー豆を買う", done: false },
  { id: 2, text: "日記を書く", done: false },
  { id: 3, text: "散歩する", done: true },
];

function App() {
  const [todos, setTodos] = useState(
    dummyTodos
  );
  const nextId = useRef(dummyTodos.length + 1);

  const handleAddTodo = (text) => {
    setTodos([
      ...todos,
      { id: nextId.current++, text: text, done: false }
    ]);
  };

  const handleToggleTodo = (id) => {
    setTodos(todos =>
      todos.map(todo =>
        todo.id === id ? { ...todo, done: !todo.done } : todo
      )
    )
  }

  const handleDeleteTodo = (id) => {
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
